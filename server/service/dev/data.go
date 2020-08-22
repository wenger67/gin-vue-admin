package dev

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

func CreateRunningData(adDeviceData model.AdDeviceData) (err error) {
	err = global.GVA_DB.Create(&adDeviceData).Error
	return err
}
