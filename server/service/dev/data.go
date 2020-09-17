package dev

import (
	"panta/global"
	"panta/model"
)

func CreateRunningData(adDeviceData model.AdDeviceData) (err error) {
	err = global.PantaDb.Create(&adDeviceData).Error
	return err
}
