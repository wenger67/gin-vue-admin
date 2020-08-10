package model

import "github.com/jinzhu/gorm"

type LiftChange struct {
	gorm.Model
	LiftId  int `json:"liftId" form:"liftId"`
	Lift    Lift `json:"lift" form:"lift" gorm:"ForiegnKey:LiftId;AssociationForiegnKey:ID"`
	Content string `json:"content" form:"content"`
}
