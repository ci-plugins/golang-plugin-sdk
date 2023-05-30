/*
 * Tencent is pleased to support the open source community by making BK-CI 蓝鲸持续集成平台 available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company.  All rights reserved.
 *
 * BK-CI 蓝鲸持续集成平台 is licensed under the MIT license.
 *
 * A copy of the MIT License is included in this file.
 *
 *
 * Terms of the MIT License:
 * ---------------------------------------------------
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation
 * files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy,
 * modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT
 * LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
 * NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
 * WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
 * SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package api

// http head keys
const (
	AuthHeaderBuildId              = "X-SODA-BID"
	AuthHeaderProjectId            = "X-SODA-PID"
	AuthHeaderDevopsBuildType      = "X-DEVOPS-BUILD-TYPE"
	AuthHeaderDevopsProjectId      = "X-DEVOPS-PROJECT-ID"
	AuthHeaderDevopsBuildId        = "X-DEVOPS-BUILD-ID"
	AuthHeaderDevopsVmSeqId        = "X-DEVOPS-VM-SID"
	AuthHeaderDevopsVmSeqName      = "X-DEVOPS-VM-NAME"
	AuthHeaderDevopsAgentId        = "X-DEVOPS-AGENT-ID"
	AuthHeaderDevopsAgentSecretKey = "X-DEVOPS-AGENT-SECRET-KEY"
	AuthHeaderDevopsCiTaskId       = "X-DEVOPS-CI-TASK-ID"
)

// consts
const (
	DataDirEnv    = "bk_data_dir"
	InputFileEnv  = "bk_data_input"
	OutputFileEnv = "bk_data_output"
)

// SdkEnv 插件运行环境变量
type SdkEnv struct {
	BuildType string `json:"buildType"`
	ProjectId string `json:"projectId"`
	AgentId   string `json:"agentId"`
	SecretKey string `json:"secretKey"`
	Gateway   string `json:"gateway"`
	BuildId   string `json:"buildId"`
	VmSeqId   string `json:"vmSeqId"`
	TaskId    string `json:"taskId"`
}

// AtomBaseParam 插件基本参数
type AtomBaseParam struct {
	PipelineVersion        string            `json:"BK_CI_PIPELINE_VERSION"`
	ProjectName            string            `json:"BK_CI_PROJECT_NAME"`
	ProjectNameCn          string            `json:"BK_CI_PROJECT_NAME_CN"`
	PipelineId             string            `json:"BK_CI_PIPELINE_ID"`
	PipelineBuildNum       string            `json:"BK_CI_BUILD_NUM"`
	PipelineBuildId        string            `json:"BK_CI_BUILD_ID"`
	PipelineName           string            `json:"BK_CI_PIPELINE_NAME"`
	PipelineStartTimeMills string            `json:"BK_CI_BUILD_START_TIME"`
	PipelineStartType      string            `json:"BK_CI_START_TYPE"`
	PipelineStartUserId    string            `json:"BK_CI_START_USER_ID"`
	PipelineStartUserName  string            `json:"BK_CI_START_USER_NAME"`
	BkWorkspace            string            `json:"bkWorkspace"`
	PipelineCreateUser     string            `json:"BK_CI_PIPELINE_CREATE_USER"`
	PipelineModifyUser     string            `json:"BK_CI_PIPELINE_UPDATE_USER"`
	BkSensitiveConfInfo    map[string]string `json:"bkSensitiveConfInfo"`
	PostActionParam        string            `json:"postEntryParam"`
	TestVersionFlag        string            `json:"testVersionFlag"`
	TaskId                 string            `json:"BK_CI_BUILD_TASK_ID"`
	TaskAtomCode           string            `json:"BK_CI_ATOM_CODE"`
	TaskAtomName           string            `json:"BK_CI_ATOM_NAME"`
	TaskAtomVersion        string            `json:"BK_CI_ATOM_VERSION"`
	TaskName               string            `json:"BK_CI_TASK_NAME"`
	StepId                 string            `json:"BK_CI_STEP_ID"`
}

// BuildType 构建类型
type BuildType string

// 构建类型
const (
	BuildTypeWorker      = "WORKER"
	BuildTypeAgent       = "AGENT"
	BuildTypePluginAgent = "PLUGIN_AGENT"
	BuildTypeDocker      = "DOCKER"
	BuildTypeDockerHost  = "DOCKER_HOST"
	BuildTypeTstackAgent = "TSTACK_AGENT"
)

// DataType 输出数据类型
type DataType string

// 输出数据类型
const (
	DataTypeString   DataType = "string"
	DataTypeArtifact DataType = "artifact"
	DataTypeReport   DataType = "report"
)

// Status 插件执行状态
type Status string

// 插件执行状态
const (
	StatusSuccess Status = "success"
	StatusFailure Status = "failure"
	StatusError   Status = "error"
)

// ErrorType 执行错误类型
type ErrorType int

// 插件执行状态
const (
	UserError       ErrorType = 1
	ThirdPartyError ErrorType = 2
	PluginError     ErrorType = 3
)

// ReportType 报告类型
type ReportType string

// 报告类型
const (
	ReportTypeInternal   ReportType = "INTERNAL"   // 内部报告，文件
	ReportTypeThirdparty ReportType = "THIRDPARTY" // 外部报告，URL连接
)

// ArtifactData 构件输出数据
type ArtifactData struct {
	Type         DataType     `json:"type"`
	Value        []string     `json:"value"`
	ArtifactType artifactType `json:"artifactoryType"`
	Path         string       `json:"path"`
}

// AddArtifact 添加待归档构件输出
func (a *ArtifactData) AddArtifact(artifact string) {
	a.Value = append(a.Value, artifact)
	a.ArtifactType = Pipeline
}

// AddArtifactAll 批量添加待归档构件输出
func (a *ArtifactData) AddArtifactAll(artifacts []string) {
	a.Value = append(a.Value, artifacts...)
	a.ArtifactType = Pipeline
}

// AddArtifactToCustomRepo 添加待归档构件输出至自定义仓库
func (a *ArtifactData) AddArtifactToCustomRepo(artifact string, customPath string) {
	a.Value = append(a.Value, artifact)
	a.ArtifactType = CustomDir
	a.Path = customPath
}

// AddArtifactAllToCustomRepo 批量添加待归档构件输出至自定义仓库
func (a *ArtifactData) AddArtifactAllToCustomRepo(artifacts []string, customPath string) {
	a.Value = append(a.Value, artifacts...)
	a.ArtifactType = CustomDir
	a.Path = customPath
}

// StringData 变量输出数据
type StringData struct {
	Type  DataType `json:"type"`
	Value string   `json:"value"`
}

// ReportData 报告输出数据
type ReportData struct {
	Type       DataType   `json:"type"`
	Label      string     `json:"label"`
	Path       string     `json:"path"`
	Url        string     `json:"url"`
	Target     string     `json:"target"`
	ReportType ReportType `json:"reportType"`
}

// NewReportData 添加报告输出
func NewReportData(label string, path string, target string) *ReportData {
	return NewInternalReportData(label, path, target)
}

// NewInternalReportData 添加内部报告输出
func NewInternalReportData(label string, path string, target string) *ReportData {
	return &ReportData{
		Type:       DataTypeReport,
		Label:      label,
		Path:       path,
		Target:     target,
		ReportType: ReportTypeInternal,
	}
}

// NewThirdpartyReportData 添加第三方报告输出
func NewThirdpartyReportData(label string, url string) *ReportData {
	return &ReportData{
		Type:       DataTypeReport,
		Label:      label,
		Url:        url,
		ReportType: ReportTypeThirdparty,
	}
}

// NewStringData 添加变量输出
func NewStringData(value string) *StringData {
	return &StringData{
		Type:  DataTypeString,
		Value: value,
	}
}

// NewArtifactData 创建构件输出数据
func NewArtifactData() *ArtifactData {
	return &ArtifactData{
		Type:  DataTypeArtifact,
		Value: []string{},
	}
}

// Qualitydata 质量红线数据
type Qualitydata struct {
	Value string `json:"value"`
}

// String 质量红线数据
func (a *Qualitydata) String() string {
	return a.Value
}

// NewQualityData 创建质量红线数据
func NewQualityData(value string) *Qualitydata {
	return &Qualitydata{
		Value: value,
	}
}

// AtomOutput 插件输出
type AtomOutput struct {
	Status            Status                  `json:"status"`
	Message           string                  `json:"message"`
	ErrorCode         int                     `json:"errorCode"`
	ErrorType         ErrorType               `json:"errorType"`
	Type              string                  `json:"type"`
	Data              map[string]interface{}  `json:"data"`
	QualityData       map[string]*Qualitydata `json:"qualityData"`
	PlatformCode      string                  `json:"platformCode"`
	PlatformErrorCode int                     `json:"platformErrorCode"`
}

// NewAtomOutput 创建插件输出
func NewAtomOutput() *AtomOutput {
	output := new(AtomOutput)
	output.Status = StatusSuccess
	output.Message = "success"
	output.Type = "default"
	output.Data = make(map[string]interface{})
	output.QualityData = make(map[string]*Qualitydata)
	return output
}
