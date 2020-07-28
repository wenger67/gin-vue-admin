package addev

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type UserAdmin struct {
	gorm.Model
	UUID uuid.UUID
	PhoneNumber string
	Password string
	RealName string
	NickName string
	Avatar string
	RoleId int
	RoleName string
	CompanyId int
	Company Company
	Address string
}
