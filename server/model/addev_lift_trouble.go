package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type LiftTrouble struct {
	gorm.Model
	LiftId           int `json:"liftId" form:"liftId"`
	Lift             Lift `json:"lift" form:"lift" gorm:"ForiegnKey:LiftId;AssociationForiegnKey:ID"`
	FromCategoryId   int `json:"fromCategoryId" form:"fromCategoryId"`  // subject:104
	FromCategory     Category `json:"fromCategory" form:"fromCategory" gorm:"ForiegnKey:FromCategoryId;AssociationForiegnKey:ID"`
	StartTime        time.Time `json:"startTime" form:"startTime"`
	StartUserId      int `json:"startUserId" from:"startUserId"`
	StartUser        SysUser `json:"startUser" form:"startUser" gorm:"ForiegnKey:StartUserId;AssociationForiegnKey:ID"`
	ResponseTime     time.Time `json:"responseTime" form:"responseTime"`
	ResponseUserId   int  `json:"responseUserId" from:"responseUserId"`
	ResponseUser     SysUser `json:"responseUser" form:"responseUser" gorm:"ForiegnKey:ResponseUserId;AssociationForiegnKey:ID"`
	SceneTime        time.Time `json:"sceneTime" form:"sceneTime"`
	SceneUserId      int `json:"sceneUserId" from:"sceneUserId"`
	SceneUser        SysUser `json:"sceneUser" form:"sceneUser" gorm:"ForiegnKey:SceneUserId;AssociationForiegnKey:ID"`
	FixTime          time.Time `json:"fixTime" form:"fixTime"`
	FixUserId        int `json:"fixUserId" from:"fixUserId"`
	FixUser          SysUser `json:"fixUser" form:"fixUser" gorm:"ForiegnKey:FixUserId;AssociationForiegnKey:ID"`
	FixCategoryId    int `json:"fixCategoryId" form:"fixCategoryId"` // subject:105
	FixCategory      Category `json:"fixCategory" form:"fixCategory" gorm:"ForiegnKey:FixCategoryId;AssociationForiegnKey:ID"`
	ReasonCategoryId int `json:"reasonCategoryId" form:"reasonCategoryId"` // subject:106
	ReasonCategory   Category `json:"reasonCategory" form:"reasonCategory" gorm:"ForiegnKey:ReasonCategoryId;AssociationForiegnKey:ID"`
	Content          string `json:"content" form:"content"`
	Progress         int `json:"progress" form:"progress"`
	RecorderId       int  `json:"recorderId" from:"recorderId"`
	Recorder         SysUser `json:"recorder" form:"recorder" gorm:"ForiegnKey:RecorderId;AssociationForiegnKey:ID"`
	FeedbackContent  string `json:"feedbackContent" form:"feedbackContent"`
	FeedbackRate     int `json:"feedbackRate" form:"feedbackRate"`
}
