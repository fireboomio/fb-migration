package consts

import "os"

const (
	OperationApiSuffix    = ".graphql"
	OperationAPISwitchOff = ".off"
	OperationAPISwitchOn  = ""
	JsonExt               = ".json"
	ConfigJsonExt         = ".config.json"
)

var OperationAPISwitchMap = map[bool]string{
	true:  OperationAPISwitchOn,
	false: OperationAPISwitchOff,
}

const (
	PathSep = string(os.PathSeparator)
)

const (
	ExportedPath                        = "exported"
	ExportedOperationsPath              = ExportedPath + PathSep + "operations"
	ExportedGeneratedPath               = ExportedPath + PathSep + "generated"
	ExportedGeneratedFireboomConfigPath = ExportedGeneratedPath + PathSep + "fireboom.config.json"
	ExportedGeneratedWdgPostmanPath     = ExportedGeneratedPath + PathSep + "wundergraph.postman.json"
	ExportedGeneratedSwaggerPath        = ExportedGeneratedPath + PathSep + "swagger.json"
	ExportedGeneratedSchemaGraphql      = ExportedGeneratedPath + PathSep + "fireboom.app.schema.graphql"
)

const (
	StorePath      = "store"                      // 各种配置文件
	StoreCloudPath = "store-cloud"                // 各种配置文件
	StoreListPath  = StorePath + PathSep + "list" // json db里面存储的内容

	StoreListDatasourcePath  = StoreListPath + PathSep + "FbDataSource"
	StoreCloudDatasourcePath = StoreCloudPath + PathSep + "datasource"

	StoreListSDKPath  = StoreListPath + PathSep + "FbSDK"
	StoreCloudSDKPath = StoreCloudPath + PathSep + "sdk"

	StoreListOperationPath  = StoreListPath + PathSep + "FbOperation"
	StoreCloudOperationPath = StoreCloudPath + PathSep + "operation"

	StoreListRolePath  = StoreListPath + PathSep + "FbRole"
	StoreCloudRolePath = StoreCloudPath + PathSep + "role"

	StoreListStoragePath  = StoreListPath + PathSep + "FbStorageBucket"
	StoreCloudStoragePath = StoreCloudPath + PathSep + "storage"

	StoreObjectPath                     = StorePath + PathSep + "object"                             // 其他配置
	StoreObjectOperationsPath           = StoreObjectPath + PathSep + "operations"                   // 其他配置
	StoreGlobalConfigPath               = StoreObjectPath + PathSep + "global_config.json"           // store/object/global_config.json
	StoreGlobalSystemConfigPath         = StoreObjectPath + PathSep + "global_system_config.json"    // store/object/global_system_config.json
	StoreGlobalOperationConfigPath      = StoreObjectPath + PathSep + "global_operation_config.json" // store/object/global_operation_config.json
	StoreCloudGlobalOperationConfigPath = StoreCloudPath + PathSep + "config" + PathSep + "global.operation.json"
	StoreCloudGlobalSettingPath         = StoreCloudPath + PathSep + "config" + PathSep + "global.setting.json"

	StoreHooksPath          = StorePath + PathSep + "hooks"              //钩子config目录
	StoreHooksAuthPath      = StoreHooksPath + PathSep + CustomAuth      //auth钩子config目录
	StoreHooksCustomizePath = StoreHooksPath + PathSep + CustomCustomize //customize钩子config目录
	StoreHooksGlobalPath    = StoreHooksPath + PathSep + CustomGlobal    //global钩子config目录
	StoreHooksHooksPath     = StoreHooksPath + PathSep + CustomHooks     //operations钩子config目录
	StoreHooksUploadsPath   = StoreHooksPath + PathSep + CustomUploads   //uploads钩子config目录
)

const (
	CustomAuth          = "auth"
	CustomGlobal        = "global"
	CustomCustomize     = "customize"
	CustomHooks         = "hooks"
	CustomUploads       = "uploads"
	CustomGenerated     = "generated"
	CustomOperations    = "operations"
	CustomProxyRequests = "proxys"
)

const (
	UploadPath      = "upload"
	UploadCloudPath = "upload-cloud"

	UploadOasPath      = UploadPath + PathSep + "oas"
	UploadCloudOasPath = UploadPath + PathSep + "oas"

	UploadDbpath      = UploadPath + PathSep + "db"
	UploadCloudDbpath = UploadCloudPath + PathSep + "db"

	UploadGraphqlPath      = UploadPath + PathSep + "graphql"
	UploadCloudGraphqlPath = UploadCloudPath + PathSep + "graphql"
)
