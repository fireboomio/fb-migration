package origin

type FbDataSource struct {
	Id         int64       `db:"id" json:"id"`
	Name       string      `db:"name" json:"name"`              // 数据源名称
	SourceType int64       `db:"source_type" json:"sourceType"` // 数据源类型: 1-db 2-rest 3-graphql 4-自定义
	Config     interface{} `db:"config" json:"config"`          // 数据源对应的配置项：命名空间、请求配置、连接配置、文件路径、是否配置为外部数据源等
	Enabled    bool        `db:"enabled" json:"enabled"`
	CreateTime string      `db:"create_time" json:"createTime"`
	UpdateTime string      `db:"update_time" json:"updateTime"`
	DeleteTime string      `db:"delete_time" json:"deleteTime"`
}
