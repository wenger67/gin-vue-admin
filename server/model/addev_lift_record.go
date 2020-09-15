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
	Medias     []ExaFileUploadAndDownload    `json:"medias" form:"medias" gorm:"foreignKey:RecordId"`
	Content    string    `json:"content" form:"content"`
	StartTime  time.Time `json:"startTime" form:"startTime" gorm:"default:null"`
	EndTime    time.Time `json:"endTime" form:"endTime" gorm:"default:null"`
	// belong to
	WorkerId   uint       `json:"workerId" form:"workerId" gorm:"default:null"`
	Worker     SysUser   `json:"worker" form:"worker" gorm:"ForeignKey:WorkerId"`
	// belong to
	RecorderId uint       `json:"recorderId" form:"recorderId" gorm:"default:null"`
	Recorder   SysUser   `json:"recorder" form:"recorder" gorm:"ForeignKey:RecorderId"`
	// 1: created, 2: started 3:ended 4:recorderid
	Progress   int       `json:"progress" form:"progress" gorm:"default:1"`
}
