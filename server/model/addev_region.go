package model

import "github.com/jinzhu/gorm"

type Region struct {
	gorm.Model
	Province string `json:"province" form:"province"`
	City string	`json:"city" form:"city"`
	District string	`json:"district" form:"district"`
}
