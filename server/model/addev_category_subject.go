package model

import "github.com/jinzhu/gorm"

type CategorySubject struct {
	gorm.Model
	SubjectName string `gorm:"comment:'类别主体名'"`
}
