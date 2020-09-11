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
	LiftId     uint       `json:"liftId" form:"liftId"`
	CategoryId uint       `json:"categoryId" form:"categoryId"`
	WorkerId   uint       `json:"workerId" form:"workerId"`
}

// step 2 fill
type LiftRecordFill struct {
	RecordId uint       `json:"recordId" form:"recordId"`
	Images   string    `json:"images" form:"images"`
	Content  string    `json:"content" form:"content"`
}

// step3 review

type LiftRecordReview struct {
	RecordId   uint `json:"recordId" form:"recordId"`
	RecorderId uint `json:"recorderId" form:"recorderId"`
}
