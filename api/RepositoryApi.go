package api

import (
	"encoding/json"
	"io"

	"github.com/ci-plugins/golang-plugin-sdk/log"
)

// CommitData 代码变更记录详细信息
type CommitData struct {
	Type       int8   `json:"type"`
	PipelineId string `json:"pipelineId"`
	BuildId    string `json:"buildId"`
	Commit     string `json:"commit"`
	Committer  string `json:"committer"`
	CommitTime int64  `json:"commitTime"`
	Comment    string `json:"comment"`
	RepoId     string `json:"repoId"`
	RepoName   string `json:"repoName"`
	ElementId  string `json:"elementId"`
	Url        string `json:"url"`
}

// CommitResult 代码变更记录信息
type CommitResult struct {
	Status int              `json:"status"`
	Data   []CommitResponse `json:"data"`
}

// CommitResponse 代码变更记录信息响应体
type CommitResponse struct {
	Name      string       `json:"name"`
	ElementId string       `json:"elementId"`
	Records   []CommitData `json:"records"`
}

// 获取代码库信息的获取方式，根据仓库ID 或 仓库别名
type repositoryType string

// 代码库选择方式
const (
	RepoTypeId   repositoryType = "ID"
	RepoTypeName repositoryType = "NAME"
)

// GetCommit 获取当前流水线构建下的“代码变更记录”
func GetCommit() (*CommitResult, error) {
	url := buildUrl("/repository/api/build/commit/getCommitsByBuildId")
	headers := getAllHeaders()
	headers["Content-type"] = "application/json"
	build := BuildRequest{path: url, headers: headers, requestBody: nil}
	req, err := buildGet(build)
	if err != nil {
		log.Error("fail to generate request: ", err)
		return nil, err
	}

	respByte, err := request(*req, "fail to get commit history")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info(string(respByte))
	result := new(CommitResult)
	err = json.Unmarshal(respByte, result)
	if err != nil {
		log.Error("fail to resolve response message: ", err)
		return nil, err
	}

	return result, nil
}

// GetRepoInfo 获取指定GIT仓库的信息（包括代码库地址），这里的返回值字段因代码库的类型（@type字段）决定，所以返回值设为map
func GetRepoInfo(repoType repositoryType, repoId string) (map[string]interface{}, error) {
	address := "/repository/api/build/repositories/?repositoryId=" + repoId + "&repositoryType=" + string(repoType)
	result, err := sendGetHttp(address, nil)
	if err != nil {
		log.Error("get git repo info error: " + err.Error())
		return nil, err
	}

	data := result.Data.(map[string]interface{})
	return data, nil
}

// GetGitOauth 获取 git  oauth 信息
func GetGitOauth(userID string) (map[string]interface{}, error) {
	address := "/repository/api/build/oauth/git/" + userID
	result, err := sendGetHttp(address, nil)
	if err != nil {
		log.Error("get git oauth error: " + err.Error())
		return nil, err
	}

	data := result.Data.(map[string]interface{})
	return data, nil
}

func sendGetHttp(address string, requestBody io.Reader) (*Result, error) {
	url := buildUrl(address)
	headers := getAllHeaders()
	headers["Content-type"] = "application/json"
	build := BuildRequest{path: url, headers: headers, requestBody: requestBody}
	req, err := buildGet(build)
	if err != nil {
		log.Error("fail to generate request: ", err)
		return nil, err
	}

	respByte, err := request(*req, "fail to get request info")
	if err != nil {
		log.Error("request fail: ", err)
		return nil, err
	}

	log.Info(string(respByte))
	result := new(Result)
	err = json.Unmarshal(respByte, result)
	if err != nil {
		log.Error("fail to unmarshal response message: ", err)
		return nil, err
	}
	return result, err
}
