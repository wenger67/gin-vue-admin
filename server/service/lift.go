package service

import (
	"panta/global"
	"panta/model"
	"panta/model/request"
	"panta/model/response"
)

// @title    CreateLift
// @description   create a Lift
// @param     lift               model.Lift
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateLift(lift *model.Lift) (err error) {
	err = global.PantaDb.Create(lift).Error
	return err
}

// @title    DeleteLift
// @description   delete a Lift
// @auth                     （2020/04/05  20:22）
// @param     lift               model.Lift
// @return                    error

func DeleteLift(lift model.Lift) (err error) {
	err = global.PantaDb.Delete(lift).Error
	return err
}

// @title    DeleteLiftByIds
// @description   delete Lifts
// @auth                     （2020/04/05  20:22）
// @param     lift               model.Lift
// @return                    error

func DeleteLiftByIds(ids request.IdsReq) (err error) {
	err = global.PantaDb.Delete(&[]model.Lift{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateLift
// @description   update a Lift
// @param     lift          *model.Lift
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateLift(lift *model.Lift) (err error) {
	err = global.PantaDb.Save(lift).Error
	return err
}

// @title    GetLift
// @description   get the info of a Lift
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Lift        Lift

func GetLift(id uint) (err error, lift model.Lift) {
	err = global.PantaDb.Where("id = ?", id).Preload("Installer").Preload("Installer.Admin").Preload("Installer." +
		"Category").Preload("Maintainer").Preload("Maintainer.Admin").Preload("Maintainer.Category").
		Preload("Checker").Preload("Checker.Admin").Preload("Checker.Category").
		Preload("Owner").Preload("Owner.Admin").Preload("Owner.Category").Preload("LiftModel").Preload("Category").
		Preload("Address").Preload("Address.Region").Preload("AdDevice").First(&lift).Error
	return
}

// @title    GetLiftInfoList
// @description   get Lift list by pagination, 分页获取Lift
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetLiftInfoList(info request.LiftSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.PantaDb.Model(&model.Lift{})
    var lifts []model.Lift
	var res []response.LiftCountByAddress

	if info.Key != "" {
		if info.Key == "address" {
			db.Select("address_id").Group("address_id").Count(&total)
			err = db.Select("address_id as address,Count(*) as total, addresses.address_name").
				Joins("left join addresses on addresses.id = lifts.address_id").Group("address_id").Scan(&res).Error
		}
		return err, res, total
	} else {
		err = db.Count(&total).Error
		err = db.Limit(limit).Offset(offset).Preload("Installer").Preload("Installer.Category").Preload("Maintainer").
			Preload("Maintainer.Category").Preload("Checker").Preload("Checker.Category").
			Preload("Owner").Preload("Owner.Category").Preload("LiftModel").Preload("Category").
			Preload("Address").Preload("AdDevice").Find(&lifts).Error
		return err, lifts, total
	}
}