package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type LiftRecord struct {
	gorm.Model
	LiftId     int `json:"liftId" form:"liftId"`
	Lift       Lift `json:"lift" form:"lift" gorm:"ForeignKey:LiftId;AssociationForeignKey:ID"`
	CategoryId int `json:"categoryId" form:"categoryId"`
	Category   Category `json:"category" form:"category"`
	Images     string `json:"images" form:"images"`
	Content    string `json:"content" form:"content"`
	StartTime  time.Time `json:"startTime" form:"startTime"`
	EndTime    time.Time `json:"endTime" form:"endTime"`
	WorkerId   int `json:"workerId" form:"workerId"`
	Worker     SysUser `json:"worker" form:"worker" gorm:"ForeignKey:WorkerId;AssociationForeignKey:ID"`
	RecorderId int `json:"recorderId" form:"recorderId"`
	Recorder   SysUser `json:"recorder" form:"recorder" gorm:"ForeignKey:RecorderId;AssociationForeignKey:ID"`
	Progress int `json:"progress" form:"progress"`
}
