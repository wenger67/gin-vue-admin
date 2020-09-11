package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	FullName    string   `json:"fullName"`
	Alias       string   `json:"alias"`
	LegalPerson string   `json:"legalPerson"`
	PhoneNumber string   `json:"phone"`
	Status      string   `json:"status"`
	RegCode     string   `json:"regCode"`
	OrgCode     string   `json:"orgCode"`
	CreditCode  string   `json:"creditCode"`
	TaxCode     string   `json:"taxCode"`
	Address     string   `json:"address"`
	// belong to
	CategoryID  int      `json:"categoryId"`
	Category    Category `json:"category"`
	AdminId     int      `json:"adminId"`
}
