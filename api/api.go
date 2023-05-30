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

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ci-plugins/golang-plugin-sdk/log"
)

var gSdkEvn *SdkEnv
var gAtomBaseParam *AtomBaseParam
var gAllAtomParam map[string]interface{}
var gAtomOutput *AtomOutput

var gDataDir string
var gInputFile string
var gOutputFile string

// StringResult 蓝盾后台返回结果
type StringResult struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func init() {
	gAtomOutput = NewAtomOutput()
	gDataDir = getDataDir()
	gInputFile = getInputFile()
	gOutputFile = getOutputFile()
	initSdkEnv()
	initAtomParam()
}

func initAtomParam() {
	err := LoadInputParam(&gAllAtomParam)
	if err != nil {
		log.Error("init atom base param failed: ", err.Error())
		FinishBuildWithError(StatusError, "init atom base param failed", 2189503, PluginError)
	}

	gAtomBaseParam = new(AtomBaseParam)
	err = LoadInputParam(gAtomBaseParam)
	postActionParam := flag.String("postAction", "noPostAction", "后置动作")
	flag.Parse()
	gAtomBaseParam.PostActionParam = *postActionParam
	if err != nil {
		log.Error("init atom base param failed: ", err.Error())
		FinishBuildWithError(StatusError, "init atom base param failed", 2189503, PluginError)
	}
}

// GetInputParam 获取输入参数
// @name	参数名称
func GetInputParam(name string) string {
	value := gAllAtomParam[name]
	if value == nil {
		return ""
	}
	strValue, ok := value.(string)
	if !ok {
		return ""
	}
	return strValue
}

// LoadInputParam 加载输入参数
func LoadInputParam(v interface{}) error {
	file := gDataDir + "/" + gInputFile
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Error("load input param failed:", err.Error())
		return errors.New("load input param failed")
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		log.Error("parse input param failed:", err.Error())
		return errors.New("parse input param failed")
	}
	return nil
}

func initSdkEnv() {
	filePath := gDataDir + "/.sdk.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error("read .sdk.json failed: ", err.Error())
		FinishBuildWithError(StatusError, "read .sdk.json failed", 2189503, PluginError)
	}

	gSdkEvn = new(SdkEnv)
	err = json.Unmarshal(data, gSdkEvn)
	if err != nil {
		log.Error("parse .sdk.json failed: ", err.Error())
		FinishBuildWithError(StatusError, "read .sdk.json failed", 2189503, PluginError)
	}

	// os.Remove(filePath)
}

func getDataDir() string {
	dir := strings.TrimSpace(os.Getenv(DataDirEnv))
	if len(dir) == 0 {
		dir, _ = os.Getwd()
	}
	return dir
}

func getInputFile() string {
	file := strings.TrimSpace(os.Getenv(InputFileEnv))
	if len(file) == 0 {
		file = "input.json"
	}
	return file
}

func getOutputFile() string {
	file := strings.TrimSpace(os.Getenv(OutputFileEnv))
	if len(file) == 0 {
		file = "output.json"
	}
	return file
}

// GetOutputData 获取输出参数
func GetOutputData(key string) interface{} {
	return gAtomOutput.Data[key]
}

// AddOutputData 添加输出参数
func AddOutputData(key string, data interface{}) {
	gAtomOutput.Data[key] = data
}

// RemoveOutputData 删除输出参数
func RemoveOutputData(key string) {
	delete(gAtomOutput.Data, key)
}

// GetQualityData 获取质量红线信息
func GetQualityData(qualityKey string) interface{} {
	return gAtomOutput.QualityData[qualityKey]
}

// AddQualityData 添加质量红线信息
func AddQualityData(qualityKey string, qualitydata *Qualitydata) {
	gAtomOutput.Type = "quality"
	gAtomOutput.QualityData[qualityKey] = qualitydata
}

// RemoveQualityData 删除质量红线信息
func RemoveQualityData(qualityKey string) {
	delete(gAtomOutput.QualityData, qualityKey)
}

