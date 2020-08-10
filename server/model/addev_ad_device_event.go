package model

import "github.com/jinzhu/gorm"

type AdDeviceEvent struct {
	gorm.Model
	DeviceId int `json:"deviceId" form:"deviceId"`
	Device AdDevice `json:"device" form:"device"`
	TypeId int `json:"typeId" form:"typeId"`
	Type Category `json:"type" form:"type" gorm:"ForiegnKey:TypeId;AssociationForiegnKey:ID;comment:'事件类型'"`
	Content string `json:"content" form:"content" gorm:"comment:'事件内容'"`
}
