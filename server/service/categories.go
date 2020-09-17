package service

import (
	"panta/global"
	"panta/model"
	"panta/model/request"
)

// @title    CreateCategories
// @description   create a Categories
// @param     category               model.Categories
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateCategories(category model.Category) (err error) {
	err = global.PantaDb.Create(&category).Error
	return err
}

// @title    DeleteCategories
// @description   delete a Categories
// @auth                     （2020/04/05  20:22）
// @param     category               model.Categories
// @return                    error

func DeleteCategories(category model.Category) (err error) {
	err = global.PantaDb.Delete(category).Error
	return err
}

// @title    DeleteCategoriesByIds
// @description   delete Categoriess
// @auth                     （2020/04/05  20:22）
// @param     category               model.Categories
// @return                    error

func DeleteCategoriesByIds(ids request.IdsReq) (err error) {
	err = global.PantaDb.Delete(&[]model.Category{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateCategories
// @description   update a Categories
// @param     category          *model.Categories
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateCategories(category *model.Category) (err error) {
	err = global.PantaDb.Save(category).Error
	return err
}

// @title    GetCategories
// @description   get the info of a Categories
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Categories        Categories

func GetCategories(id uint) (err error, category model.Category) {
	err = global.PantaDb.Where("id = ?", id).Preload("Subject").First(&category).Error
	return
}

// @title    GetCategoriesInfoList
// @description   get Categories list by pagination, 分页获取Categories
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetCategoriesInfoList(info request.CategorySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	subjectId := info.ID
	db := global.PantaDb.Model(&model.Category{})
    var categories []model.Category
	if subjectId != 0 {
		db = db.Where("category_subject_id = ?", subjectId)
	}
	err = db.Count(&total).Limit(limit).Offset(offset).Preload("Subject").Find(&categories).Error
	return err, categories, total
}