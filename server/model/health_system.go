package model

import "gorm.io/gorm"

type HealthSystem struct {
	gorm.Model
	LiftId            int  `json:"liftId" form:"liftId"`
	Lift              Lift `json:"lift" form:"lift" gorm:"ForeignKey:LiftId;"`
	TimeDimension     int  `json:"timeDimension" gorm:"timeDimension"`
	MaintainDimension int  `json:"maintainDimension" gorm:"maintainDimension"`
	HumanDimension    int  `json:"humanDimension" gorm:"humanDimension"`
	InnerDimension    int  `json:"innerDimension" gorm:"innerDimension"`
	SensorDimension   int  `json:"sensorDimension" gorm:"sensorDimension"`
}
