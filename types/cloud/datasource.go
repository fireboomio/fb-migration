package cloud

import (
	"fireboom-migrate/types/wgpb"
)

type Datasource struct {
	Name       string `json:"name"`
	Enabled    bool   `json:"enabled"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	DeleteTime string `json:"deleteTime"`

	Kind           wgpb.DataSourceKind  `json:"kind"`
	CustomRest     *wgpb.CustomRest     `json:"customRest"`
	CustomGraphql  *wgpb.CustomGraphql  `json:"customGraphql"`
	CustomDatabase *wgpb.CustomDatabase `json:"customDatabase"`
}
