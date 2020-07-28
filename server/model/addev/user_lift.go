package addev

import "github.com/jinzhu/gorm"

type UserLift struct {
	gorm.Model
	UserId int
	UserAdmin UserAdmin
	LiftId int
	Lift Lift
	CategoryId int
	Category Category `gorm:"comment:'用户电梯关系类型'"`
}