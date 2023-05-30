package api

import (
	"encoding/json"

	"github.com/ci-plugins/golang-plugin-sdk/log"
)

// Certificate 请求凭证结果
type Certificate struct {
	status int
	Data   map[string]string `json:"data"`
}

// GetCertificate 获取指定ID的凭证
func GetCertificate(certificateId string) map[string]string {
	log.Info("Begin to get certificate")
	url := buildUrl("/ticket/api/build/credentials/" + certificateId + "/detail")
	var build = BuildRequest{path: url, requestBody: nil, headers: getAllHeaders()}
	req, err := buildGet(build)
	if err != nil {
		log.Error("build request failed: " + err.Error())
		return nil
	}

	respByte, err := request(*req, "failed to get certificate")
	if err != nil {
		log.Error("get certificate failed: " + err.Error())
		return nil
	}
	var certificate = new(Certificate)
	err = json.Unmarshal(respByte, &certificate)

	if err != nil {
		log.Error("resolve response json failed: " + err.Error())
	}

	return certificate.Data
}
