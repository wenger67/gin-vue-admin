package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateLiftChange
// @description   create a LiftChange
// @param     liftChange               model.LiftChange
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateLiftChange(liftChange model.LiftChange) (err error) {
	err = global.PantaDb.Create(&liftChange).Error
	return err
}

// @title    DeleteLiftChange
// @description   delete a LiftChange
// @auth                     （2020/04/05  20:22）
// @param     liftChange               model.LiftChange
// @return                    error

func DeleteLiftChange(liftChange model.LiftChange) (err error) {
	err = global.PantaDb.Delete(liftChange).Error
	return err
}

// @title    DeleteLiftChangeByIds
// @description   delete LiftChanges
// @auth                     （2020/04/05  20:22）
// @param     liftChange               model.LiftChange
// @return                    error

func DeleteLiftChangeByIds(ids request.IdsReq) (err error) {
	err = global.PantaDb.Delete(&[]model.LiftChange{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateLiftChange
// @description   update a LiftChange
// @param     liftChange          *model.LiftChange
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateLiftChange(liftChange *model.LiftChange) (err error) {
	err = global.PantaDb.Save(liftChange).Error
	return err
}

// @title    GetLiftChange
// @description   get the info of a LiftChange
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    LiftChange        LiftChange

func GetLiftChange(id uint) (err error, liftChange model.LiftChange) {
	err = global.PantaDb.Where("id = ?", id).Preload("Lift").First(&liftChange).Error
	return
}

// @title    GetLiftChangeInfoList
// @description   get LiftChange list by pagination, 分页获取LiftChange
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetLiftChangeInfoList(info request.LiftChangeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.PantaDb.Model(&model.LiftChange{})
    var liftChanges []model.LiftChange
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Lift").Find(&liftChanges).Error
	return err, liftChanges, total
}