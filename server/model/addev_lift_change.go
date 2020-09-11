package model

import "gorm.io/gorm"

type LiftChange struct {
	gorm.Model
	LiftId  int    `json:"liftId" form:"liftId"`
	Lift    Lift   `json:"lift" form:"lift" gorm:"foreignKey:LiftId"`
	Content string `json:"content" form:"content"`
}
