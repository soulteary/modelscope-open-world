package modelscope

import (
	"encoding/json"
	"io"
	"net/http"
)

const MODELSCOPE_LOGIN_USERINFO_URL = "https://modelscope.cn/api/v1/users/login/info"

type LoginUserInfoResponse struct {
	Code      int               `json:"Code"`
	Data      LoginUserInfoData `json:"Data"`
	Message   string            `json:"Message"`
	RequestID string            `json:"RequestId"`
	Success   bool              `json:"Success"`
}

type LoginUserInfoData struct {
	AccessToken       string `json:"AccessToken"`
	Avatar            string `json:"Avatar"`
	Description       string `json:"Description"`
	Email             string `json:"Email"`
	FromSite          string `json:"FromSite"`
	FullName          string `json:"FullName"`
	GitlabAccessToken string `json:"GitlabAccessToken"`
	GitlabUserID      int    `json:"GitlabUserId"`
	HavanaID          string `json:"HavanaId"`
	IsCertification   string `json:"IsCertification"`
	Name              string `json:"Name"`
	OrgRoleMap        any    `json:"OrgRoleMap,omitempty"`
	Roles             any    `json:"Roles"`
	SecurityModel     string `json:"SecurityModel"`
	WorkNo            string `json:"WorkNo"`
	IsRealName        bool   `json:"is_real_name"`
}

// Obtain user information from ModelScope
func GetUserInfo(cookies string) (result LoginUserInfoResponse, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", MODELSCOPE_LOGIN_USERINFO_URL, nil)
	if err != nil {
		return result, err
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", cookies)
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://modelscope.cn/my/overview")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")
	req.Header.Set("x-modelscope-accept-language", "zh_CN")
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
