package model

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Subject           CategorySubject `json:"categorySubject" gorm:"ForeignKey:CategorySubjectId;AssociationForeignKey:ID;comment:'Category Subject'"`
	CategorySubjectId int             `json:"categorySubjectId"`
	CategoryName      string          `json:"categoryName"`
}
