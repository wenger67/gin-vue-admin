package model

import "github.com/jinzhu/gorm"

type AdDeviceData struct {
	gorm.Model
	DeviceId int `json:"device_id" form:"device_id"`
	Device AdDevice `json:"device" form:"device" gorm:"ForiegnKey:DeviceId;AssociationForiegnKey:ID;"`
	TroubleId int `json:"trouble_id" form:"trouble_id"`
	Trouble Category `json:"trouble" form:"trouble" gorm:"ForiegnKey:TroubleId;AssociationForiegnKey:ID;comment:'电梯状态/故障类型'"`
	Accx float32 `json:"accx" form:"accx"`
	Accy float32 `json:"accy" form:"accy"`
	Accz float32 `json:"accz" form:"accz"`
	Degx float32 `json:"degx" form:"degx"`
	Degy float32 `json:"degy" form:"degy"`
	Degz float32 `json:"degz" form:"degz"`
	Speedz float32 `json:"speedz" form:"speedz"`
	Floor float32 `json:"floor" form:"floor"`
	DoorStateId int `json:"door_state_id" form:"door_state_id" gorm:"comment:'开门状态'"`
	DoorState Category `json:"door_state" form:"door_state" gorm:"ForiegnKey:DoorStateId;AssociationForiegnKey:ID"`
	PeopleInside bool `json:"people_inside" form:"people_inside"`
}

