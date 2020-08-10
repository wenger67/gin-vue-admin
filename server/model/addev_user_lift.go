package model

import "github.com/jinzhu/gorm"

type UserLift struct {
	gorm.Model
	UserId     int
	UserAdmin  SysUser
	LiftId     int
	Lift       Lift
	CategoryId int
	Category   Category `gorm:"comment:'用户电梯关系类型'"`
}