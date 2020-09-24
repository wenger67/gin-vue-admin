package response

import (
	"panta/model"
)

type SysUserResponse struct {
	User model.SysUser `json:"user"`
}

type LoginResponse struct {
	User      model.SysUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}

type UserCountByCompany struct {
	CompanyId   int    `json:"companyId"`
	FullName string `json:"companyName"`
	Total       int    `json:"total"`
}
