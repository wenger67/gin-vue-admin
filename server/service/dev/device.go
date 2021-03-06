package dev

import (
	"panta/global"
	"panta/model"
)

func FindLift(id uint) (err error, lift model.Lift) {
	err = global.PantaDb.Where("id = ?", id).Preload("Installer").Preload("Installer.Category").Preload("Maintainer").
		Preload("Maintainer.Category").Preload("Checker").Preload("Checker.Category").
		Preload("Owner").Preload("Owner.Category").Preload("LiftModel").Preload("Category").
		Preload("Address").Preload("Address.Region").Preload("AdDevice").First(&lift).Error
	return
}

func CreateEvent(adDeviceEvent *model.AdDeviceEvent) (err error) {
	err = global.PantaDb.Create(adDeviceEvent).Error
	return err
}

func UpdateEvent(adDeviceEvent *model.AdDeviceEvent) (err error) {
	err = global.PantaDb.Save(adDeviceEvent).Error
	return err
}