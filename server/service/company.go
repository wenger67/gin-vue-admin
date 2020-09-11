package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateCompany
// @description   create a Company
// @param     company               model.Company
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateCompany(company model.Company) (err error) {
	err = global.GVA_DB.Create(&company).Error
	return err
}

// @title    DeleteCompany
// @description   delete a Company
// @auth                     （2020/04/05  20:22）
// @param     company               model.Company
// @return                    error

func DeleteCompany(company model.Company) (err error) {
	err = global.GVA_DB.Delete(company).Error
	return err
}

// @title    DeleteCompanyByIds
// @description   delete Companys
// @auth                     （2020/04/05  20:22）
// @param     company               model.Company
// @return                    error

func DeleteCompanyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Company{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateCompany
// @description   update a Company
// @param     company          *model.Company
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateCompany(company *model.Company) (err error) {
	err = global.GVA_DB.Save(company).Error
	return err
}

// @title    GetCompany
// @description   get the info of a Company
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Company        Company

func GetCompany(id uint) (err error, company model.Company) {
	err = global.GVA_DB.Where("id = ?", id).First(&company).Error
	return
}

// @title    GetCompanyInfoList
// @description   get Company list by pagination, 分页获取Company
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetCompanyInfoList(info request.CompanySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	categoryId := info.ID
    // 创建db
	db := global.GVA_DB.Model(&model.Company{})
    var companys []model.Company

	if categoryId != 0 {
		db = db.Where("category_id = ?", categoryId);
	}

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Category").Preload("Category.Subject").Find(&companys).Error
	return err, companys, total
}