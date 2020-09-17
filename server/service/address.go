package service

import (
	"panta/global"
	"panta/model"
	"panta/model/request"
)

// @title    CreateAddress
// @description   create a Address
// @param     address               model.Address
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateAddress(address model.Address) (err error) {
	err = global.PantaDb.Create(&address).Error
	return err
}

// @title    DeleteAddress
// @description   delete a Address
// @auth                     （2020/04/05  20:22）
// @param     address               model.Address
// @return                    error

func DeleteAddress(address model.Address) (err error) {
	err = global.PantaDb.Delete(address).Error
	return err
}

// @title    DeleteAddressByIds
// @description   delete Addresss
// @auth                     （2020/04/05  20:22）
// @param     address               model.Address
// @return                    error

func DeleteAddressByIds(ids request.IdsReq) (err error) {
	err = global.PantaDb.Delete(&[]model.Address{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateAddress
// @description   update a Address
// @param     address          *model.Address
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateAddress(address *model.Address) (err error) {
	err = global.PantaDb.Save(address).Error
	return err
}

// @title    GetAddress
// @description   get the info of a Address
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Address        Address

func GetAddress(id uint) (err error, address model.Address) {
	err = global.PantaDb.Where("id = ?", id).Preload("Region").Preload("Tags").First(&address).Error
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
	db := global.PantaDb.Model(&model.Address{})
    var addresses []model.Address

	if info.AddressName != "" {
		db = db.Where("address_name LIKE ?", "%" + info.AddressName + "%")
	}

	err = db.Count(&total).Limit(limit).Offset(offset).Preload("Region").Preload("Tags").Find(&addresses).Error
	return err, addresses, total
}