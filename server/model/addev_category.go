package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	// belong to
	Subject           CategorySubject `json:"categorySubject" gorm:"foreignKey:CategorySubjectId"`
	CategorySubjectId uint             `json:"categorySubjectId"`
	CategoryName      string          `json:"categoryName"`
}
