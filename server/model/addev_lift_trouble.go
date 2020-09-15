package model

import (
	"gorm.io/gorm"
	"time"
)

type LiftTrouble struct {
	gorm.Model
	// belong to
	LiftId           uint `json:"liftId" form:"liftId"`
	Lift             Lift `json:"lift" form:"lift" gorm:"ForeignKey:LiftId;"`
	// belong to
	FromCategoryId   uint `json:"fromCategoryId" form:"fromCategoryId"`  // subject:104
	FromCategory     Category `json:"fromCategory" form:"fromCategory" gorm:"ForeignKey:FromCategoryId"`
	StartTime        time.Time `json:"startTime" form:"startTime" gorm:"default:null"`
	// belong to
	StartUserId      uint `json:"startUserId" from:"startUserId" gorm:"default:null"`
	StartUser        SysUser `json:"startUser" form:"startUser" gorm:"ForeignKey:StartUserId"`
	ResponseTime     time.Time `json:"responseTime" form:"responseTime" gorm:"default:null"`
	// belong to
	ResponseUserId   uint  `json:"responseUserId" from:"responseUserId" gorm:"default:null"`
	ResponseUser     SysUser `json:"responseUser" form:"responseUser" gorm:"ForeignKey:ResponseUserId"`
	SceneTime        time.Time `json:"sceneTime" form:"sceneTime" gorm:"default:null"`
	// belong to
	SceneUserId      uint `json:"sceneUserId" from:"sceneUserId" gorm:"default:null"`
	SceneUser        SysUser `json:"sceneUser" form:"sceneUser" gorm:"ForeignKey:SceneUserId"`
	FixTime          time.Time `json:"fixTime" form:"fixTime" gorm:"default:null"`
	// belong to
	FixUserId        uint `json:"fixUserId" from:"fixUserId" gorm:"default:null"`
	FixUser          SysUser `json:"fixUser" form:"fixUser" gorm:"ForeignKey:FixUserId"`
	// belong to
	FixCategoryId    uint `json:"fixCategoryId" form:"fixCategoryId" gorm:"default:null"` // subject:105
	FixCategory      Category `json:"fixCategory" form:"fixCategory" gorm:"ForeignKey:FixCategoryId"`
	// belong to
	ReasonCategoryId uint `json:"reasonCategoryId" form:"reasonCategoryId" gorm:"default:null"` // subject:106
	ReasonCategory   Category `json:"reasonCategory" form:"reasonCategory" gorm:"ForeignKey:ReasonCategoryId"`

	Medias []ExaFileUploadAndDownload `json:"medias" gorm:"foreignKey:TroubleId"`
	Content          string `json:"content" form:"content" gorm:"default:null"`
	// belong to
	RecorderId       uint  `json:"recorderId" from:"recorderId" gorm:"default:null"`
	Recorder         SysUser `json:"recorder" form:"recorder" gorm:"ForeignKey:RecorderId"`
	FeedbackContent  string `json:"feedbackContent" form:"feedbackContent"`
	FeedbackRate     int `json:"feedbackRate" form:"feedbackRate"`
	// 1: create 2: response 3:scene 4:fix 5:feedback 6:recorder
	Progress         int `json:"progress" form:"progress"`
}
