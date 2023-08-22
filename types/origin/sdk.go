package origin

type FbSDK struct {
	Id          int64  `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Name        string `db:"name" json:"name"`
	Author      string `db:"author" json:"author"`
	Version     string `db:"version" json:"version"`
	Description string `db:"description" json:"description"`
	Type        string `db:"type" json:"type"`
	Language    string `db:"language" json:"language"`
	Url         string `db:"url" json:"url"`
	OutputPath  string `db:"outputPath" json:"outputPath"`
	DirName     string `db:"dirName" json:"dirName"`
	Icon        string `db:"icon" json:"icon"`
	Enabled     bool   `db:"enabled" json:"enabled"`
	CreateTime  string `db:"create_time" json:"createTime"`
	UpdateTime  string `db:"update_time" json:"updateTime"`
	DeleteTime  string `db:"delete_time" json:"deleteTime"`
}
