package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateAdDeviceConfig
// @description   create a AdDeviceConfig
// @param     adDeviceConfig               model.AdDeviceConfig
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateAdDeviceConfig(adDeviceConfig model.AdDeviceConfig) (err error) {
	err = global.GVA_DB.Create(&adDeviceConfig).Error
	return err
}

// @title    DeleteAdDeviceConfig
// @description   delete a AdDeviceConfig
// @auth                     （2020/04/05  20:22）
// @param     adDeviceConfig               model.AdDeviceConfig
// @return                    error

func DeleteAdDeviceConfig(adDeviceConfig model.AdDeviceConfig) (err error) {
	err = global.GVA_DB.Delete(adDeviceConfig).Error
	return err
}

// @title    DeleteAdDeviceConfigByIds
// @description   delete AdDeviceConfigs
// @auth                     （2020/04/05  20:22）
// @param     adDeviceConfig               model.AdDeviceConfig
// @return                    error

func DeleteAdDeviceConfigByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.AdDeviceConfig{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateAdDeviceConfig
// @description   update a AdDeviceConfig
// @param     adDeviceConfig          *model.AdDeviceConfig
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateAdDeviceConfig(adDeviceConfig *model.AdDeviceConfig) (err error) {
	err = global.GVA_DB.Save(adDeviceConfig).Error
	return err
}

// @title    GetAdDeviceConfig
// @description   get the info of a AdDeviceConfig
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    AdDeviceConfig        AdDeviceConfig

func GetAdDeviceConfig(id uint) (err error, adDeviceConfig model.AdDeviceConfig) {
	err = global.GVA_DB.Where("id = ?", id).First(&adDeviceConfig).Error
	return
}

// @title    GetAdDeviceConfigInfoList
// @description   get AdDeviceConfig list by pagination, 分页获取AdDeviceConfig
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetAdDeviceConfigInfoList(info request.AdDeviceConfigSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.AdDeviceConfig{})
    var adDeviceConfigs []model.AdDeviceConfig
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&adDeviceConfigs).Error
	return err, adDeviceConfigs, total
}