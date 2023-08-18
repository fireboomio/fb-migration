package origin

type Setting struct {
	System      SystemConfig      `json:"system"`      // 系统
	Version     VersionConfig     `json:"version"`     // 版本
	Environment EnvironmentConfig `json:"environment"` // 环境变量
}

type SystemConfig struct {
	ApiPublicAddr       string `json:"apiPublicAddr"`   // API对外暴露地址
	ApiInternalAddr     string `json:"apiInternalAddr"` // wundergraph服务内网地址，用于钩子服务调用api使用
	APIListenHost       string `json:"apiListenHost"`   // wundergraph服务监听地址和端口
	APIListenPort       string `json:"apiListenPort"`
	HooksServerURL      string `json:"hooksServerURL"`      // 钩子服务地址
	HooksServerLanguage string `json:"hooksServerLanguage"` // 钩子使用语言
	LogLevel            int8   `json:"logLevel"`            // 日志水平
	IsDev               bool   `json:"isDev"`               // 开发者模式开关 true开 false关 0-开(debug模式 用up 启动) 1-关(生产模式，用start启动)
	ForcedJumpEnabled   bool   `json:"forcedJumpEnabled"`   // 强制跳转 强制重定向跳转，开启后强制使用https协议
	// TODO 暂时未使用
	DebugEnabled bool `json:"debugEnabled"` // 是否开启调试
	UsageReport  bool `json:"usageReport"`  // 是否上报使用情况
}

type VersionConfig struct {
	VersionNum          string `json:"versionNum"`          // 版本号
	PrismaVersion       string `json:"prismaVersion"`       // prisma版本
	PrismaEngineVersion string `json:"prismaEngineVersion"` // prisma引擎版本
	Copyright           string `json:"copyright"`           // 版权
}

type EnvironmentConfig struct {
	EnvironmentList []EnvironmentDetail `json:"environmentList"` // 环境变量列表
	SystemVariable  string              `json:"systemVariable"`  // 系统变量
}

type EnvironmentDetail struct {
	Name string `json:"name"` // 变量名
	Dev  string `json:"dev"`  // 开发环境
	Pro  string `json:"pro"`  // 生产环境
}
