package origin

type HookStruct struct {
	Path    string `json:"path"` // 钩子的路径 钩子分四种operation、auth、customize、global 以这4个为一级目录，二级目录为钩子的类型(operation的则以operation的名称为二级目录，钩子类型为三级目录)
	Enabled *bool  `json:"enabled"`
}
