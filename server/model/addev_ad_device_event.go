package model

import "github.com/jinzhu/gorm"

type AdDeviceEvent struct {
	gorm.Model
	DeviceId int `json:"device_id" form:"device_id"`
	Device AdDevice `json:"device" form:"device"`
	TypeId int `json:"type_id" form:"type_id"`
	Type Category `json:"type" form:"type" gorm:"ForiegnKey:TypeId;AssociationForiegnKey:ID;comment:'事件类型'"`
	Content string `json:"content" form:"content" gorm:"comment:'事件内容'"`
}
