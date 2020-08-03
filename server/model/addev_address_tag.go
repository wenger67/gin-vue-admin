package model

import "github.com/jinzhu/gorm"

type AddressTag struct {
	gorm.Model
	AddressId int `json:"addressId" form:"addressId"`
	Address Address `json:"address" form:"address" gorm:"ForeignKey:AddressId;AssociationForeignKey:ID"`
	TagId int `json:"tagId" form:"tagId"`
	Tag Category `json:"tag" form:"tag" gorm:"ForeignKey:TagId;AssociationForeignKey:ID"`
}
