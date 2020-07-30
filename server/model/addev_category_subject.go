package model

import "github.com/jinzhu/gorm"

type CategorySubject struct {
	gorm.Model
	SubjectName string `json:"subjectName" gorm:"comment:'类别主体名'"`
}
