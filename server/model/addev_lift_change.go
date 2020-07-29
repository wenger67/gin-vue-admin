package model

import "github.com/jinzhu/gorm"

type LiftChange struct {
	gorm.Model
	LiftId  int
	Lift    Lift
	Content string
}
