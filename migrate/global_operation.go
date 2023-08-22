package migrate

import (
	"encoding/json"
	"fireboom-migrate/consts"
	"fireboom-migrate/types/cloud"
	"fireboom-migrate/types/origin"
	"fireboom-migrate/types/wgpb"
	"fireboom-migrate/utils"
	"fmt"
	"strings"
)

// 迁移全局operation配置
// store/object/global_operations_config.json --> store-cloud/config/global.operation.json
func migrateGlobalOperation() {
	path := consts.StoreGlobalOperationConfigPath
	content, err := utils.ReadFile(path)
	if err != nil {
		fmt.Errorf("read file failed, err: %v\n", err.Error())
	}

	globalOperationConfig := &origin.OperationSetting{}
	_ = json.Unmarshal(content, globalOperationConfig)

	authenticationConfigs := make(map[wgpb.OperationType]*wgpb.OperationAuthenticationConfig, 3)
	authenticationConfigs[wgpb.OperationType_QUERY] = &wgpb.OperationAuthenticationConfig{
		AuthRequired: globalOperationConfig.AuthenticationQueriesRequired,
	}
	authenticationConfigs[wgpb.OperationType_MUTATION] = &wgpb.OperationAuthenticationConfig{
		AuthRequired: globalOperationConfig.AuthenticationMutationsRequired,
	}
	authenticationConfigs[wgpb.OperationType_SUBSCRIPTION] = &wgpb.OperationAuthenticationConfig{
		AuthRequired: globalOperationConfig.AuthenticationSubscriptionsRequired,
	}

	newConfig := &cloud.GlobalOperation{
		CacheConfig: &wgpb.OperationCacheConfig{
			Enabled:              globalOperationConfig.Enabled,
			MaxAge:               globalOperationConfig.CachingMaxAge,
			Public:               false,
			StaleWhileRevalidate: globalOperationConfig.CachingStaleWhileRevalidate,
		},
		LiveQueryConfig: &wgpb.OperationLiveQueryConfig{
			Enabled:                globalOperationConfig.Enabled,
			PollingIntervalSeconds: globalOperationConfig.LiveQueryPollingIntervalSeconds,
		},
		AuthenticationConfigs:    authenticationConfigs,
		ApiAuthenticationHooks:   getApiAuthHooks(),
		GlobalHttpTransportHooks: getGlobalHttpTransportHooks(),
	}

	bytes, _ := json.MarshalIndent(newConfig, "", "    ")
	err = utils.WriteFile(consts.StoreCloudGlobalOperationConfigPath, bytes)
	fmt.Printf("write file %s\n", consts.StoreCloudGlobalOperationConfigPath)
	if err != nil {
		fmt.Errorf("write file failed, path: %s, err: %v\n", consts.StoreCloudGlobalOperationConfigPath, err.Error())
	}
}

func getGlobalHttpTransportHooks() map[string]bool {
	globalPath := consts.StoreHooksGlobalPath
	nameLists, err := utils.ReadDir(globalPath)
	if len(nameLists) == 0 {
		return nil
	}
	if err != nil {
		fmt.Errorf("read dir failed,err: %v\n", err.Error())
		return nil
	}

	res := make(map[string]bool, 0)
	for _, name := range nameLists {
		path := globalPath + consts.PathSep + name
		suffix := strings.TrimSuffix(name, ".config.json")

		switch suffix {
		case consts.HttpTransportBeforeRequest:
			res[consts.HttpTransportBeforeRequest] = utils.GetHookConfigEnable(path)
			break
		case consts.HttpTransportOnRequest:
			res[consts.HttpTransportOnRequest] = utils.GetHookConfigEnable(path)
			break
		case consts.HttpTransportOnResponse:
			res[consts.HttpTransportOnResponse] = utils.GetHookConfigEnable(path)
			break
		}
	}

	return res
}

// 全局认证钩子
func getApiAuthHooks() map[string]bool {
	authPath := consts.StoreHooksAuthPath
	nameLists, err := utils.ReadDir(authPath)
	if len(nameLists) == 0 {
		return nil
	}
	if err != nil {
		fmt.Errorf("read dir failed,err: %v", err.Error())
		return nil
	}

	res := make(map[string]bool, 0)
	for _, name := range nameLists {
		path := authPath + consts.PathSep + name
		suffix := strings.TrimSuffix(name, ".config.json")

		switch suffix {
		case consts.PostAuthentication:
			res[consts.PostAuthentication] = utils.GetHookConfigEnable(path)
			break
		case consts.MutatingPostAuthentication:
			res[consts.MutatingPostAuthentication] = utils.GetHookConfigEnable(path)
			break
		case consts.RevalidateAuthentication:
			res[consts.RevalidateAuthentication] = utils.GetHookConfigEnable(path)
			break
		case consts.PostLogout:
			res[consts.PostLogout] = utils.GetHookConfigEnable(path)
			break
		}
	}
	return res
}
