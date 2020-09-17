package response

import "panta/model"

type SubjectResponse struct {
	Subject model.CategorySubject `json:"subject"`
}

type SubjectListResponse struct {
	SubjectList []model.CategorySubject `json:"subjectList"`
}