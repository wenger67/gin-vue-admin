package model

import (
	"gorm.io/gorm"
	"time"
)

type LiftRecord struct {
	// progress 1: created
	gorm.Model
	LiftId uint `json:"liftId" form:"liftId"`	// belong
	Lift   Lift `json:"lift" form:"lift" gorm:"ForeignKey:LiftId"`
	CategoryId uint      `json:"categoryId" form:"categoryId"`	// belong to
	Category   Category  `json:"category" form:"category" gorm:"foreignKey:CategoryId"`

	// progress 2: start
	StartTime  time.Time `json:"startTime" form:"startTime" gorm:"default:null"`
	WorkerId   uint       `json:"workerId" form:"workerId" gorm:"default:null"`	// belong to
	Worker     SysUser   `json:"worker" form:"worker" gorm:"ForeignKey:WorkerId"`

	// progress 3: end
	Medias     []ExaFileUploadAndDownload    `json:"medias" form:"medias" gorm:"foreignKey:RecordId"`
	Content    string    `json:"content" form:"content"`
	EndTime    time.Time `json:"endTime" form:"endTime" gorm:"default:null"`

	// progress 4: review
	RecorderId uint       `json:"recorderId" form:"recorderId" gorm:"default:null"`	// belong to
	Recorder   SysUser   `json:"recorder" form:"recorder" gorm:"ForeignKey:RecorderId"`
	// 1: created, 2: started 3:ended 4:recorderid
	Progress   int       `json:"progress" form:"progress" gorm:"default:1"`
}
