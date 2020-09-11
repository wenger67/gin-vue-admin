package model

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	UUID        uuid.UUID `json:"uuid" form:"uuid"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Password    string `json:"-"`
	RealName    string `json:"realName" form:"realName"`
	NickName    string `json:"nickName" form:"nickName"`
	Avatar      string `json:"avatar" form:"avatar"`
	// belong to
	CompanyId   int `json:"companyId" form:"companyId"`
	Company     Company `json:"company" form:"company" gorm:"ForeignKey:CompanyId"`
	Address     string `json:"address" form:"address"`
	// belong to
	Authority   SysAuthority `json:"authority" gorm:"ForeignKey:AuthorityId;comment:'用户角色'"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:'用户角色ID'"`
}
