package model

import "github.com/jinzhu/gorm"

type Company struct {
	gorm.Model
	FullName string
	Alias       string
	LegalPerson string
	PhoneNumber string
	Status      string
	RegCode     string
	OrgCode     string
	CreditCode  string
	TaxCode     string
	Address     string
	CategoryId  int
	Category    Category
}