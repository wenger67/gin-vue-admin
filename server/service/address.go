package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateAddress
// @description   create a Address
// @param     address               model.Address
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateAddress(address model.Address) (err error) {
	err = global.GVA_DB.Create(&address).Error
	return err
}

// @title    DeleteAddress
// @description   delete a Address
// @auth                     （2020/04/05  20:22）
// @param     address               model.Address
// @return                    error

func DeleteAddress(address model.Address) (err error) {
	err = global.GVA_DB.Delete(address).Error
	return err
}

// @title    DeleteAddressByIds
// @description   delete Addresss
// @auth                     （2020/04/05  20:22）
// @param     address               model.Address
// @return                    error

func DeleteAddressByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Address{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateAddress
// @description   update a Address
// @param     address          *model.Address
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateAddress(address *model.Address) (err error) {
	err = global.GVA_DB.Save(address).Error
	return err
}

// @title    GetAddress
// @description   get the info of a Address
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Address        Address

func GetAddress(id uint) (err error, address model.Address) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Region").Preload("Tags").First(&address).Error
	return
}

// @title    GetAddressInfoList
// @description   get Address list by pagination, 分页获取Address
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetAddressInfoList(info request.AddressSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Address{})
    var addresss []model.Address

	if info.AddressName != "" {
		db = db.Where("address_name LIKE ?", "%" + info.AddressName + "%")
	}

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Region").Preload("Tags").Find(&addresss).Error
	return err, addresss, total
}