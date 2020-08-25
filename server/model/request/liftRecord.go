package request

import (
	"gin-vue-admin/model"
)

type LiftRecordSearch struct {
	model.LiftRecord
	PageInfo
}

// step 1 create
type LiftRecordCreate struct {
	LiftId     int       `json:"liftId" form:"liftId"`
	CategoryId int       `json:"categoryId" form:"categoryId"`
	WorkerId   int       `json:"workerId" form:"workerId"`
}

// step 2 fill
type LiftRecordFill struct {
	RecordId int       `json:"recordId" form:"recordId"`
	Images   string    `json:"images" form:"images"`
	Content  string    `json:"content" form:"content"`
}

// step3 review

type LiftRecordReview struct {
	RecordId   int `json:"recordId" form:"recordId"`
	RecorderId int `json:"recorderId" form:"recorderId"`
}
