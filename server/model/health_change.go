package model

import "gorm.io/gorm"

type HealthChange struct {
	gorm.Model
	LiftId      uint      `json:"liftId" form:"liftId"`
	Lift        Lift     `json:"lift" form:"lift" gorm:"ForeignKey:LiftId;AssociationForeignKey:ID"`
	DimensionId uint      `json:"dimensionId" form:"dimensionId"`
	Dimension   Category `json:"dimension" form:"dimension" gorm:"ForeignKey:DimensionId;AssociationForeignKey:ID"`
	Content     string   `json:"content" form:"content"`
	Score       int      `json:"score" form:"score"`
}
