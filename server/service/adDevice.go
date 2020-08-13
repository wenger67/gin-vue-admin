package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateAdDevice
// @description   create a AdDevice
// @param     adDevice               model.AdDevice
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateAdDevice(adDevice model.AdDevice) (err error) {
	adDevice.LastOnlineTime = adDevice.InstallTime
	adDevice.LastOfflineTime = adDevice.InstallTime
	err = global.GVA_DB.Create(&adDevice).Error
	return err
}

// @title    DeleteAdDevice
// @description   delete a AdDevice
// @auth                     （2020/04/05  20:22）
// @param     adDevice               model.AdDevice
// @return                    error

func DeleteAdDevice(adDevice model.AdDevice) (err error) {
	err = global.GVA_DB.Delete(adDevice).Error
	return err
}

// @title    DeleteAdDeviceByIds
// @description   delete AdDevices
// @auth                     （2020/04/05  20:22）
// @param     adDevice               model.AdDevice
// @return                    error

func DeleteAdDeviceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.AdDevice{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateAdDevice
// @description   update a AdDevice
// @param     adDevice          *model.AdDevice
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateAdDevice(adDevice *model.AdDevice) (err error) {
	err = global.GVA_DB.Save(adDevice).Error
	return err
}

// @title    GetAdDevice
// @description   get the info of a AdDevice
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    AdDevice        AdDevice

func GetAdDevice(id uint) (err error, adDevice model.AdDevice) {
	err = global.GVA_DB.Where("id = ?", id).First(&adDevice).Error
	return
}

// @title    GetAdDeviceInfoList
// @description   get AdDevice list by pagination, 分页获取AdDevice
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetAdDeviceInfoList(info request.AdDeviceSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.AdDevice{})
    var adDevices []model.AdDevice
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Type").Preload("Factory").
		Preload("Status").Preload("Owners").Find(&adDevices).Error
	return err, adDevices, total
}