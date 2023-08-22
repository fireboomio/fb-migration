package origin

type GlobalConfig struct {
	// 是否强制https重定向
	ForceHttpsRedirects             bool                               `json:"forceHttpsRedirects"`
	AuthRedirectURL                 []string                           `json:"authorizedRedirectUris"` // 身份鉴权-重定向url配置
	ConfigureWunderGraphApplication WunderGraphConfigApplicationConfig `json:"configureWunderGraphApplication"`
}

type WunderGraphConfigApplicationConfig struct {
	Security *SecurityConfig    `json:"security"`
	Cors     *CorsConfiguration `json:"cors"`
}

type SecurityConfig struct {
	AllowedHostsEnabled   bool     `json:"allowedHostsEnabled"`
	EnableGraphQLEndpoint bool     `json:"enableGraphQLEndpoint"` // GraphQL端点 0-关 1-开
	AllowedHosts          []string `json:"allowedHosts"`          // 允许主机,多个域名
	// 开启 csrf 保护
	EnableCSRFProtect bool `json:"enableCSRF"`
}

type CorsConfiguration struct {
	AllowedOriginsEnabled bool     `json:"allowedOriginsEnabled"` // 允许域名开关
	AllowedOrigins        []string `json:"allowedOrigins"`        // 允许域名
	AllowedMethods        []string `json:"allowedMethods"`        // 允许方法 0-* 1-GET 2-POST 3-PUT
	AllowedHeadersEnabled bool     `json:"allowedHeadersEnabled"` // 允许投开关
	AllowedHeaders        []string `json:"allowedHeaders"`        // 请求头部
	ExposedHeaders        []string `json:"exposedHeaders"`        // 排除头部
	ExposedHeadersEnabled bool     `json:"exposedHeadersEnabled"` // 排除头部开关
	MaxAge                int64    `json:"maxAge"`                // 跨域时间(s)
	AllowCredentials      bool     `json:"allowCredentials"`      // 允许证书开关 0-开 1-关
}