// SetPlatformCode 设置插件对接平台代码
func SetPlatformCode(platformCode string) {
	gAtomOutput.PlatformCode = platformCode
}

// SetPlatformErrorCode 设置插件对接平台错误码
func SetPlatformErrorCode(platformErrorCode int) {
	gAtomOutput.PlatformErrorCode = platformErrorCode
}

// WriteOutput 将输出写到文件
func WriteOutput() error {
	data, _ := json.Marshal(gAtomOutput)

	file := gDataDir + "/" + gOutputFile
	err := ioutil.WriteFile(file, data, 0644)
	if err != nil {
		log.Error("write output failed: ", err.Error())
		return errors.New("write output failed")
	}
	return nil
}

// FinishBuild 结束构建
func FinishBuild(status Status, msg string) {
	gAtomOutput.Message = msg
	gAtomOutput.Status = status
	WriteOutput()
	switch status {
	case StatusSuccess:
		os.Exit(0)
	case StatusFailure:
		os.Exit(1)
	case StatusError:
		os.Exit(2)
	default:
		os.Exit(0)
	}
}

// FinishBuildWithErrorCode 结束构建
// @status		任务状态
// @msg			消息
// @errorCode	错误码
func FinishBuildWithErrorCode(status Status, msg string, errorCode int) {
	gAtomOutput.Message = msg
	gAtomOutput.Status = status
	gAtomOutput.ErrorCode = errorCode
	WriteOutput()
	switch status {
	case StatusSuccess:
		os.Exit(0)
	case StatusFailure:
		os.Exit(1)
	case StatusError:
		os.Exit(2)
	default:
		os.Exit(0)
	}
}

// FinishBuildWithError 结束构建
// @status		任务状态
// @msg			消息
// @errorCode	错误码
// @errorType	错误类型
func FinishBuildWithError(status Status, msg string, errorCode int, errorType ErrorType) {
	gAtomOutput.Message = msg
	gAtomOutput.Status = status
	gAtomOutput.ErrorCode = errorCode
	gAtomOutput.ErrorType = errorType
	WriteOutput()
	switch status {
	case StatusSuccess:
		os.Exit(0)
	case StatusFailure:
		os.Exit(1)
	case StatusError:
		os.Exit(2)
	default:
		os.Exit(0)
	}
}

// SetAtomOutputType 获取插件输出类型
func SetAtomOutputType(atomOutputType string) {
	gAtomOutput.Type = atomOutputType
}

// GetProjectName 获取项目名称
func GetProjectName() string {
	return gAtomBaseParam.ProjectName
}

// GetProjectDisplayName 获取项目显示名称
func GetProjectDisplayName() string {
	return gAtomBaseParam.ProjectNameCn
}

// GetPipelineId 获取流水线ID
func GetPipelineId() string {
	return gAtomBaseParam.PipelineId
}

// GetPipelineName 获取流水线名称
func GetPipelineName() string {
	return gAtomBaseParam.PipelineName
}

// GetPipelineBuildId 获取构建ID
func GetPipelineBuildId() string {
	return gAtomBaseParam.PipelineBuildId
}

// GetPipelineBuildNumber 获取构建号
func GetPipelineBuildNumber() string {
	return gAtomBaseParam.PipelineBuildNum
}

// GetPipelineStartType 获取流水线启动方式
func GetPipelineStartType() string {
	return gAtomBaseParam.PipelineStartType
}

// GetPipelineStartUserId 获取流水线启动用户ID
func GetPipelineStartUserId() string {
	return gAtomBaseParam.PipelineStartUserId
}

// GetPipelineStartUserName 获取流水线启动用户名
func GetPipelineStartUserName() string {
	return gAtomBaseParam.PipelineStartUserName
}

// GetPipelineStartTimeMills 获取流水线启动时间
func GetPipelineStartTimeMills() string {
	return gAtomBaseParam.PipelineStartTimeMills
}

