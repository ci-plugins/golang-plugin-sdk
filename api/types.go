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

const (
	AuthHeaderBuildId              = "X-SODA-BID"
	AuthHeaderProjectId            = "X-SODA-PID"
	AuthHeaderDevopsBuildType      = "X-DEVOPS-BUILD-TYPE"
	AuthHeaderDevopsProjectId      = "X-DEVOPS-PROJECT-ID"
	AuthHeaderDevopsBuildId        = "X-DEVOPS-BUILD-ID"
	AuthHeaderDevopsVmSeqId        = "X-DEVOPS-VM-SID"
	AuthHeaderDevopsAgentId        = "X-DEVOPS-AGENT-ID"
	AuthHeaderDevopsAgentSecretKey = "X-DEVOPS-AGENT-SECRET-KEY"
)

const (
	DataDirEnv    = "bk_data_dir"
	InputFileEnv  = "bk_data_input"
	OutputFileEnv = "bk_data_output"
)

type SdkEnv struct {
	BuildType string `json:"buildType"`
	ProjectId string `json:"projectId"`
	AgentId   string `json:"agentId"`
	SecretKey string `json:"secretKey"`
	Gateway   string `json:"gateway"`
	BuildId   string `json:"buildId"`
	VmSeqId   string `json:"vmSeqId"`
}

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
	TestVersionFlag        string            `json:"testVersionFlag"`
	BkSensitiveConfInfo    map[string]string `json:"bkSensitiveConfInfo"`
	PipelineTaskId         string            `json:"BK_CI_BUILD_TASK_ID"`
	PipelineUpdateUserName string            `json:"BK_CI_PIPELINE_UPDATE_USER"`
}

type BuildType string

const (
	BuildTypeWorker      = "WORKER"
	BuildTypeAgent       = "AGENT"
	BuildTypePluginAgent = "PLUGIN_AGENT"
	BuildTypeDocker      = "DOCKER"
	BuildTypeDockerHost  = "DOCKER_HOST"
	BuildTypeTstackAgent = "TSTACK_AGENT"
)

type DataType string

const (
	DataTypeString   DataType = "string"
	DataTypeArtifact DataType = "artifact"
	DataTypeReport   DataType = "report"
)

type Status string

const (
	StatusSuccess Status = "success"
	StatusFailure Status = "failure"
	StatusError   Status = "error"
)

type ArtifactData struct {
	Type  DataType `json:"type"`
	Value []string `json:"value"`
}

func (a *ArtifactData) AddArtifact(artifact string) {
	a.Value = append(a.Value, artifact)
}

func (a *ArtifactData) AddArtifactAll(artifacts []string) {
	a.Value = append(a.Value, artifacts...)
}

type StringData struct {
	Type  DataType `json:"type"`
	Value string   `json:"value"`
}

type ReportData struct {
	Type   DataType `json:"type"`
	Label  string   `json:"label"`
	Path   string   `json:"path"`
	Target string   `json:"target"`
}

func NewReportData(label string, path string, target string) *ReportData {
	return &ReportData{
		Type:   DataTypeReport,
		Label:  label,
		Path:   path,
		Target: target,
	}
}

func NewStringData(value string) *StringData {
	return &StringData{
		Type:  DataTypeString,
		Value: value,
	}
}

func NewArtifactData() *ArtifactData {
	return &ArtifactData{
		Type:  DataTypeArtifact,
		Value: []string{},
	}
}

type AtomOutput struct {
	Status    Status                 `json:"status"`
	Message   string                 `json:"message"`
	ErrorCode int                    `json:"errorCode"`
	Type      string                 `json:"type"`
	Data      map[string]interface{} `json:"data"`
}

func NewAtomOutput() *AtomOutput {
	output := new(AtomOutput)
	output.Status = StatusSuccess
	output.Message = "success"
	output.Type = "default"
	output.Data = make(map[string]interface{})
	return output
}
