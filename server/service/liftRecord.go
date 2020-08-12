package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateLiftRecord
// @description   create a LiftRecord
// @param     liftRecord               model.LiftRecord
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateLiftRecord(liftRecord model.LiftRecord) (err error) {
	err = global.GVA_DB.Create(&liftRecord).Error
	return err
}

// @title    DeleteLiftRecord
// @description   delete a LiftRecord
// @auth                     （2020/04/05  20:22）
// @param     liftRecord               model.LiftRecord
// @return                    error

func DeleteLiftRecord(liftRecord model.LiftRecord) (err error) {
	err = global.GVA_DB.Delete(liftRecord).Error
	return err
}

// @title    DeleteLiftRecordByIds
// @description   delete LiftRecords
// @auth                     （2020/04/05  20:22）
// @param     liftRecord               model.LiftRecord
// @return                    error

func DeleteLiftRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.LiftRecord{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateLiftRecord
// @description   update a LiftRecord
// @param     liftRecord          *model.LiftRecord
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateLiftRecord(liftRecord *model.LiftRecord) (err error) {
	err = global.GVA_DB.Save(liftRecord).Error
	return err
}

// @title    GetLiftRecord
// @description   get the info of a LiftRecord
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    LiftRecord        LiftRecord

func GetLiftRecord(id uint) (err error, liftRecord model.LiftRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&liftRecord).Error
	return
}

// @title    GetLiftRecordInfoList
// @description   get LiftRecord list by pagination, 分页获取LiftRecord
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetLiftRecordInfoList(info request.LiftRecordSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.LiftRecord{})
    var liftRecords []model.LiftRecord
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&liftRecords).Error
	return err, liftRecords, total
}