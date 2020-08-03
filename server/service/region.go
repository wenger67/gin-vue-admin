package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateRegion
// @description   create a Region
// @param     region               model.Region
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateRegion(region model.Region) (err error) {
	err = global.GVA_DB.Create(&region).Error
	return err
}

// @title    DeleteRegion
// @description   delete a Region
// @auth                     （2020/04/05  20:22）
// @param     region               model.Region
// @return                    error

func DeleteRegion(region model.Region) (err error) {
	err = global.GVA_DB.Delete(region).Error
	return err
}

// @title    DeleteRegionByIds
// @description   delete Regions
// @auth                     （2020/04/05  20:22）
// @param     region               model.Region
// @return                    error

func DeleteRegionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Region{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateRegion
// @description   update a Region
// @param     region          *model.Region
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateRegion(region *model.Region) (err error) {
	err = global.GVA_DB.Save(region).Error
	return err
}

// @title    GetRegion
// @description   get the info of a Region
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Region        Region

func GetRegion(id uint) (err error, region model.Region) {
	err = global.GVA_DB.Where("id = ?", id).First(&region).Error
	return
}

// @title    GetRegionInfoList
// @description   get Region list by pagination, 分页获取Region
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetRegionInfoList(info request.RegionSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Region{})
    var regions []model.Region
	if info.GroupKey == "province" {
		db = db.Select("province").Where("province LIKE ?", "%" + info.Province + "%").Group(info.GroupKey)
	}

	if info.GroupKey == "city" {
		db = db.Select("city").Where("province = ? AND city LIKE ?", info.Province, "%" + info.City + "%").Group(info.GroupKey)
	}

	if info.GroupKey == "district" {
		db = db.Select("district").Where("city = ? AND district LIKE ?", info.City, "%" + info.District + "%").Group(info.GroupKey)
	}
	err = db.Limit(limit).Offset(offset).Find(&regions).Error
	return err, regions, total
}