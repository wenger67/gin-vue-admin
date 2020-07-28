package addev

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	CategorySubjectId int
	CategorySubject CategorySubject
	CategoryName string
}
