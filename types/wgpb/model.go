package wgpb

import (
	"google.golang.org/protobuf/runtime/protoimpl"
)

type (
	CustomGraphql struct {
		Url          string                 `json:"url"`
		Headers      map[string]*HTTPHeader `json:"headers"`
		Customized   bool                   `json:"customized"`
		SchemaString string                 `json:"schemaString"`
	}
	CustomRest struct {
		OasFilepath string                 `json:"oasFilepath"`
		BaseUrl     string                 `json:"baseUrl"`
		Headers     map[string]*HTTPHeader `json:"headers"`
	}
	CustomDatabase struct {
		Kind          int64                  `json:"kind"`
		DatabaseUrl   *ConfigurationVariable `json:"databaseUrl"`
		DatabaseAlone *CustomDatabaseAlone   `json:"databaseAlone"`
	}
	CustomDatabaseAlone struct {
		Host     string `json:"host"`
		Port     int64  `json:"port"`
		Database string `json:"database"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

type HTTPHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*ConfigurationVariable `protobuf:"bytes,1,rep,name=values,proto3" json:"values"`
}

type ConfigurationVariable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind                            int32  `protobuf:"varint,1,opt,name=kind,proto3,enum=wgpb.ConfigurationVariableKind" json:"kind"`
	StaticVariableContent           string `protobuf:"bytes,2,opt,name=staticVariableContent,proto3" json:"staticVariableContent,omitempty"`                     // [omitempty]
	EnvironmentVariableName         string `protobuf:"bytes,3,opt,name=environmentVariableName,proto3" json:"environmentVariableName,omitempty"`                 // [omitempty]
	EnvironmentVariableDefaultValue string `protobuf:"bytes,4,opt,name=environmentVariableDefaultValue,proto3" json:"environmentVariableDefaultValue,omitempty"` // [omitempty]
	PlaceholderVariableName         string `protobuf:"bytes,5,opt,name=placeholderVariableName,proto3" json:"placeholderVariableName,omitempty"`                 // [omitempty]
}

type ConfigurationVariableKind int32

const (
	ConfigurationVariableKind_STATIC_CONFIGURATION_VARIABLE      ConfigurationVariableKind = 0
	ConfigurationVariableKind_ENV_CONFIGURATION_VARIABLE         ConfigurationVariableKind = 1
	ConfigurationVariableKind_PLACEHOLDER_CONFIGURATION_VARIABLE ConfigurationVariableKind = 2
)

type CustomDatabaseKind int32

const (
	customDatabaseKindUrl   CustomDatabaseKind = 0
	customDatabaseKindAlone CustomDatabaseKind = 1
)

type DataSourceKind int32

const (
	DataSourceKind_STATIC     DataSourceKind = 0
	DataSourceKind_REST       DataSourceKind = 1
	DataSourceKind_GRAPHQL    DataSourceKind = 2
	DataSourceKind_POSTGRESQL DataSourceKind = 3
	DataSourceKind_MYSQL      DataSourceKind = 4
	DataSourceKind_SQLSERVER  DataSourceKind = 5
	DataSourceKind_MONGODB    DataSourceKind = 6
	DataSourceKind_SQLITE     DataSourceKind = 7
	DataSourceKind_PRISMA     DataSourceKind = 8
)

type OperationExecutionEngine int32

const (
	OperationExecutionEngine_ENGINE_GRAPHQL  OperationExecutionEngine = 0
	OperationExecutionEngine_ENGINE_FUNCTION OperationExecutionEngine = 1
	OperationExecutionEngine_ENGINE_PROXY    OperationExecutionEngine = 2
)

type OperationType int32

const (
	OperationType_QUERY        OperationType = 0
	OperationType_MUTATION     OperationType = 1
	OperationType_SUBSCRIPTION OperationType = 2
)

type OperationHooksConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreResolve                 bool                          `protobuf:"varint,1,opt,name=preResolve,proto3" json:"preResolve"`
	PostResolve                bool                          `protobuf:"varint,2,opt,name=postResolve,proto3" json:"postResolve"`
	MutatingPreResolve         bool                          `protobuf:"varint,3,opt,name=mutatingPreResolve,proto3" json:"mutatingPreResolve"`
	MutatingPostResolve        bool                          `protobuf:"varint,4,opt,name=mutatingPostResolve,proto3" json:"mutatingPostResolve"`
	CustomResolve              bool                          `protobuf:"varint,8,opt,name=customResolve,proto3" json:"customResolve"`
	MockResolve                *MockResolveHookConfiguration `protobuf:"bytes,5,opt,name=mockResolve,proto3" json:"mockResolve"`
	HttpTransportBeforeRequest bool                          `protobuf:"varint,11,opt,name=httpTransportBeforeRequest,proto3" json:"httpTransportBeforeRequest"`
	HttpTransportOnRequest     bool                          `protobuf:"varint,6,opt,name=httpTransportOnRequest,proto3" json:"httpTransportOnRequest"`
	HttpTransportOnResponse    bool                          `protobuf:"varint,7,opt,name=httpTransportOnResponse,proto3" json:"httpTransportOnResponse"`
	OnConnectionInit           bool                          `protobuf:"varint,10,opt,name=onConnectionInit,proto3" json:"onConnectionInit"`
	TsPathMap                  map[string]string             `protobuf:"bytes,99,rep,name=tsPathMap,proto3" json:"tsPathMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // [omitempty]
}

type MockResolveHookConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled                           bool  `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled"`
	SubscriptionPollingIntervalMillis int64 `protobuf:"varint,2,opt,name=subscriptionPollingIntervalMillis,proto3" json:"subscriptionPollingIntervalMillis"`
}

type OperationCacheConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled              bool  `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled"`
	MaxAge               int64 `protobuf:"varint,2,opt,name=maxAge,proto3" json:"maxAge"`
	Public               bool  `protobuf:"varint,3,opt,name=public,proto3" json:"public"`
	StaleWhileRevalidate int64 `protobuf:"varint,4,opt,name=staleWhileRevalidate,proto3" json:"staleWhileRevalidate"`
}

type OperationLiveQueryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled                bool  `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled"`
	PollingIntervalSeconds int64 `protobuf:"varint,2,opt,name=pollingIntervalSeconds,proto3" json:"pollingIntervalSeconds"`
}

type OperationAuthenticationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthRequired bool `protobuf:"varint,1,opt,name=authRequired,proto3" json:"authRequired"`
}

type OperationAuthorizationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Claims     []*ClaimConfig       `protobuf:"bytes,1,rep,name=claims,proto3" json:"claims"`
	RoleConfig *OperationRoleConfig `protobuf:"bytes,2,opt,name=roleConfig,proto3" json:"roleConfig"`
}

type ClaimConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VariablePathComponents []string `protobuf:"bytes,1,rep,name=variablePathComponents,proto3" json:"variablePathComponents"`
	ClaimType              int32    `protobuf:"varint,2,opt,name=claimType,proto3,enum=wgpb.ClaimType" json:"claimType"`
	// Available iff claimType == CUSTOM
	Custom *CustomClaim `protobuf:"bytes,3,opt,name=custom,proto3,oneof" json:"custom"`
}

type CustomClaim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	JsonPathComponents []string `protobuf:"bytes,2,rep,name=jsonPathComponents,proto3" json:"jsonPathComponents"`
	Type               int32    `protobuf:"varint,3,opt,name=type,proto3,enum=wgpb.ValueType" json:"type"`
	Required           bool     `protobuf:"varint,4,opt,name=required,proto3" json:"required"`
}

type OperationRoleConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequireMatchAll []string `protobuf:"bytes,1,rep,name=requireMatchAll,proto3" json:"requireMatchAll"`
	RequireMatchAny []string `protobuf:"bytes,2,rep,name=requireMatchAny,proto3" json:"requireMatchAny"`
	DenyMatchAll    []string `protobuf:"bytes,3,rep,name=denyMatchAll,proto3" json:"denyMatchAll"`
	DenyMatchAny    []string `protobuf:"bytes,4,rep,name=denyMatchAny,proto3" json:"denyMatchAny"`
}

type S3UploadProfileHooksConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreUpload  bool `protobuf:"varint,1,opt,name=preUpload,proto3" json:"preUpload"`
	PostUpload bool `protobuf:"varint,2,opt,name=postUpload,proto3" json:"postUpload"`
}

