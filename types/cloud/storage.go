package cloud

import (
	"fireboom-migrate/types/wgpb"
)

type Storage struct {
	Enabled    bool   `json:"enabled"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	DeleteTime string `json:"deleteTime"`

	wgpb.S3UploadConfiguration
}
