package request

import "panta/model"

type SearchSubjectParams struct {
	model.CategorySubject
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc bool `json:"desc"`
}