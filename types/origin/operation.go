package origin

type FbOperation struct {
	Id            int64  `db:"id" json:"id"`
	Method        string `db:"method" json:"method"`                // 方法类型 GET、POST、PUT、DELETE
	RestURL       string `db:"rest_url" json:"restUrl"`             // 方法类型 GET、POST、PUT、DELETE
	OperationType string `db:"operation_type" json:"operationType"` // 请求类型 queries,mutations,subscriptions
	IsPublic      bool   `db:"is_public" json:"isPublic"`           // 状态 true 公有 false 私有
	Remark        string `db:"remark" json:"remark"`                // 说明
	Illegal       bool   `db:"illegal" json:"illegal"`              // 是否非法 true 非法 false 合法
	LiveQuery     bool   `db:"live_query" json:"liveQuery"`         // 是否实时 true 是 false 否
	Path          string `db:"path" json:"title"`                   // 路径
	Content       string `db:"content" json:"content"`              // 内容
	Enabled       bool   `db:"enabled" json:"enabled"`              // 开关 true开 false关
	CreateTime    string `db:"create_time" json:"createTime"`
	UpdateTime    string `db:"update_time" json:"updateTime"`
	DeleteTime    string `db:"delete_time" json:"deleteTime"`
	RoleType      string `db:"role_type" json:"roleType"`
	Roles         string `db:"roles" json:"roles"`
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
