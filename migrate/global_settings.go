package migrate

import (
	"encoding/json"
	"fireboom-migrate/consts"
	"fireboom-migrate/types/cloud"
	"fireboom-migrate/types/origin"
	"fireboom-migrate/types/wgpb"
	"fireboom-migrate/utils"
	"fmt"
	"time"
)

func migrateGlobalSettings() {
	globalConfigPath := consts.StoreGlobalConfigPath
	content1, err := utils.ReadFile(globalConfigPath)
	if err != nil {
		fmt.Errorf("read file failed, err: %v", err.Error())
	}

	//获取global_system_config.json
	globalSystemConfigPath := consts.StoreGlobalSystemConfigPath
	content2, err := utils.ReadFile(globalSystemConfigPath)
	if err != nil {
		fmt.Errorf("read file failed, err: %v", err.Error())
	}

	//反序列化
	globalConfig := &origin.GlobalConfig{}
	err = json.Unmarshal(content1, globalConfig)
	if err != nil {
		fmt.Errorf("json unmarshal failed, err: %v", err.Error())
		return
	}

	globalSystemConfig := &origin.Setting{}
	err = json.Unmarshal(content2, &globalSystemConfig)
	if err != nil {
		fmt.Errorf("json unmarshal failed, err: %v", err.Error())
		return
	}

	setting := cloud.GlobalSetting{
		NodeOptions:       getNodeOptions(),
		ServerOptions:     getServerOptions(),
		CorsConfiguration: getCorsConfiguration(globalConfig),
		SecurityConfig:    getSecurityConfig(globalConfig),
		Appearance:        &cloud.Appearance{Language: ""},
		BuildInfo: &cloud.BuildInfo{
			BuiltBy: "Fireboom",
		},
	}

	bytes, err := json.MarshalIndent(setting, "", "    ")
	if err != nil {
		fmt.Errorf("read file failed, err: %v\n", err.Error())
	}

	//写入文件
	dsPath := consts.StoreCloudGlobalSettingPath
	if utils.NotExistFile(dsPath) {
		utils.CreateFile(dsPath)
	}

	err = utils.WriteFile(dsPath, bytes)
	fmt.Printf("write file %s\n", dsPath)
	if err != nil {
		fmt.Errorf("write file failed, err: %v\n", err.Error())
	}
}

// 获取ServerOptions
func getServerOptions() *wgpb.ServerOptions {
	return &wgpb.ServerOptions{
		ServerUrl: utils.GetConfigurationVariable(origin.Value{
			Kind: "1",
			Val:  "FB_SERVER_URL",
		}),
		Listen: &wgpb.ListenerOptions{
			Host: utils.GetConfigurationVariable(origin.Value{
				Kind: "1",
				Val:  "FB_SERVER_LISTEN_HOST",
			}),
			Port: utils.GetConfigurationVariable(origin.Value{
				Kind: "1",
				Val:  "FB_SERVER_LISTEN_PORT",
			}),
		},
	}
}

// 获取NodeOptions
func getNodeOptions() *wgpb.NodeOptions {

	return &wgpb.NodeOptions{
		NodeUrl: utils.GetConfigurationVariable(origin.Value{
			Kind: "1",
			Val:  consts.INTERNAL_URL,
		}),

		PublicNodeUrl: utils.GetConfigurationVariable(origin.Value{
			Kind: "1",
			Val:  consts.PUBLIC_URL,
		}),

		Listen: &wgpb.ListenerOptions{
			Host: utils.GetConfigurationVariable(origin.Value{
				Kind: "1",
				Val:  consts.LISTEN_HOST,
			}),

			Port: utils.GetConfigurationVariable(origin.Value{

				Kind: "1",
				Val:  consts.LISTEN_PORT,
			}),
		},

		Logger: &wgpb.NodeLogging{Level: utils.GetConfigurationVariable(origin.Value{
			Kind: "1",
			Val:  consts.LOG_LEVEL,
		})},
		DefaultRequestTimeoutSeconds: 0,
	}
}

// 获取SecurityConfig
func getSecurityConfig(cfg *origin.GlobalConfig) cloud.SecurityConfig {

	return cloud.SecurityConfig{
		AllowedReport:          true,
		EnableGraphqlEndpoint:  cfg.ConfigureWunderGraphApplication.Security.EnableGraphQLEndpoint,
		AllowedHostNames:       transformOriginStringToConfigurationVariable(cfg.ConfigureWunderGraphApplication.Security.AllowedHosts),
		AuthorizedRedirectUris: transformOriginStringToConfigurationVariable(cfg.AuthRedirectURL),
		AuthenticationKey:      "",
		EnableCSRFProtect:      cfg.ConfigureWunderGraphApplication.Security.EnableCSRFProtect,
		ForceHttpsRedirects:    cfg.ForceHttpsRedirects,
		GlobalRateLimit: struct {
			Enabled     bool          `json:"enabled"`
			Requests    int           `json:"requests"`
			PerDuration time.Duration `json:"perDuration"`
		}{},
	}
}

// 获取corsConfiguration
func getCorsConfiguration(cfg *origin.GlobalConfig) *wgpb.CorsConfiguration {
	return &wgpb.CorsConfiguration{
		AllowedOrigins:   transformOriginStringToConfigurationVariable(cfg.ConfigureWunderGraphApplication.Cors.AllowedOrigins),
		AllowedMethods:   cfg.ConfigureWunderGraphApplication.Cors.AllowedMethods,
		AllowedHeaders:   cfg.ConfigureWunderGraphApplication.Cors.AllowedHeaders,
		ExposedHeaders:   cfg.ConfigureWunderGraphApplication.Cors.ExposedHeaders,
		MaxAge:           cfg.ConfigureWunderGraphApplication.Cors.MaxAge,
		AllowCredentials: cfg.ConfigureWunderGraphApplication.Cors.AllowCredentials,
	}
}

func transformOriginStringToConfigurationVariable(origins []string) []*wgpb.ConfigurationVariable {
	res := make([]*wgpb.ConfigurationVariable, len(origins))

	for _, origin := range origins {
		tmp := &wgpb.ConfigurationVariable{StaticVariableContent: origin}
		res = append(res, tmp)
	}
	return res
}
