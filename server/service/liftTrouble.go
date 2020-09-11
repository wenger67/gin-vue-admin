package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"time"
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
	updateMap := make(map[string]interface{})
	switch liftTrouble.Progress {
	case 1:
		updateMap["ResponseUserId"] = liftTrouble.ResponseUserId
		updateMap["ResponseTime"] = time.Now()
		updateMap["Progress"] = liftTrouble.Progress + 1
		break
	case 2:
		updateMap["SceneUserId"] = liftTrouble.SceneUserId
		updateMap["SceneTime"] = time.Now()
		updateMap["Progress"] = liftTrouble.Progress + 1
		break
	case 3:
		updateMap["FixUserId"] = liftTrouble.FixUserId
		updateMap["FixTime"] = time.Now()
		updateMap["FixCategoryId"] = liftTrouble.FixCategoryId
		updateMap["ReasonCategoryId"] = liftTrouble.ReasonCategoryId
		updateMap["Content"] = liftTrouble.Content
		updateMap["Progress"] = liftTrouble.Progress + 1
		break
	case 4:
		updateMap["FeedbackContent"] = liftTrouble.FeedbackContent
		updateMap["FeedbackRate"] = liftTrouble.FeedbackRate
		updateMap["Progress"] = liftTrouble.Progress + 1
		break
	case 5:
		updateMap["RecorderId"] = liftTrouble.RecorderId
		updateMap["Progress"] = liftTrouble.Progress + 1
		break
	default:
		global.GVA_LOG.Warning("unknown lift trouble progress, ", liftTrouble.Progress)

	}
	err = global.GVA_DB.Model(model.LiftTrouble{}).Where("ID = ?", liftTrouble.ID).Updates(updateMap).Error
	return err
}

// @title    GetLiftTrouble
// @description   get the info of a LiftTrouble
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    LiftTrouble        LiftTrouble

func GetLiftTrouble(id uint) (err error, liftTrouble model.LiftTrouble) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Lift").Preload("FromCategory").
		Preload("StartUser").Preload("ResponseUser").Preload("SceneUser").
		Preload("FixUser").Preload("FixCategory").Preload("ReasonCategory").
		Preload("Recorder").First(&liftTrouble).Error
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
	db := global.GVA_DB.Model(&model.LiftTrouble{})
    var liftTroubles []model.LiftTrouble
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Lift").Preload("FromCategory").
		Preload("StartUser").Preload("ResponseUser").Preload("SceneUser").
		Preload("FixUser").Preload("FixCategory").Preload("ReasonCategory").
		Preload("Recorder").Find(&liftTroubles).Error
	return err, liftTroubles, total
}