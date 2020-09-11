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
	StartTime        time.Time `json:"startTime" form:"startTime"`
	// belong to
	StartUserId      uint `json:"startUserId" from:"startUserId"`
	StartUser        SysUser `json:"startUser" form:"startUser" gorm:"ForeignKey:StartUserId"`
	ResponseTime     time.Time `json:"responseTime" form:"responseTime"`
	// belong to
	ResponseUserId   uint  `json:"responseUserId" from:"responseUserId"`
	ResponseUser     SysUser `json:"responseUser" form:"responseUser" gorm:"ForeignKey:ResponseUserId"`
	SceneTime        time.Time `json:"sceneTime" form:"sceneTime"`
	// belong to
	SceneUserId      uint `json:"sceneUserId" from:"sceneUserId"`
	SceneUser        SysUser `json:"sceneUser" form:"sceneUser" gorm:"ForeignKey:SceneUserId"`
	FixTime          time.Time `json:"fixTime" form:"fixTime"`
	// belong to
	FixUserId        uint `json:"fixUserId" from:"fixUserId"`
	FixUser          SysUser `json:"fixUser" form:"fixUser" gorm:"ForeignKey:FixUserId"`
	// belong to
	FixCategoryId    uint `json:"fixCategoryId" form:"fixCategoryId"` // subject:105
	FixCategory      Category `json:"fixCategory" form:"fixCategory" gorm:"ForeignKey:FixCategoryId"`
	// belong to
	ReasonCategoryId int `json:"reasonCategoryId" form:"reasonCategoryId"` // subject:106
	ReasonCategory   Category `json:"reasonCategory" form:"reasonCategory" gorm:"ForeignKey:ReasonCategoryId"`
	Content          string `json:"content" form:"content"`
	Progress         int `json:"progress" form:"progress"`
	// belong to
	RecorderId       uint  `json:"recorderId" from:"recorderId"`
	Recorder         SysUser `json:"recorder" form:"recorder" gorm:"ForeignKey:RecorderId"`
	FeedbackContent  string `json:"feedbackContent" form:"feedbackContent"`
	FeedbackRate     int `json:"feedbackRate" form:"feedbackRate"`
}
