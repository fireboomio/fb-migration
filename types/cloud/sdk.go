package cloud

type sdkType string

const (
	SdkServer, SdkClient sdkType = "server", "client"
	fieldEnabled                 = "enabled"
	fieldOutputPath              = "outputPath"
	fieldGitpull                 = "gitpull"
)

type Sdk struct {
	Name       string `json:"name"`
	Enabled    bool   `json:"enabled"`
	Type       string `json:"type"`
	Language   string `json:"language"`
	Extension  string `json:"extension"`
	GitUrl     string `json:"gitUrl"`
	GitBranch  string `json:"gitBranch"`
	OutputPath string `json:"outputPath"`

	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
	DeleteTime  string `json:"deleteTime"`
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Version     string `json:"version"`
	Description string `json:"description"`
}
