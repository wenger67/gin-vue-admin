package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gorm.io/gorm"
)

func CreateSubject(subject model.CategorySubject) (err error) {
	findOne := global.GVA_DB.Where("subject_name = ?", subject.SubjectName).Find(&model.CategorySubject{}).Error
	if findOne == nil {
		 return errors.New("same subject name exists")
	} else {
		err = global.GVA_DB.Create(&subject).Error
	}
	return err
}

func DeleteSubject(subject model.CategorySubject) (err error) {
	err = global.GVA_DB.Delete(subject).Error
	// todo delete relative logic
	return err
}

func GetSubjectList(subject model.CategorySubject, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.CategorySubject{})
	var subjectList []model.CategorySubject

	if subject.SubjectName != "" {
		db = db.Where("subject_name LIKE ?", "%" + subject.SubjectName + "%")
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, subjectList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Find(&subjectList).Error
		} else {
			err = db.Order("id").Find(&subjectList).Error
		}
	}
	return err, subjectList, total
}

func GetSubjectById(id float64) (err error, subject model.CategorySubject) {
	err = global.GVA_DB.Where("id = ?", id).First(&subject).Error
	return
}

func UpdateSubject(subject model.CategorySubject) (err error) {
	var origin model.CategorySubject
	err = global.GVA_DB.Where("id = ?", subject.ID).First(&origin).Error

	if err != nil {
		return err
	} else {
		if origin.SubjectName != subject.SubjectName {
			flag := errors.Is(global.GVA_DB.Where("subject_name = ?",
				subject.SubjectName).Find(&model.CategorySubject{}).Error, gorm.ErrRecordNotFound)
			if !flag {
				return errors.New("same subject name exist!")
			}
		} else {
			return errors.New("same subject name, nothing changed!")
		}

		err = global.GVA_DB.Save(&subject).Error
	}
	return err
}

func GetAllSubjects() (err error, subjects []model.CategorySubject)  {
	err = global.GVA_DB.Find(&subjects).Error
	return
}