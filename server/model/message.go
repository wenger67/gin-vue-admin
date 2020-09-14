package model

import "gorm.io/gorm"

/**

 */
type Message struct {
	gorm.Model
	FromId     int      `json:"fromId" form:"fromId" gorm:"ForeignKey:FromId;AssociationForeignKey:ID"`
	FromType   Category `json:"fromType" form:"fromType"`
	TargetId   int      `json:"targetId" form:"targetId" gorm:"ForeignKey:TargetId;AssociationForeignKey:ID"`
	TargetType Category `json:"targetType" form:"targetType"`
	Content    string   `json:"content" form:"content"`
	TypeId     int      `json:"typeId" form:"typeId" gorm:"ForeignKey:TypeId;AssociationForeignKey:ID"`
	Type       Category `json:"type" form:"type"`
	Send       bool     `json:"send" form:"send"`
	Read       bool     `json:"read" form:"read"`
}
