package model

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	RegionId int `json:"regionId" form:"regionId"`
	Region Region `json:"region" form:"region" gorm:"ForeignKey:RegionId;AssociationForeignKey:ID"`
	AddressName string `json:"addressName" form:"addressName"`
	UserAmount int `json:"userAmount" form:"userAmount"`
	Tags []AddressTag `json:"tags" form:"tags" gorm:"many2many:address_tags"`
}
