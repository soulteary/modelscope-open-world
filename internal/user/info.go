package user

import (
	"fmt"

	"github.com/soulteary/modelscope-open-world/api/modelscope"
)

type UserInfo struct {
	Email        string `json:"Email"`
	GitlabUserID int    `json:"GitlabUserID"`
	HavanaID     string `json:"HavanaID"`
	Name         string `json:"Name"`
	IsRealName   bool   `json:"IsRealName"`
}

// Basic information provided to third-party software for user login
func GetBasicUserInfo(cookies string) (result UserInfo, err error) {
	info, err := modelscope.GetUserInfo(cookies)
	if err != nil {
		return result, err
	}

	if !info.Success {
		return result, fmt.Errorf(info.Message)
	}

	return UserInfo{
		Email:        info.Data.Email,
		GitlabUserID: info.Data.GitlabUserID,
		HavanaID:     info.Data.HavanaID,
		Name:         info.Data.Name,
		IsRealName:   info.Data.IsRealName,
	}, nil
}