// GetPipelineVersion 获取流水线版本号
func GetPipelineVersion() string {
	return gAtomBaseParam.PipelineVersion
}

// GetWorkspace 获取工作目录
func GetWorkspace() string {
	if gAtomBaseParam.BkWorkspace == "" {
		return "."
	}
	return gAtomBaseParam.BkWorkspace
}

// GetPipelineCreateUser 获取流水线创建人
func GetPipelineCreateUser() string {
	return gAtomBaseParam.PipelineCreateUser
}

// GetPipelineModifyUser 获取流水线最后修改人
func GetPipelineModifyUser() string {
	return gAtomBaseParam.PipelineModifyUser
}

// GetSensitiveConfParam 获取插件敏感参数
func GetSensitiveConfParam(fieldName string) string {
	return gAtomBaseParam.BkSensitiveConfInfo[fieldName]
}

// GetPostActionParam 获取后置执行参数
func GetPostActionParam() string {
	return gAtomBaseParam.PostActionParam
}

// GetBuildVarByKey 获取指定构建下的构建参数
func GetBuildVarByKey(key string) (string, error) {
	if key == "" {
		return "", errors.New("key is empty")
	}

	vars, err := GetBuildVar()
	if err != nil {
		return "", err
	}

	value, ok := vars[key].(string)
	if !ok {
		log.Error("key ", key, " is not exist in map")
		return "", nil
	}
	return value, nil
}

// GetBuildVar 获取指定构建下的构建参数
func GetBuildVar() (map[string]interface{}, error) {
	url := buildUrl("process/api/build/variable/getBuildVariable")
	headers := getAllHeaders()
	headers["X-DEVOPS-BUILD-ID"] = gAtomBaseParam.PipelineBuildId
	headers["X-DEVOPS-PROJECT-ID"] = gAtomBaseParam.ProjectName
	headers["X-DEVOPS-PIPELINE-ID"] = gAtomBaseParam.PipelineId
	build := BuildRequest{path: url, headers: headers, requestBody: nil}
	req, err := buildGet(build)
	if err != nil {
		log.Error("fail to generate request: ", err)
		return nil, err
	}

	respByte, err := request(*req, "fail to get build variable")
	if err != nil {
		return nil, err
	}

	log.Info(string(respByte))
	result := new(Result)
	err = json.Unmarshal(respByte, result)
	if err != nil {
		log.Error("fail to unmarshal response message: ", err)
		return nil, err
	}

	buildVar := result.Data.(map[string]interface{})
	return buildVar, nil
}

// GetBuildContextByKey 获取指定构建下的构建上下文
// Deprecated: 该接口将在后期弃用,请改用 GetVariableByName.
func GetBuildContextByKey(key string) string {
	return getVariable(key, false)
}

// GetVariableByName 获取指定构建下的构建上下文新版
func GetVariableByName(name string) string {
	return getVariable(name, true)
}

func getVariable(name string, check bool) string {
	url := buildUrl("process/api/build/variable/get_build_context?contextName=" + name + "&check=" + fmt.Sprintf("%t", check))
	headers := getAllHeaders()
	headers["X-DEVOPS-BUILD-ID"] = gAtomBaseParam.PipelineBuildId
	headers["X-DEVOPS-PROJECT-ID"] = gAtomBaseParam.ProjectName
	headers["X-DEVOPS-PIPELINE-ID"] = gAtomBaseParam.PipelineId
	build := BuildRequest{path: url, headers: headers, requestBody: nil}
	req, err := buildGet(build)
	if err != nil {
		log.Error("fail to generate request: ", err)
		return ""
	}

	respByte, err := request(*req, "fail to get build context")
	if err != nil {
		return ""
	}

	log.Info(string(respByte))
	result := new(Result)
	err = json.Unmarshal(respByte, result)
	if err != nil {
		log.Error("fail to unmarshal response message: ", err)
		return ""
	}
	respStr := result.Data.(string)
	return respStr
}
