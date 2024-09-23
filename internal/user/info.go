package user

import (
	"fmt"

	"github.com/soulteary/modelscope-open-world/api/modelscope"
)

type Info struct {
	Email        string `json:"Email"`
	GitlabUserID int    `json:"GitlabUserID"`
	HavanaID     string `json:"HavanaID"`
	Name         string `json:"Name"`
	IsRealName   bool   `json:"IsRealName"`
	AccessToken  string `json:"AccessToken"`
}

// Basic information provided to third-party software for user login
func GetBasicUserInfo(cookies string) (result Info, err error) {
	info, err := modelscope.GetUserInfo(cookies)
	if err != nil {
		return result, err
	}

	if !info.Success {
		return result, fmt.Errorf(info.Message)
	}

	result.Email = info.Data.Email
	result.GitlabUserID = info.Data.GitlabUserID
	result.HavanaID = info.Data.HavanaID
	result.Name = info.Data.Name
	result.IsRealName = info.Data.IsRealName
	result.AccessToken = info.Data.AccessToken

	return result, nil
}
