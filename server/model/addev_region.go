package model

import "github.com/jinzhu/gorm"

type Region struct {
	gorm.Model
	Province string
	City string
	District string
}
