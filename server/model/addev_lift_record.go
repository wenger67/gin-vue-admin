package model

import (
	"gorm.io/gorm"
	"time"
)

type LiftRecord struct {
	gorm.Model
	// belong
	LiftId uint `json:"liftId" form:"liftId"`
	Lift   Lift `json:"lift" form:"lift" gorm:"ForeignKey:LiftId"`
	// belong to
	CategoryId uint      `json:"categoryId" form:"categoryId"`
	Category   Category  `json:"category" form:"category" gorm:"foreignKey:CategoryId"`
	Images     string    `json:"images" form:"images"`
	Content    string    `json:"content" form:"content"`
	StartTime  time.Time `json:"startTime" form:"startTime"`
	EndTime    time.Time `json:"endTime" form:"endTime"`
	// belong to
	WorkerId   uint       `json:"workerId" form:"workerId"`
	Worker     SysUser   `json:"worker" form:"worker" gorm:"ForeignKey:WorkerId"`
	// belong to
	RecorderId uint       `json:"recorderId" form:"recorderId"`
	Recorder   SysUser   `json:"recorder" form:"recorder" gorm:"ForeignKey:RecorderId"`
	Progress   int       `json:"progress" form:"progress"`
}
