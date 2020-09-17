package service

import (
	"panta/global"
	"panta/model"
	"panta/model/request"
)

// @title    CreateLiftModel
// @description   create a LiftModel
// @param     liftModel               model.LiftModel
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateLiftModel(liftModel model.LiftModel) (err error) {
	err = global.PantaDb.Create(&liftModel).Error
	return err
}

// @title    DeleteLiftModel
// @description   delete a LiftModel
// @auth                     （2020/04/05  20:22）
// @param     liftModel               model.LiftModel
// @return                    error

func DeleteLiftModel(liftModel model.LiftModel) (err error) {
	err = global.PantaDb.Delete(liftModel).Error
	return err
}

// @title    DeleteLiftModelByIds
// @description   delete LiftModels
// @auth                     （2020/04/05  20:22）
// @param     liftModel               model.LiftModel
// @return                    error

func DeleteLiftModelByIds(ids request.IdsReq) (err error) {
	err = global.PantaDb.Delete(&[]model.LiftModel{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateLiftModel
// @description   update a LiftModel
// @param     liftModel          *model.LiftModel
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateLiftModel(liftModel *model.LiftModel) (err error) {
	err = global.PantaDb.Save(liftModel).Error
	return err
}

// @title    GetLiftModel
// @description   get the info of a LiftModel
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    LiftModel        LiftModel

func GetLiftModel(id uint) (err error, liftModel model.LiftModel) {
	err = global.PantaDb.Where("id = ?", id).Preload("Company").First(&liftModel).Error
	return
}

// @title    GetLiftModelInfoList
// @description   get LiftModel list by pagination, 分页获取LiftModel
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetLiftModelInfoList(info request.LiftModelSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.PantaDb.Model(&model.LiftModel{})
    var liftModels []model.LiftModel
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Company").Find(&liftModels).Error
	return err, liftModels, total
}