package model

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	AddressName string
}
