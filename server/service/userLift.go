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
	err = global.GVA_DB.Create(&userLift).Error
	return err
}

// @title    DeleteUserLift
// @description   delete a UserLift
// @auth                     （2020/04/05  20:22）
// @param     userLift               model.UserLift
// @return                    error

func DeleteUserLift(userLift model.UserLift) (err error) {
	err = global.GVA_DB.Delete(userLift).Error
	return err
}

// @title    DeleteUserLiftByIds
// @description   delete UserLifts
// @auth                     （2020/04/05  20:22）
// @param     userLift               model.UserLift
// @return                    error

func DeleteUserLiftByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.UserLift{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateUserLift
// @description   update a UserLift
// @param     userLift          *model.UserLift
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateUserLift(userLift *model.UserLift) (err error) {
	err = global.GVA_DB.Save(userLift).Error
	return err
}

// @title    GetUserLift
// @description   get the info of a UserLift
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    UserLift        UserLift

func GetUserLift(id uint) (err error, userLift model.UserLift) {
	err = global.GVA_DB.Where("id = ?", id).First(&userLift).Error
	return
}

// @title    GetUserLiftInfoList
// @description   get UserLift list by pagination, 分页获取UserLift
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetUserLiftInfoList(info request.UserLiftSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.UserLift{})
    var userLift []model.UserLift
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&userLift).Error
	return err, userLift, total
}