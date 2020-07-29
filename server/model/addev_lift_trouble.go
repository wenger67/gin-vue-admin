package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type LiftTrouble struct {
	gorm.Model
	LiftId           int
	Lift             Lift
	FromCategoryId   int
	FromCategory     Category
	StartTime        time.Time
	StartUserId      int
	StartUser        UserAdmin
	ResponseTime     time.Time
	ResponseUserId   int
	ResponseUser     UserAdmin
	SceneTime        time.Time
	SceneUserId      int
	SceneUser        UserAdmin
	FixTime          time.Time
	FixUserId        int
	FixUser          UserAdmin
	FixCategoryId    int
	FixCategory      Category
	ReasonCategoryId int
	ReasonCategory   Category
	Content          string
	Progress         int
	RecorderId       int
	Recorder         UserAdmin
	FeedbackContent  string
	FeedbackRate     int
}
