package model

import "gorm.io/gorm"

type AdDeviceEvent struct {
	gorm.Model
	// belong to
	DeviceID uint      `json:"deviceId" form:"deviceId"`
	Device   Device   `json:"device" form:"device"`
	// belong to
	TypeId   uint      `json:"typeId" form:"typeId"`
	Type     Category `json:"type" form:"type" gorm:"foreignKey:TypeId;comment:'事件类型'"`
	/**
	{cl:[{brief:"", images:[], createdAt:},{brief:"", images:[], createdAt:},
	{brief:"", images:[], createdAt:}], v1:[]]}
	 */
	Content string `json:"content" form:"content" gorm:"comment:'事件内容'"`
}
