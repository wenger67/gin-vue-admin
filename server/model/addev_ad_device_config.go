package model

import "gorm.io/gorm"

type AdDeviceConfig struct {
	gorm.Model
	Key string `json:"key" form:"key"`
	Value string `json:"value" form:"value"`
	Comment string `json:"comment" form:"comment" grom:"comment:'说明'"`
}
