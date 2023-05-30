package api

// 版本仓库类型
type artifactType string

// 仓库类型
const (
	Pipeline  artifactType = "PIPELINE"   // 流水线仓库
	CustomDir artifactType = "CUSTOM_DIR" // 自定义仓库
)

// FileChecksums 文件校验值
type FileChecksums struct {
	Sha1 string `json:"sha1"`
	Md5  string `json:"md5"`
}

// FileDetail 构建文件信息
type FileDetail struct {
	Name         string                 `json:"name"`
	FullName     string                 `json:"fullName"`
	FullPath     string                 `json:"fullPath"`
	Size         float64                `json:"size"`
	CreateTime   float64                `json:"createTime"`
	ModifiedTime float64                `json:"modifiedTime"`
	CheckSums    FileChecksums          `json:"checksums"`
	Meta         map[string]interface{} `json:"meta"`
}
