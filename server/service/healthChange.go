package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateHealthChange
// @description   create a HealthChange
// @param     healthChange               model.HealthChange
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateHealthChange(healthChange model.HealthChange) (err error) {
	err = global.PantaDb.Create(&healthChange).Error
	return err
}

// @title    DeleteHealthChange
// @description   delete a HealthChange
// @auth                     （2020/04/05  20:22）
// @param     healthChange               model.HealthChange
// @return                    error

func DeleteHealthChange(healthChange model.HealthChange) (err error) {
	err = global.PantaDb.Delete(healthChange).Error
	return err
}

// @title    DeleteHealthChangeByIds
// @description   delete HealthChanges
// @auth                     （2020/04/05  20:22）
// @param     healthChange               model.HealthChange
// @return                    error

func DeleteHealthChangeByIds(ids request.IdsReq) (err error) {
	err = global.PantaDb.Delete(&[]model.HealthChange{},"id in (?)",ids.Ids).Error
	return err
}

func DeleteHealthChangeByLiftIds(ids request.IdsReq) (err error) {
	err = global.PantaDb.Delete(&[]model.HealthChange{},"lift_id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateHealthChange
// @description   update a HealthChange
// @param     healthChange          *model.HealthChange
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateHealthChange(healthChange *model.HealthChange) (err error) {
	err = global.PantaDb.Save(healthChange).Error
	return err
}

// @title    GetHealthChange
// @description   get the info of a HealthChange
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    HealthChange        HealthChange

func GetHealthChange(id uint) (err error, healthChange model.HealthChange) {
	err = global.PantaDb.Where("id = ?", id).Preload("Lift").Preload("Dimension").First(&healthChange).Error
	return
}

// @title    GetHealthChangeList
// @description   get HealthChange list by pagination, 分页获取HealthChange
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetHealthChangeList(info request.HealthChangeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.PantaDb.Model(&model.HealthChange{})
    var healthChanges []model.HealthChange
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Lift").Preload("Dimension").Find(&healthChanges).Error
	return err, healthChanges, total
}