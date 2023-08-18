package origin

type FbRole struct {
	Id         int64  `db:"id" json:"id"`
	Code       string `db:"code" json:"code"`     // 角色编码
	Remark     string `db:"remark" json:"remark"` // 描述
	CreateTime string `db:"create_time" json:"createTime"`
	UpdateTime string `db:"update_time" json:"updateTime"`
	DeleteTime string `db:"delete_time" json:"deleteTime"`
}
