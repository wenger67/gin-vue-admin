package service

import (
	"panta/global"
	"panta/model"
	"panta/model/request"
)

// @title    CreateAdDevice
// @description   create a Device
// @param     adDevice               model.Device
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateAdDevice(adDevice model.Device) (err error) {
	adDevice.LastOnlineTime = adDevice.InstallTime
	adDevice.LastOfflineTime = adDevice.InstallTime
	err = global.PantaDb.Create(&adDevice).Error
	return err
}

// @title    DeleteAdDevice
// @description   delete a Device
// @auth                     （2020/04/05  20:22）
// @param     adDevice               model.Device
// @return                    error

func DeleteAdDevice(adDevice model.Device) (err error) {
	err = global.PantaDb.Delete(adDevice).Error
	return err
}

// @title    DeleteAdDeviceByIds
// @description   delete AdDevices
// @auth                     （2020/04/05  20:22）
// @param     adDevice               model.Device
// @return                    error

func DeleteAdDeviceByIds(ids request.IdsReq) (err error) {
	err = global.PantaDb.Delete(&[]model.Device{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateAdDevice
// @description   update a Device
// @param     adDevice          *model.Device
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateAdDevice(adDevice *model.Device) (err error) {
	err = global.PantaDb.Save(adDevice).Error
	return err
}

// @title    GetAdDevice
// @description   get the info of a Device
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Device        Device

func GetAdDevice(id uint) (err error, adDevice model.Device) {
	err = global.PantaDb.Where("id = ?", id).First(&adDevice).Error
	return
}

// @title    GetAdDeviceInfoList
// @description   get Device list by pagination, 分页获取AdDevice
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetAdDeviceInfoList(info request.AdDeviceSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.PantaDb.Model(&model.Device{})
    var adDevices []model.Device
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Type").Preload("Factory").
		Preload("Status").Preload("Owners").Find(&adDevices).Error
	return err, adDevices, total
}