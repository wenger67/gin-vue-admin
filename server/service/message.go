package service

import (
	"panta/global"
	"panta/model"
	"panta/model/request"
)

// @title    CreateMessage
// @description   create a Message
// @param     message               model.Message
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateMessage(message model.Message) (err error) {
	err = global.PantaDb.Create(&message).Error
	return err
}

// @title    DeleteMessage
// @description   delete a Message
// @auth                     （2020/04/05  20:22）
// @param     message               model.Message
// @return                    error

func DeleteMessage(message model.Message) (err error) {
	err = global.PantaDb.Delete(message).Error
	return err
}

// @title    DeleteMessageByIds
// @description   delete Messages
// @auth                     （2020/04/05  20:22）
// @param     message               model.Message
// @return                    error

func DeleteMessageByIds(ids request.IdsReq) (err error) {
	err = global.PantaDb.Delete(&[]model.Message{}, "id in (?)", ids.Ids).Error
	return err
}

// @title    UpdateMessage
// @description   update a Message
// @param     message          *model.Message
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateMessage(message *model.Message) (err error) {
	err = global.PantaDb.Save(message).Error
	return err
}

// @title    GetMessage
// @description   get the info of a Message
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Message        Message

func GetMessage(id uint) (err error, message model.Message) {
	err = global.PantaDb.Where("id = ?", id).First(&message).Error
	return
}

// @title    GetMessageInfoList
// @description   get Message list by pagination, 分页获取Message
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetMessageInfoList(info request.MessageSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.PantaDb.Model(&model.Message{})
	var messages []model.Message
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("FromUser").Preload("TargetUser").Preload("Type").Find(&messages).Error
	return err, messages, total
}
