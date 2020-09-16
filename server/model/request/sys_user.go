package request

import (
	"gin-vue-admin/model"
	uuid "github.com/satori/go.uuid"
)

// User register structure
type RegisterStruct struct {
	PhoneNumber  string `json:"phoneNumber"`
	Password    string `json:"passWord"`
	NickName    string `json:"nickName" gorm:"default:'QMPlusUser'"`
	Avatar   string `json:"avatar" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	AuthorityId string `json:"authorityId" gorm:"default:888"`
}

// User login structure
type RegisterAndLoginStruct struct {
	PhoneNumber  string `json:"phoneNumber"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

// Modify password structure
type ChangePasswordStruct struct {
	PhoneNumber  string `json:"phoneNumber"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}

type CreateStruct struct {
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Password    string `json:"password" form:"password"`
	RealName    string `json:"realName" form:"realName"`
	NickName    string `json:"nickName" form:"nickName"`
	Avatar      string `json:"avatar" form:"avatar"`
	CompanyId   uint `json:"companyId" form:"companyId"`
	Address     string `json:"address" form:"address"`
	AuthorityId string       `json:"authorityId" form:"authorityId" gorm:"default:888;comment:'用户角色ID'"`
}

type SearchUserParams struct {
	model.SysUser
	PageInfo
	CompanyId uint `json:"companyId" form:"companyId"`
}
