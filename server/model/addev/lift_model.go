package addev

import "github.com/jinzhu/gorm"

type LiftModel struct {
	gorm.Model
	FactoryId int
	Company Company `gorm:"ForeignKey:FactoryId;AssociationForeignKey:FactoryId;comment:'电梯型号'"`
	Brand string `gorm:"comment:'电梯品牌'"`
	Modal string `gorm:"comment:'电梯型号'"`
	Load int `gorm:"comment:'电梯载重'"`
}
