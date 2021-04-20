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
		FinishBuildWithErrorCode(StatusError, "init atom base param failed", 16015100)
	}

	gAtomBaseParam = new(AtomBaseParam)
	err = LoadInputParam(gAtomBaseParam)
	if err != nil {
		log.Error("init atom base param failed: ", err.Error())
		FinishBuildWithErrorCode(StatusError, "init atom base param failed", 16015100)
	}
}

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
		FinishBuildWithErrorCode(StatusError, "read .sdk.json failed", 16015102)
	}

	gSdkEvn = new(SdkEnv)
	err = json.Unmarshal(data, gSdkEvn)
	if err != nil {
		log.Error("parse .sdk.json failed: ", err.Error())
		FinishBuildWithErrorCode(StatusError, "parse .sdk.json failed", 16015102)
	}

	os.Remove(filePath)
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

func GetOutputData(key string) interface{} {
	return gAtomOutput.Data[key]
}

func AddOutputData(key string, data interface{}) {
	gAtomOutput.Data[key] = data
}

func RemoveOutputData(key string) {
	delete(gAtomOutput.Data, key)
}

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

func SetAtomOutputType(atomOutputType string) {
	gAtomOutput.Type = atomOutputType
}

func GetProjectName() string {
	return gAtomBaseParam.ProjectName
}

func GetProjectDisplayName() string {
	return gAtomBaseParam.ProjectNameCn
}

func GetPipelineId() string {
	return gAtomBaseParam.PipelineId
}

func GetPipelineName() string {
	return gAtomBaseParam.PipelineName
}

func GetPipelineBuildId() string {
	return gAtomBaseParam.PipelineBuildId
}

func GetPipelineBuildNumber() string {
	return gAtomBaseParam.PipelineBuildNum
}

func GetPipelineStartType() string {
	return gAtomBaseParam.PipelineStartType
}

func GetPipelineStartUserId() string {
	return gAtomBaseParam.PipelineStartUserId
}

func GetPipelineStartUserName() string {
	return gAtomBaseParam.PipelineStartUserName
}

func GetPipelineStartTimeMills() string {
	return gAtomBaseParam.PipelineStartTimeMills
}

func GetPipelineVersion() string {
	return gAtomBaseParam.PipelineVersion
}

func GetWorkspace() string {
	return gAtomBaseParam.BkWorkspace
}

func GetTestVersionFlag() string {
	return gAtomBaseParam.TestVersionFlag
}

func GetBkSensitiveConfInfo() map[string]string {
	return gAtomBaseParam.BkSensitiveConfInfo
}

func GetPipelineTaskId() string {
	return gAtomBaseParam.PipelineTaskId
}

func GetPipelineUpdateUserName() string {
	return gAtomBaseParam.PipelineUpdateUserName
}
