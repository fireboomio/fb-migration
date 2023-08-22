package migrate

import (
	"encoding/json"
	"fireboom-migrate/consts"
	"fireboom-migrate/types/cloud"
	"fireboom-migrate/types/origin"
	"fireboom-migrate/types/wgpb"
	"fireboom-migrate/utils"
	"fmt"
	"path/filepath"
	"strings"
)

var opSettingMap map[string]*cloud.OperationSetting

func init() {
	opSettingMap = make(map[string]*cloud.OperationSetting)
}

// 迁移operation
// 1. store/list/FbOperation --> store-cloud/operation/xx.json
// 2. store/object/operations/xx.json --> 独立配置
// 3. exported/operations/xx.graphql --> copy to store-cloud/operation/xx.graphql
// 4. store/hooks/hooks  --> api钩子
func migrateOperation() {
	operationPath := consts.StoreListOperationPath
	content, err := utils.ReadFile(operationPath)
	if err != nil {
		fmt.Errorf("read file failed, err: %v\n", err.Error())
		return
	}

	var ops = make([]origin.FbOperation, 0)
	_ = json.Unmarshal(content, &ops)

	for _, op := range ops {
		operation := &cloud.Operation{
			Enabled:       op.Enabled,
			Remark:        op.Remark,
			CreateTime:    op.CreateTime,
			DeleteTime:    op.DeleteTime,
			UpdateTime:    op.UpdateTime,
			Path:          strings.TrimPrefix(op.Path, "/"),
			Engine:        wgpb.OperationExecutionEngine_ENGINE_GRAPHQL,
			OperationType: switchOperationType(op.OperationType),

			LiveQueryConfig:      getLiveQueryConfig(&op),
			ConfigCustomized:     getConfigCustomized(&op),
			CacheConfig:          getCacheConfig(op.Path),
			AuthenticationConfig: getAuthenticationConfig(op.Path),
			HooksConfiguration:   getHooksConfig(op.Path),
		}

		//copy graphql
		srcPath := filepath.Join(consts.ExportedOperationsPath, op.Path) + consts.OperationApiSuffix
		if !op.Enabled {
			srcPath += consts.OperationAPISwitchOff
		}

		err = utils.CopyFile(srcPath,
			filepath.Join(consts.StoreCloudOperationPath, op.Path)+consts.OperationApiSuffix)
		if err != nil {
			fmt.Errorf("copy file failded, err: %v\n", err.Error())
		}

		bytes, _ := json.MarshalIndent(operation, "", "    ")
		opPath := utils.GetStoreCloudModelPath(consts.StoreCloudOperationPath, operation.Path)
		if utils.NotExistFile(opPath) {
			utils.CreateFile(opPath)
		}

		err = utils.WriteFile(opPath, bytes)
		fmt.Printf("write file %s\n", opPath)
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}
}

func getConfigCustomized(op *origin.FbOperation) bool {
	setting, ok := opSettingMap[op.Path]
	if !ok {
		return false
	}
	return setting.Enabled
}

func getHooksConfig(path string) *wgpb.OperationHooksConfiguration {
	path = filepath.Join(consts.StoreHooksHooksPath, path)
	if utils.NotExistFile(path) {
		return nil
	}

	res := &wgpb.OperationHooksConfiguration{}

	// 扫描 store/hooks/hooks 目录
	fileList, err := utils.ReadDir(path)
	if err != nil {
		fmt.Errorf("read dir %s failed, err: %s", path, err.Error())
		return res
	}

	// api局部钩子
	for _, fileName := range fileList {
		shortName := strings.TrimSuffix(fileName, consts.ConfigJsonExt)
		switch shortName {
		case consts.PreResolve:
			res.PreResolve = utils.GetHookConfigEnable(filepath.Join(path, fileName))
		case consts.MutatingPreResolve:
			res.MutatingPreResolve = utils.GetHookConfigEnable(filepath.Join(path, fileName))
		case consts.MockResolve:
			res.MockResolve = &wgpb.MockResolveHookConfiguration{
				Enabled: utils.GetHookConfigEnable(filepath.Join(path, fileName)),
			}
		case consts.CustomResolve:
			res.CustomResolve = utils.GetHookConfigEnable(filepath.Join(path, fileName))
		case consts.PostResolve:
			res.PostResolve = utils.GetHookConfigEnable(filepath.Join(path, fileName))
		case consts.MutatingPostResolve:
			res.MutatingPostResolve = utils.GetHookConfigEnable(filepath.Join(path, fileName))
		}
	}

	return res
}

func getAuthenticationConfig(path string) *wgpb.OperationAuthenticationConfig {
	setting, ok := opSettingMap[path]
	if !ok {
		return nil
	}
	return &wgpb.OperationAuthenticationConfig{
		AuthRequired: setting.AuthenticationRequired,
	}
}

func getCacheConfig(path string) *wgpb.OperationCacheConfig {
	setting, ok := opSettingMap[path]
	if !ok {
		return nil
	}
	return &wgpb.OperationCacheConfig{
		Enabled:              setting.CachingEnabled,
		MaxAge:               setting.CachingMaxAge,
		Public:               false, // ??
		StaleWhileRevalidate: setting.CachingStaleWhileRevalidate,
	}
}

func switchOperationType(typeStr string) wgpb.OperationType {
	switch typeStr {
	case "queries":
		return wgpb.OperationType_QUERY
	case "mutations":
		return wgpb.OperationType_MUTATION
	case "subscriptions":
		return wgpb.OperationType_SUBSCRIPTION
	default:
		return wgpb.OperationType_QUERY
	}
}

func getLiveQueryConfig(operation *origin.FbOperation) *wgpb.OperationLiveQueryConfig {
	res := &wgpb.OperationLiveQueryConfig{
		Enabled: operation.LiveQuery,
	}
	path := filepath.Join(consts.StoreObjectOperationsPath, operation.Path) + consts.JsonExt
	content, err := utils.ReadFile(path)
	if err != nil {
		fmt.Errorf("read file failed, err: %v", err.Error())
		return res
	}

	opSetting := &cloud.OperationSetting{}
	_ = json.Unmarshal(content, opSetting)

	if operation.LiveQuery {
		res.PollingIntervalSeconds = opSetting.LiveQueryPollingIntervalSeconds
	}

	opSettingMap[operation.Path] = opSetting
	fmt.Printf("operation setting path: %s\n", operation.Path)
	return res
}
