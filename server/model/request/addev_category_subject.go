package request

import "gin-vue-admin/model"

type SearchSubjectParams struct {
	model.CategorySubject
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc bool `json:"desc"`
}