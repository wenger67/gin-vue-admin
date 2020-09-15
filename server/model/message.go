package model

import "gorm.io/gorm"

/**

 */
type Message struct {
	gorm.Model
	FromId     uint      `json:"fromId" form:"fromId" gorm:"default:null"`
	FromUser SysUser	`json:"fromUser" form:"fromUser" gorm:"foreignKey:FromId"`
	TargetId   uint      `json:"targetId" form:"targetId"`
	TargetUser SysUser `json:"targetType" form:"targetType" gorm:"ForeignKey:TargetId"`
	Content    string   `json:"content" form:"content"`
	TypeId     uint      `json:"typeId" form:"typeId"`
	Type       Category `json:"type" form:"type" gorm:"ForeignKey:TypeId;"`
	Send       bool     `json:"send" form:"send" gorm:"default:0"`
	Read       bool     `json:"read" form:"read" gorm:"default:0"`
}
