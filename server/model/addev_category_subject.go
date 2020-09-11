package model

import "gorm.io/gorm"

type CategorySubject struct {
	gorm.Model
	SubjectName string `json:"subjectName" gorm:"comment:'类别主体名'"`
}
