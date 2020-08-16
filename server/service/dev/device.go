package dev

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

func FindLift(id uint) (err error, lift model.Lift) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Installer").Preload("Installer.Category").Preload("Maintainer").
		Preload("Maintainer.Category").Preload("Checker").Preload("Checker.Category").
		Preload("Owner").Preload("Owner.Category").Preload("LiftModel").Preload("Category").
		Preload("Address").Preload("Address.Region").Preload("AdDevice").First(&lift).Error
	return
}
