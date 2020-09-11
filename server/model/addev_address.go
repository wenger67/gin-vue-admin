package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	RegionId int `json:"regionId" form:"regionId"`
	Region Region `json:"region" form:"region" gorm:"foreignKey:RegionId;"`
	AddressName string `json:"addressName" form:"addressName"`
	Location string `json:"location" form:"location"`
	UserAmount int `json:"userAmount" form:"userAmount"`
	Tags []Category `json:"tags" form:"tags" gorm:"many2many:address_tags"`
}
