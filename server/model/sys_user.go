package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	gorm.Model
	UUID        uuid.UUID `json:"uuid" form:"uuid"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Password    string `json:"-"`
	RealName    string `json:"realName" form:"realName"`
	NickName    string `json:"nickName" form:"nickName"`
	Avatar      string `json:"avatar" form:"avatar"`
	CompanyId   int `json:"companyId" form:"companyId"`
	Company     Company `json:"company" form:"company" gorm:"ForeignKey:CompanyId;AssociationForeignKey:ID"`
	Address     string `json:"address" form:"address"`
	Authority   SysAuthority `json:"authority" gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId;comment:'用户角色'"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:'用户角色ID'"`
}
