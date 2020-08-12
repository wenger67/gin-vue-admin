package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateAdDeviceData
// @description   create a AdDeviceData
// @param     adDeviceData               model.AdDeviceData
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateAdDeviceData(adDeviceData model.AdDeviceData) (err error) {
	err = global.GVA_DB.Create(&adDeviceData).Error
	return err
}

// @title    DeleteAdDeviceData
// @description   delete a AdDeviceData
// @auth                     （2020/04/05  20:22）
// @param     adDeviceData               model.AdDeviceData
// @return                    error

func DeleteAdDeviceData(adDeviceData model.AdDeviceData) (err error) {
	err = global.GVA_DB.Delete(adDeviceData).Error
	return err
}

// @title    DeleteAdDeviceDataByIds
// @description   delete AdDeviceDatas
// @auth                     （2020/04/05  20:22）
// @param     adDeviceData               model.AdDeviceData
// @return                    error

func DeleteAdDeviceDataByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.AdDeviceData{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateAdDeviceData
// @description   update a AdDeviceData
// @param     adDeviceData          *model.AdDeviceData
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateAdDeviceData(adDeviceData *model.AdDeviceData) (err error) {
	err = global.GVA_DB.Save(adDeviceData).Error
	return err
}

// @title    GetAdDeviceData
// @description   get the info of a AdDeviceData
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    AdDeviceData        AdDeviceData

func GetAdDeviceData(id uint) (err error, adDeviceData model.AdDeviceData) {
	err = global.GVA_DB.Where("id = ?", id).First(&adDeviceData).Error
	return
}

// @title    GetAdDeviceDataInfoList
// @description   get AdDeviceData list by pagination, 分页获取AdDeviceData
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetAdDeviceDataInfoList(info request.AdDeviceDataSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.AdDeviceData{})
    var adDeviceDatas []model.AdDeviceData
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&adDeviceDatas).Error
	return err, adDeviceDatas, total
}