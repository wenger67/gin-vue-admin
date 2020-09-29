package model

import (
	"gorm.io/gorm"
	"time"
)

type LiftTrouble struct {
	// progress 1: created
	gorm.Model
	LiftId           uint `json:"liftId" form:"liftId"`	// belong to
	Lift             Lift `json:"lift" form:"lift" gorm:"foreignKey:LiftId;"`
	FromCategoryId   uint `json:"fromCategoryId" form:"fromCategoryId"`  // subject:104	// belong to
	FromCategory     Category `json:"fromCategory" form:"fromCategory" gorm:"foreignKey:FromCategoryId"`
	StartTime        time.Time `json:"startTime" form:"startTime" gorm:"default:null"`
	StartUserId      uint `json:"startUserId" from:"startUserId" gorm:"default:null"`	// belong to
	StartUser        SysUser `json:"startUser" form:"startUser" gorm:"foreignKey:StartUserId"`

	// progress 2: response
	ResponseTime     time.Time `json:"responseTime" form:"responseTime" gorm:"default:null"`
	ResponseUserId   uint  `json:"responseUserId" from:"responseUserId" gorm:"default:null"`	// belong to
	ResponseUser     SysUser `json:"responseUser" form:"responseUser" gorm:"foreignKey:ResponseUserId"`

	// progress 3: scene
	SceneTime        time.Time `json:"sceneTime" form:"sceneTime" gorm:"default:null"`
	SceneUserId      uint `json:"sceneUserId" from:"sceneUserId" gorm:"default:null"`	// belong to
	SceneUser        SysUser `json:"sceneUser" form:"sceneUser" gorm:"foreignKey:SceneUserId"`

	// progress 4: fix
	FixTime          time.Time `json:"fixTime" form:"fixTime" gorm:"default:null"`
	FixUserId        uint `json:"fixUserId" from:"fixUserId" gorm:"default:null"`	// belong to
	FixUser          SysUser `json:"fixUser" form:"fixUser" gorm:"foreignKey:FixUserId"`
	FixCategoryId    uint `json:"fixCategoryId" form:"fixCategoryId" gorm:"default:null"` // subject:105	// belong to
	FixCategory      Category `json:"fixCategory" form:"fixCategory" gorm:"foreignKey:FixCategoryId"`
	ReasonCategoryId uint `json:"reasonCategoryId" form:"reasonCategoryId" gorm:"default:null"` // subject:106	// belong to
	ReasonCategory   Category `json:"reasonCategory" form:"reasonCategory" gorm:"foreignKey:ReasonCategoryId"`

	Medias []ExaFileUploadAndDownload `json:"medias" gorm:"foreignKey:TroubleId"`
	Content          string `json:"content" form:"content" gorm:"default:null"`	// belong to

	// progress 5: feedback
	FeedbackContent  string `json:"feedbackContent" form:"feedbackContent"`
	FeedbackRate     float32 `json:"feedbackRate" form:"feedbackRate"`

	// progress 6: review
	RecorderId       uint  `json:"recorderId" from:"recorderId" gorm:"default:null"`
	Recorder         SysUser `json:"recorder" form:"recorder" gorm:"foreignKey:RecorderId"`
	
	// 1: create 2: response 3:scene 4:fix 5:feedback 6:recorder
	Progress         int `json:"progress" form:"progress"`
}
