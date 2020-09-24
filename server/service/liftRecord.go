package service

import (
	"panta/global"
	"panta/model"
	"panta/model/request"
	"panta/model/response"
	"time"
)

// @title    CreateLiftRecord
// @description   create a LiftRecord
// @param     liftRecord               model.LiftRecord
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateLiftRecord(liftRecord model.LiftRecord) (err error) {
	liftRecord.Progress = 1
	err = global.PantaDb.Create(&liftRecord).Error
	return err
}

func FillLiftRecord(fill request.LiftRecordFill) (err error) {
	global.PantaLog.Debug(fill)
	db := global.PantaDb.Model(model.LiftRecord{}).Where("id = ?", fill.RecordId)
	err = db.Debug().Updates(map[string]interface{}{"content": fill.Content, "images": fill.Images, "EndTime": time.Now(), "progress": 2}).Error
	return
}

func ReviewLiftRecord(review request.LiftRecordReview) (err error) {
	db := global.PantaDb.Model(model.LiftRecord{}).Where("id = ?", review.RecordId)
	err = db.Updates(map[string]interface{}{"recorderId": review.RecorderId, "progress": 3}).Error
	return
}


// @title    DeleteLiftRecord
// @description   delete a LiftRecord
// @auth                     （2020/04/05  20:22）
// @param     liftRecord               model.LiftRecord
// @return                    error

func DeleteLiftRecord(liftRecord model.LiftRecord) (err error) {
	err = global.PantaDb.Delete(liftRecord).Error
	return err
}

// @title    DeleteLiftRecordByIds
// @description   delete LiftRecords
// @auth                     （2020/04/05  20:22）
// @param     liftRecord               model.LiftRecord
// @return                    error

func DeleteLiftRecordByIds(ids request.IdsReq) (err error) {
	err = global.PantaDb.Delete(&[]model.LiftRecord{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateLiftRecord
// @description   update a LiftRecord
// @param     liftRecord          *model.LiftRecord
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateLiftRecord(liftRecord *model.LiftRecord) (err error) {
	err = global.PantaDb.Save(liftRecord).Error
	return err
}

// @title    GetLiftRecord
// @description   get the info of a LiftRecord
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    LiftRecord        LiftRecord

func GetLiftRecord(id uint) (err error, liftRecord model.LiftRecord) {
	err = global.PantaDb.Where("id = ?", id).Preload("Lift").Preload("Category").Preload("Medias").
		Preload("Worker").Preload("Recorder").First(&liftRecord).Error
	return
}

// @title    GetLiftRecordInfoList
// @description   get LiftRecord list by pagination, 分页获取LiftRecord
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetLiftRecordInfoList(info request.LiftRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.PantaDb.Model(&model.LiftRecord{})
    var liftRecords []model.LiftRecord
	if info.Key != "" {
		var res []response.LiftRecordCountByType
		if info.Key == "type" {
			db = db.Count(&total).Select("category_id as type,Count(*) as total").Group("category_id").
				Order("total desc, category_id desc")
			if info.Threshold != 0 {
				db = db.Having("total > ?", info.Threshold)
			}
			if info.Limit != 0 {
				db = db.Limit(info.Limit)
			}
			err = db.Scan(&res).Error
		}
		return err, res, total
	} else {
		err = db.Count(&total).Error
		err = db.Limit(limit).Offset(offset).Preload("Lift").Preload("Category").Preload("Medias").
			Preload("Worker").Preload("Recorder").Find(&liftRecords).Error
		return err, liftRecords, total
	}
}