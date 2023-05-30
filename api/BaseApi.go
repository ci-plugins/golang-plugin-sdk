package api

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ci-plugins/golang-plugin-sdk/log"
)

// Result 插件执行结果
type Result struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// BuildRequest 构建请求
type BuildRequest struct {
	path        string
	headers     map[string]string
	requestBody io.Reader
}

var client = http.Client{
	Timeout: 30 * time.Second,
}

func request(r http.Request, errMessage string) ([]byte, error) {
	response, err := client.Do(&r)
	if err != nil {
		log.Error("do http request failed: " + err.Error())
		return nil, errors.New(errMessage)
	}

	if !(response.StatusCode >= 200 && response.StatusCode < 300) {
		log.Error("http request failed, status: " + strconv.Itoa(response.StatusCode))
		return nil, errors.New(errMessage)
	}

	respStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error("get response content failed: " + err.Error())
		return nil, errors.New(errMessage)
	}

	return respStr, nil
}

func buildGet(build BuildRequest) (*http.Request, error) {
	if build.path == "" {
		return nil, errors.New("can not generate request without path")
	}

	url := buildUrl(build.path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if build.headers != nil {
		for k, v := range build.headers {
			req.Header.Add(k, v)
		}
	}

	return req, nil
}

func buildPost(build BuildRequest) (*http.Request, error) {
	if build.path == "" {
		return nil, errors.New("can not generate request without path")
	}

	url := buildUrl(build.path)
	req, err := http.NewRequest("POST", url, build.requestBody)
	if err != nil {
		return nil, err
	}

	if build.headers != nil {
		for k, v := range build.headers {
			req.Header.Add(k, v)
		}
	}

	return req, nil
}

func buildPut(build BuildRequest) (*http.Request, error) {
	if build.path == "" {
		return nil, errors.New("can not generate request without path")
	}

	url := buildUrl(build.path)
	req, err := http.NewRequest("PUT", url, build.requestBody)
	if err != nil {
		return nil, err
	}

	if build.headers != nil {
		for k, v := range build.headers {
			req.Header.Add(k, v)
		}
	}

	return req, nil
}

func buildDelete(build BuildRequest) (*http.Request, error) {
	if build.path == "" {
		return nil, errors.New("can not generate request without path")
	}

	url := buildUrl(build.path)
	req, err := http.NewRequest("DELETE", url, build.requestBody)
	if err != nil {
		return nil, err
	}

	if build.headers != nil {
		for k, v := range build.headers {
			req.Header.Add(k, v)
		}
	}

	return req, nil
}

func getAllHeaders() map[string]string {
	var headers = make(map[string]string)
	headers[AuthHeaderDevopsBuildType] = gSdkEvn.BuildType
	headers[AuthHeaderProjectId] = gSdkEvn.ProjectId
	headers[AuthHeaderDevopsProjectId] = gSdkEvn.ProjectId
	headers[AuthHeaderDevopsBuildId] = gSdkEvn.BuildId
	headers[AuthHeaderDevopsAgentSecretKey] = gSdkEvn.SecretKey
	headers[AuthHeaderDevopsAgentId] = gSdkEvn.AgentId
	headers[AuthHeaderDevopsVmSeqId] = gSdkEvn.VmSeqId
	headers[AuthHeaderBuildId] = gSdkEvn.BuildId
	headers[AuthHeaderDevopsCiTaskId] = gSdkEvn.TaskId
	return headers
}

func buildUrl(path string) string {
	var gateway = strings.TrimSuffix(gSdkEvn.Gateway, "/")
	if strings.HasPrefix(gateway, "http") {
		return gateway + "/" + strings.TrimPrefix(strings.TrimSpace(path), "/")
	} else {
		return "http://" + gateway + "/" + strings.TrimPrefix(strings.TrimSpace(path), "/")
	}
}
