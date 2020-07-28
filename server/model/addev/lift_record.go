package addev

import (
	"github.com/jinzhu/gorm"
	"time"
)

type LiftRecord struct {
	gorm.Model
	LiftId int
	Lift Lift
	CategoryId int
	Category Category
	Images string
	Content string
	StartTime time.Time
	EndTime time.Time
	WorkerId int
	Worker UserAdmin
	RecorderId int
	Recorder UserAdmin
}
