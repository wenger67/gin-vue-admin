package model

import "github.com/jinzhu/gorm"

type UserLift struct {
	gorm.Model
	UserId     int	`json:"userId" form:"userId"`
	UserAdmin  SysUser `json:"user" form:"user"`
	LiftId     int `json:"liftId" form:"liftId"`
	Lift       Lift	`json:"lift" form:"lift"`
	CategoryId int	`json:"categoryId" form:"categoryId"`
	Category   Category `json:"category" form:"category" gorm:"comment:'用户电梯关系类型'"`
}