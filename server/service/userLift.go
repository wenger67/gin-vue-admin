package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateUserLift
// @description   create a UserLift
// @param     userLift               model.UserLift
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateUserLift(userLift model.UserLift) (err error) {
	err = global.PantaDb.Create(&userLift).Error
	return err
}

// @title    DeleteUserLift
// @description   delete a UserLift
// @auth                     （2020/04/05  20:22）
// @param     userLift               model.UserLift
// @return                    error

func DeleteUserLift(userLift model.UserLift) (err error) {
	err = global.PantaDb.Delete(userLift).Error
	return err
}

// @title    DeleteUserLiftByIds
// @description   delete UserLifts
// @auth                     （2020/04/05  20:22）
// @param     userLift               model.UserLift
// @return                    error

func DeleteUserLiftByIds(ids request.IdsReq) (err error) {
	err = global.PantaDb.Delete(&[]model.UserLift{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateUserLift
// @description   update a UserLift
// @param     userLift          *model.UserLift
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateUserLift(userLift *model.UserLift) (err error) {
	err = global.PantaDb.Save(userLift).Error
	return err
}

// @title    GetUserLift
// @description   get the info of a UserLift
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    UserLift        UserLift

func GetUserLift(id uint) (err error, userLift model.UserLift) {
	err = global.PantaDb.Where("id = ?", id).Preload("Lift").Preload("User").Preload("Category").First(&userLift).Error
	return
}

// @title    GetUserLiftList
// @description   get UserLift list by pagination, 分页获取UserLift
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetUserLiftList(info request.UserLiftSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.PantaDb.Model(&model.UserLift{})
    var userLift []model.UserLift
    // 如果有条件搜索 下方会自动创建搜索语句
	if info.LiftId != 0 {
		db = db.Where("lift_id = ?", info.LiftId)
	}
	if info.CategoryId != 0 {
		db = db.Where("category_id = ?", info.CategoryId)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Lift").Preload("User").Preload("Category").Find(&userLift).Error
	return err, userLift, total
}