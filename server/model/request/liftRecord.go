package request

import (
	"panta/model"
)

type LiftRecordSearch struct {
	model.LiftRecord
	PageInfo
	RankInfo
}

// step 1 create
type LiftRecordCreate struct {
	LiftId     uint       `json:"liftId" form:"liftId"`
	CategoryId uint       `json:"categoryId" form:"categoryId"`
}

type LiftRecordUpdate struct {
	model.LiftRecord
}
