package cloud

import (
	"fireboom-migrate/types/wgpb"
)

type GlobalOperation struct {
	CacheConfig              *wgpb.OperationCacheConfig                                 `json:"cacheConfig"`
	LiveQueryConfig          *wgpb.OperationLiveQueryConfig                             `json:"liveQueryConfig"`
	AuthenticationConfigs    map[wgpb.OperationType]*wgpb.OperationAuthenticationConfig `json:"authenticationConfigs"`
	ApiAuthenticationHooks   map[string]bool                                            `json:"apiAuthenticationHooks"`
	GlobalHttpTransportHooks map[string]bool                                            `json:"globalHttpTransportHooks"`
}
