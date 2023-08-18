package origin

import (
	"fireboom-migrate/types/wgpb"
)

type S3Upload struct {
	Name            string                           `json:"name"`
	EndPoint        string                           `json:"endPoint"`
	BucketName      string                           `json:"bucketName"`
	BucketLocation  string                           `json:"bucketLocation"`
	AccessKeyID     Value                            `json:"accessKeyID"`
	SecretAccessKey Value                            `json:"secretAccessKey"`
	UseSSL          bool                             `json:"useSSL"`
	UploadProfiles  map[string]*wgpb.S3UploadProfile `json:"uploadProfiles"`
}

type Value struct {
	Key  string `json:"key"`
	Kind string `json:"kind"` // 0-值 1-环境变量 2-转发值客户端
	Val  string `json:"val"`
}

// FbStorageBucket  存储配置
type FbStorageBucket struct {
	Id         int64    `db:"id" json:"id"`
	Name       string   `db:"name" json:"name"`       //  名称
	Enabled    bool     `db:"enabled" json:"enabled"` //  启用
	Config     S3Upload `db:"config" json:"config"`   //  配置
	CreateTime string   `db:"create_time" json:"createTime"`
	UpdateTime string   `db:"update_time" json:"updateTime"`
	DeleteTime string   `db:"delete_time" json:"deleteTime"`
}