type S3UploadProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequireAuthentication     bool                               `protobuf:"varint,1,opt,name=requireAuthentication,proto3" json:"requireAuthentication"`
	MaxAllowedUploadSizeBytes int32                              `protobuf:"varint,2,opt,name=maxAllowedUploadSizeBytes,proto3" json:"maxAllowedUploadSizeBytes"`
	MaxAllowedFiles           int32                              `protobuf:"varint,3,opt,name=maxAllowedFiles,proto3" json:"maxAllowedFiles"`
	AllowedMimeTypes          []string                           `protobuf:"bytes,4,rep,name=allowedMimeTypes,proto3" json:"allowedMimeTypes"`
	AllowedFileExtensions     []string                           `protobuf:"bytes,5,rep,name=allowedFileExtensions,proto3" json:"allowedFileExtensions"`
	MetadataJSONSchema        string                             `protobuf:"bytes,6,opt,name=metadataJSONSchema,proto3" json:"metadataJSONSchema"`
	Hooks                     *S3UploadProfileHooksConfiguration `protobuf:"bytes,7,opt,name=hooks,proto3" json:"hooks"`
}

type S3UploadConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string                      `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Endpoint        *ConfigurationVariable      `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint"`
	AccessKeyID     *ConfigurationVariable      `protobuf:"bytes,3,opt,name=accessKeyID,proto3" json:"accessKeyID"`
	SecretAccessKey *ConfigurationVariable      `protobuf:"bytes,4,opt,name=secretAccessKey,proto3" json:"secretAccessKey"`
	BucketName      *ConfigurationVariable      `protobuf:"bytes,5,opt,name=bucketName,proto3" json:"bucketName"`
	BucketLocation  *ConfigurationVariable      `protobuf:"bytes,6,opt,name=bucketLocation,proto3" json:"bucketLocation"`
	UseSSL          bool                        `protobuf:"varint,7,opt,name=useSSL,proto3" json:"useSSL"`
	UploadProfiles  map[string]*S3UploadProfile `protobuf:"bytes,8,rep,name=uploadProfiles,proto3" json:"uploadProfiles" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

type NodeOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeUrl                      *ConfigurationVariable `protobuf:"bytes,1,opt,name=nodeUrl,proto3" json:"nodeUrl"`
	PublicNodeUrl                *ConfigurationVariable `protobuf:"bytes,4,opt,name=publicNodeUrl,proto3" json:"publicNodeUrl"`
	Listen                       *ListenerOptions       `protobuf:"bytes,2,opt,name=listen,proto3" json:"listen"`
	Logger                       *NodeLogging           `protobuf:"bytes,3,opt,name=logger,proto3" json:"logger"`
	DefaultRequestTimeoutSeconds int64                  `protobuf:"varint,5,opt,name=defaultRequestTimeoutSeconds,proto3" json:"defaultRequestTimeoutSeconds"`
}

type NodeLogging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level *ConfigurationVariable `protobuf:"bytes,1,opt,name=level,proto3" json:"level"`
}

type ListenerOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host *ConfigurationVariable `protobuf:"bytes,1,opt,name=host,proto3" json:"host"`
	Port *ConfigurationVariable `protobuf:"bytes,2,opt,name=port,proto3" json:"port"`
}

type ServerOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerUrl *ConfigurationVariable `protobuf:"bytes,1,opt,name=serverUrl,proto3" json:"serverUrl"`
	Listen    *ListenerOptions       `protobuf:"bytes,2,opt,name=listen,proto3" json:"listen"`
	Logger    *ServerLogging         `protobuf:"bytes,3,opt,name=logger,proto3" json:"logger"`
}

type ServerLogging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level *ConfigurationVariable `protobuf:"bytes,1,opt,name=level,proto3" json:"level"`
}

type CorsConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowedOrigins []*ConfigurationVariable `protobuf:"bytes,1,rep,name=allowedOrigins,proto3" json:"allowedOrigins"`
	AllowedMethods []string                 `protobuf:"bytes,2,rep,name=allowedMethods,proto3" json:"allowedMethods"`

	AllowedHeaders []string `protobuf:"bytes,3,rep,name=allowedHeaders,proto3" json:"allowedHeaders"`

	ExposedHeaders []string `protobuf:"bytes,4,rep,name=exposedHeaders,proto3" json:"exposedHeaders"`

	MaxAge           int64 `protobuf:"varint,5,opt,name=maxAge,proto3" json:"maxAge"`
	AllowCredentials bool  `protobuf:"varint,6,opt,name=allowCredentials,proto3" json:"allowCredentials"`
}
