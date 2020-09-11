package model

import "gorm.io/gorm"

type UserLift struct {
	gorm.Model
	// belong to
	UserId     uint	`json:"userId" form:"userId"`
	User  SysUser `json:"user" form:"user" gorm:"foreignKey:UserId"`
	// belong to
	LiftId     uint `json:"liftId" form:"liftId"`
	Lift       Lift	`json:"lift" form:"lift" gorm:"foreignKey:LiftId"`
	// belong to
	CategoryId uint	`json:"categoryId" form:"categoryId"`
	Category   Category `json:"category" form:"category" gorm:"foreignKey:CategoryId;comment:'用户电梯关系类型'"`
}