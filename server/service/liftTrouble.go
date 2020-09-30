package service

import (
	"panta/global"
	"panta/model"
	"panta/model/request"
	"panta/model/response"
	"panta/utils/enum"
	"time"
)

// @title    CreateLiftTrouble
// @description   create a LiftTrouble
// @param     liftTrouble               model.LiftTrouble
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateLiftTrouble(liftTrouble model.LiftTrouble) (err error) {
	err = global.PantaDb.Create(&liftTrouble).Error
	return err
}

// @title    DeleteLiftTrouble
// @description   delete a LiftTrouble
// @auth                     （2020/04/05  20:22）
// @param     liftTrouble               model.LiftTrouble
// @return                    error

func DeleteLiftTrouble(liftTrouble model.LiftTrouble) (err error) {
	err = global.PantaDb.Delete(liftTrouble).Error
	return err
}

// @title    DeleteLiftTroubleByIds
// @description   delete LiftTroubles
// @auth                     （2020/04/05  20:22）
// @param     liftTrouble               model.LiftTrouble
// @return                    error

func DeleteLiftTroubleByIds(ids request.IdsReq) (err error) {
	err = global.PantaDb.Delete(&[]model.LiftTrouble{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateLiftTrouble
// @description   update a LiftTrouble
// @param     liftTrouble          *model.LiftTrouble
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateLiftTrouble(liftTrouble *request.LiftTroubleUpdate) (err error) {
	updateMap := make(map[string]interface{})
	switch liftTrouble.Progress {
	case int(enum.TroubleCreated):
		updateMap["ResponseUserId"] = liftTrouble.ResponseUserId
		updateMap["ResponseTime"] = time.Now()
		updateMap["Progress"] = liftTrouble.Progress + 1
		break
	case int(enum.TroubleResponded):
		updateMap["SceneUserId"] = liftTrouble.SceneUserId
		updateMap["SceneTime"] = time.Now()
		updateMap["Progress"] = liftTrouble.Progress + 1
		break
	case int(enum.TroubleScened):
		updateMap["FixUserId"] = liftTrouble.FixUserId
		updateMap["FixTime"] = time.Now()
		updateMap["FixCategoryId"] = liftTrouble.FixCategoryId
		updateMap["ReasonCategoryId"] = liftTrouble.ReasonCategoryId
		updateMap["Content"] = liftTrouble.Content
		updateMap["Progress"] = liftTrouble.Progress + 1
		break
	case int(enum.TroubleFixed):
		updateMap["FeedbackContent"] = liftTrouble.FeedbackContent
		updateMap["FeedbackRate"] = liftTrouble.FeedbackRate
		updateMap["Progress"] = liftTrouble.Progress + 1
		break
	case int(enum.TroubleFeedback):
		updateMap["RecorderId"] = liftTrouble.RecorderId
		updateMap["Progress"] = liftTrouble.Progress + 1
		break
	default:
		global.PantaLog.Warning("unknown lift trouble progress, ", liftTrouble.Progress)

	}
	err = global.PantaDb.Model(model.LiftTrouble{}).Where("ID = ?", liftTrouble.ID).Updates(updateMap).Error
	return err
}

// @title    GetLiftTrouble
// @description   get the info of a LiftTrouble
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    LiftTrouble        LiftTrouble

func GetLiftTrouble(id uint) (err error, liftTrouble model.LiftTrouble) {
	err = global.PantaDb.Where("id = ?", id).Preload("Lift").Preload("FromCategory").
		Preload("StartUser").Preload("ResponseUser").Preload("SceneUser").
		Preload("FixUser").Preload("FixCategory").Preload("ReasonCategory").
		Preload("Recorder").Preload("Medias").First(&liftTrouble).Error
	return
}

// @title    GetLiftTroubleInfoList
// @description   get LiftTrouble list by pagination, 分页获取LiftTrouble
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetLiftTroubleInfoList(info request.LiftTroubleSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.PantaDb.Model(&model.LiftTrouble{})
    var liftTroubles []model.LiftTrouble

	if info.Key != "" {
		var res []response.LiftTroubleCountByProgress
		if info.Key == "progress" {
			db = db.Count(&total).Select("progress,Count(*) as total").Group("progress").
				Order("total desc, progress desc")
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
		err = db.Limit(limit).Offset(offset).Preload("Lift").Preload("FromCategory").
			Preload("StartUser").Preload("ResponseUser").Preload("SceneUser").
			Preload("FixUser").Preload("FixCategory").Preload("ReasonCategory").
			Preload("Medias").Preload("Recorder").Find(&liftTroubles).Error
		return err, liftTroubles, total
	}
}