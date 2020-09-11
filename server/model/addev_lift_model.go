package model

import "gorm.io/gorm"

type LiftModel struct {
	gorm.Model
	FactoryId int `json:"factoryId" form:"factoryId"`
	Company   Company `json:"company" form:"company" gorm:"ForeignKey:FactoryId;comment:'电梯型号'"`
	Brand     string  `json:"brand" form:"brand" gorm:"comment:'电梯品牌'"`
	Modal     string  `json:"modal" form:"modal" gorm:"comment:'电梯型号'"`
	Load      int     `json:"load" form:"load" gorm:"comment:'电梯载重'"`
}
