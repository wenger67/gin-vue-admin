package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateHealthSystem
// @description   create a HealthSystem
// @param     healthSystem               model.HealthSystem
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateHealthSystem(healthSystem model.HealthSystem) (err error) {
	err = global.GVA_DB.Create(&healthSystem).Error
	return err
}

// @title    DeleteHealthSystem
// @description   delete a HealthSystem
// @auth                     （2020/04/05  20:22）
// @param     healthSystem               model.HealthSystem
// @return                    error

func DeleteHealthSystem(healthSystem model.HealthSystem) (err error) {
	err = global.GVA_DB.Delete(healthSystem).Error
	return err
}

// @title    DeleteHealthSystemByIds
// @description   delete HealthSystems
// @auth                     （2020/04/05  20:22）
// @param     healthSystem               model.HealthSystem
// @return                    error

func DeleteHealthSystemByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.HealthSystem{},"id in (?)",ids.Ids).Error
	return err
}

func DeleteHealthSystemByLiftIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.HealthSystem{},"lift_id in (?)",ids.Ids).Error
	return err
}


// @title    UpdateHealthSystem
// @description   update a HealthSystem
// @param     healthSystem          *model.HealthSystem
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateHealthSystem(healthSystem *model.HealthSystem) (err error) {
	err = global.GVA_DB.Save(healthSystem).Error
	return err
}

// @title    GetHealthSystem
// @description   get the info of a HealthSystem
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    HealthSystem        HealthSystem

func GetHealthSystem(id uint) (err error, healthSystem model.HealthSystem) {
	err = global.GVA_DB.Where("id = ?", id).First(&healthSystem).Error
	return
}

// @title    GetHealthSystemInfoList
// @description   get HealthSystem list by pagination, 分页获取HealthSystem
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetHealthSystemInfoList(info request.HealthSystemSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.HealthSystem{})
    var healthSystems []model.HealthSystem
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&healthSystems).Error
	return err, healthSystems, total
}