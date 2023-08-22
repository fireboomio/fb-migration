package cloud

import (
	"fireboom-migrate/types/wgpb"
)

type Operation struct {
	Enabled    bool   `json:"enabled"`
	Title      string `json:"title"`
	Remark     string `json:"remark"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	DeleteTime string `json:"deleteTime"`

	Path               string                            `json:"path"`
	Engine             wgpb.OperationExecutionEngine     `json:"engine"`
	OperationType      wgpb.OperationType                `json:"operationType"`
	HooksConfiguration *wgpb.OperationHooksConfiguration `json:"hooksConfiguration"`

	ConfigCustomized     bool                                `json:"configCustomized"`
	CacheConfig          *wgpb.OperationCacheConfig          `json:"cacheConfig"`
	LiveQueryConfig      *wgpb.OperationLiveQueryConfig      `json:"liveQueryConfig"`
	AuthenticationConfig *wgpb.OperationAuthenticationConfig `json:"authenticationConfig"`

	Invalid             bool                               `json:"-"`
	Internal            bool                               `json:"-"`
	AuthorizationConfig *wgpb.OperationAuthorizationConfig `json:"-"`
}

type OperationSetting struct {
	AuthenticationQueriesRequired       bool  `json:"authenticationQueriesRequired"`       // 查询身份验证授权
	AuthenticationMutationsRequired     bool  `json:"authenticationMutationsRequired"`     // 更改身份验证授权
	AuthenticationSubscriptionsRequired bool  `json:"authenticationSubscriptionsRequired"` // 推送身份验证授权
	Enabled                             bool  `json:"enabled"`                             // 设置开关
	AuthenticationRequired              bool  `json:"authenticationRequired"`              // 需要授权
	CachingEnabled                      bool  `json:"cachingEnabled"`                      // 开启缓存
	CachingMaxAge                       int64 `json:"cachingMaxAge"`                       // 最大时长
	CachingStaleWhileRevalidate         int64 `json:"cachingStaleWhileRevalidate"`         // 重校验时长
	LiveQueryEnabled                    bool  `json:"liveQueryEnabled"`                    // 开启实时
	LiveQueryPollingIntervalSeconds     int64 `json:"liveQueryPollingIntervalSeconds"`     // 轮询间隔
}
