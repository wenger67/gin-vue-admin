package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateLiftTrouble
// @description   create a LiftTrouble
// @param     liftTrouble               model.LiftTrouble
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateLiftTrouble(liftTrouble model.LiftTrouble) (err error) {
	err = global.GVA_DB.Create(&liftTrouble).Error
	return err
}

// @title    DeleteLiftTrouble
// @description   delete a LiftTrouble
// @auth                     （2020/04/05  20:22）
// @param     liftTrouble               model.LiftTrouble
// @return                    error

func DeleteLiftTrouble(liftTrouble model.LiftTrouble) (err error) {
	err = global.GVA_DB.Delete(liftTrouble).Error
	return err
}

// @title    DeleteLiftTroubleByIds
// @description   delete LiftTroubles
// @auth                     （2020/04/05  20:22）
// @param     liftTrouble               model.LiftTrouble
// @return                    error

func DeleteLiftTroubleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.LiftTrouble{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateLiftTrouble
// @description   update a LiftTrouble
// @param     liftTrouble          *model.LiftTrouble
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateLiftTrouble(liftTrouble *model.LiftTrouble) (err error) {
	err = global.GVA_DB.Save(liftTrouble).Error
	return err
}

// @title    GetLiftTrouble
// @description   get the info of a LiftTrouble
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    LiftTrouble        LiftTrouble

func GetLiftTrouble(id uint) (err error, liftTrouble model.LiftTrouble) {
	err = global.GVA_DB.Where("id = ?", id).First(&liftTrouble).Error
	return
}

// @title    GetLiftTroubleInfoList
// @description   get LiftTrouble list by pagination, 分页获取LiftTrouble
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetLiftTroubleInfoList(info request.LiftTroubleSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.LiftTrouble{})
    var liftTroubles []model.LiftTrouble
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&liftTroubles).Error
	return err, liftTroubles, total
}