package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// reCAPTCHA验证服务的响应结构
type RecaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTS string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}

/*
// 谷歌人机验证

	responseKey: 客户端秘钥
	secretKey: 服务端秘钥
	recaptchaLink: 请求地址
*/
func VerifyReCaptchaV2(responseKey string, secretKey string, recaptchaLink string) bool {

	// 设置请求参数
	params := url.Values{}
	params.Add("response", responseKey)
	params.Add("secret", secretKey)

	// 构建完整URL
	baseURL := recaptchaLink
	requestURL := baseURL + "?" + params.Encode()

	// 发起GET请求
	resp, err := http.Get(requestURL)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
		return false
	}
	defer resp.Body.Close() // 确保请求结束后关闭连接
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read response body failed: %v", err)
		return false
	}

	var recaptchaResp RecaptchaResponse
	err = json.Unmarshal(body, &recaptchaResp)
	if err != nil {
		log.Fatalf("Parse JSON response failed: %v", err)
		return false
	}

	fmt.Println(recaptchaResp)

	// 根据success字段判断并返回结果
	if recaptchaResp.Success {
		log.Println("reCAPTCHA succeeded.")
		return true
	} else {
		log.Println("reCAPTCHA failed.")
		return false
	}

}
