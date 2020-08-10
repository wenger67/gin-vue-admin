package model

import "github.com/jinzhu/gorm"

type AdDeviceData struct {
	gorm.Model
	DeviceId int `json:"deviceId" form:"deviceId"`
	Device AdDevice `json:"device" form:"device" gorm:"ForiegnKey:DeviceId;AssociationForiegnKey:ID;"`
	TroubleId int `json:"troubleId" form:"troubleId"`
	Trouble Category `json:"trouble" form:"trouble" gorm:"ForiegnKey:TroubleId;AssociationForiegnKey:ID;comment:'电梯状态/故障类型'"`
	Accx float32 `json:"accx" form:"accx"`
	Accy float32 `json:"accy" form:"accy"`
	Accz float32 `json:"accz" form:"accz"`
	Degx float32 `json:"degx" form:"degx"`
	Degy float32 `json:"degy" form:"degy"`
	Degz float32 `json:"degz" form:"degz"`
	Speedz float32 `json:"speedz" form:"speedz"`
	Floor float32 `json:"floor" form:"floor"`
	DoorStateId int `json:"doorStateId" form:"doorStateId" gorm:"comment:'开门状态'"`
	DoorState Category `json:"doorState" form:"doorState" gorm:"ForiegnKey:DoorStateId;AssociationForiegnKey:ID"`
	PeopleInside bool `json:"peopleInside" form:"peopleInside"`
}

