package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateLift
// @description   create a Lift
// @param     lift               model.Lift
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateLift(lift *model.Lift) (err error) {
	err = global.GVA_DB.Create(lift).Error
	return err
}

// @title    DeleteLift
// @description   delete a Lift
// @auth                     （2020/04/05  20:22）
// @param     lift               model.Lift
// @return                    error

func DeleteLift(lift model.Lift) (err error) {
	err = global.GVA_DB.Delete(lift).Error
	return err
}

// @title    DeleteLiftByIds
// @description   delete Lifts
// @auth                     （2020/04/05  20:22）
// @param     lift               model.Lift
// @return                    error

func DeleteLiftByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Lift{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateLift
// @description   update a Lift
// @param     lift          *model.Lift
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateLift(lift *model.Lift) (err error) {
	err = global.GVA_DB.Save(lift).Error
	return err
}

// @title    GetLift
// @description   get the info of a Lift
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Lift        Lift

func GetLift(id uint) (err error, lift model.Lift) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Installer").Preload("Installer.Category").Preload("Maintainer").
		Preload("Maintainer.Category").Preload("Checker").Preload("Checker.Category").
		Preload("Owner").Preload("Owner.Category").Preload("LiftModel").Preload("Category").
		Preload("Address").Preload("Address.Region").Preload("AdDevice").First(&lift).Error
	return
}

// @title    GetLiftInfoList
// @description   get Lift list by pagination, 分页获取Lift
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetLiftInfoList(info request.LiftSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Lift{})
    var lifts []model.Lift
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Installer").Preload("Installer.Category").Preload("Maintainer").
		Preload("Maintainer.Category").Preload("Checker").Preload("Checker.Category").
		Preload("Owner").Preload("Owner.Category").Preload("LiftModel").Preload("Category").
		Preload("Address").Preload("AdDevice").Find(&lifts).Error
	return err, lifts, total
}