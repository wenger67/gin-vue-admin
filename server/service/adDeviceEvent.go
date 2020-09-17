package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateAdDeviceEvent
// @description   create a AdDeviceEvent
// @param     adDeviceEvent               model.AdDeviceEvent
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateAdDeviceEvent(adDeviceEvent model.AdDeviceEvent) (err error) {
	err = global.PantaDb.Create(&adDeviceEvent).Error
	return err
}

// @title    DeleteAdDeviceEvent
// @description   delete a AdDeviceEvent
// @auth                     （2020/04/05  20:22）
// @param     adDeviceEvent               model.AdDeviceEvent
// @return                    error

func DeleteAdDeviceEvent(adDeviceEvent model.AdDeviceEvent) (err error) {
	err = global.PantaDb.Delete(adDeviceEvent).Error
	return err
}

// @title    DeleteAdDeviceEventByIds
// @description   delete AdDeviceEvents
// @auth                     （2020/04/05  20:22）
// @param     adDeviceEvent               model.AdDeviceEvent
// @return                    error

func DeleteAdDeviceEventByIds(ids request.IdsReq) (err error) {
	err = global.PantaDb.Delete(&[]model.AdDeviceEvent{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateAdDeviceEvent
// @description   update a AdDeviceEvent
// @param     adDeviceEvent          *model.AdDeviceEvent
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateAdDeviceEvent(adDeviceEvent *model.AdDeviceEvent) (err error) {
	err = global.PantaDb.Save(adDeviceEvent).Error
	return err
}

// @title    GetAdDeviceEvent
// @description   get the info of a AdDeviceEvent
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    AdDeviceEvent        AdDeviceEvent

func GetAdDeviceEvent(id uint) (err error, adDeviceEvent model.AdDeviceEvent) {
	err = global.PantaDb.Where("id = ?", id).Preload("Device").Preload("Type").First(&adDeviceEvent).Error
	return
}

// @title    GetAdDeviceEventInfoList
// @description   get AdDeviceEvent list by pagination, 分页获取AdDeviceEvent
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetAdDeviceEventInfoList(info request.AdDeviceEventSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.PantaDb.Model(&model.AdDeviceEvent{})
    var adDeviceEvents []model.AdDeviceEvent
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Device").Preload("Type").Find(&adDeviceEvents).Error
	return err, adDeviceEvents, total
}