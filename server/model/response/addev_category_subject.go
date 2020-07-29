package response

import "gin-vue-admin/model"

type SubjectResponse struct {
	Subject model.CategorySubject `json:"subject"`
}

type SubjectListResponse struct {
	SubjectList []model.CategorySubject `json:"subjectList"`
}