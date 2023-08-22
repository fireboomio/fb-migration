package cloud

import (
	"fireboom-migrate/types/wgpb"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

type GlobalSetting struct {
	NodeOptions       *wgpb.NodeOptions       `json:"nodeOptions"`       // WunderGraphConfiguration.UserDefinedApi.NodeOptions
	ServerOptions     *wgpb.ServerOptions     `json:"serverOptions"`     // WunderGraphConfiguration.UserDefinedApi.ServerOptions
	CorsConfiguration *wgpb.CorsConfiguration `json:"corsConfiguration"` // WunderGraphConfiguration.UserDefinedApi.Cors
	SecurityConfig

	Appearance    *Appearance       `json:"appearance"`
	ConsoleLogger *lumberjackLogger `json:"consoleLogger"`
	BuildInfo     *BuildInfo        `json:"buildInfo"`
}

type SecurityConfig struct {
	AllowedReport         bool                          `json:"allowedReport"`
	EnableGraphqlEndpoint bool                          `json:"enableGraphqlEndpoint"` // WunderGraphConfiguration.EnableGraphqlEndpoint
	AllowedHostNames      []*wgpb.ConfigurationVariable `json:"allowedHostNames"`      // WunderGraphConfiguration.AllowedHostNames

	AuthorizedRedirectUris []*wgpb.ConfigurationVariable `json:"authorizedRedirectUris"`
	AuthenticationKey      string                        `json:"authenticationKey"`
	EnableCSRFProtect      bool                          `json:"enableCSRFProtect"`
	ForceHttpsRedirects    bool                          `json:"forceHttpsRedirects"`
	GlobalRateLimit        struct {
		Enabled     bool          `json:"enabled"`
		Requests    int           `json:"requests"`
		PerDuration time.Duration `json:"perDuration"`
	} `json:"globalRateLimit"`
}

type Appearance struct {
	Language string `json:"language"`
}

type lumberjackLogger lumberjack.Logger

type BuildInfo struct {
	Version, Commit, Date, BuiltBy string
}
