package service

import (
	"errors"
	"panta/global"
	"panta/model"
	"panta/model/request"
	"panta/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// @title    Register
// @description   register, 用户注册
// @auth                     （2020/04/05  20:22）
// @param     u               model.SysUser
// @return    err             error
// @return    userInter       *SysUser

func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	// 判断用户名是否注册
	notRegister := errors.Is(global.PantaDb.Where("phone_number = ?", u.PhoneNumber).First(&user).Error, gorm.ErrRecordNotFound)
	// notRegister为false表明读取到了 不能注册
	if !notRegister {
		return errors.New("用户名已注册"), userInter
	} else {
		// 否则 附加uuid 密码md5简单加密 注册
		u.Password = utils.MD5V([]byte(u.Password))
		u.UUID = uuid.NewV4()
		err = global.PantaDb.Create(&u).Error
	}
	return err, u
}

func CreateUser(u model.SysUser) (err error, user model.SysUser) {
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	// TODO create avatar

	err = global.PantaDb.Create(&u).Error
	return err, u
}

// @title    Login
// @description   login, 用户登录
// @auth                     （2020/04/05  20:22）
// @param     u               *model.SysUser
// @return    err             error
// @return    userInter       *SysUser

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.PantaDb.Where("phone_number = ? AND password = ?", u.PhoneNumber, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

// @title    ChangePassword
// @description   change the password of a certain user, 修改用户密码
// @auth                     （2020/04/05  20:22）
// @param     u               *model.SysUser
// @param     newPassword     string
// @return    err             error
// @return    userInter       *SysUser

func ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.PantaDb.Where("phone_number = ? AND password = ?", u.PhoneNumber, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

// @title    GetInfoList
// @description   get user list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func GetUserInfoList(params request.SearchUserParams) (err error, list interface{}, total int64) {
	limit := params.PageInfo.PageSize
	offset := params.PageInfo.PageSize * (params.PageInfo.Page - 1)
	db := global.PantaDb.Model(&model.SysUser{})
	var userList []model.SysUser
	if params.CompanyId != 0 {
		db = db.Where("company_id = ?", params.CompanyId)
	}
	err = db.Count(&total).Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return err, userList, total
}

func GetUser(id request.GetById) (err error, user model.SysUser) {
	err = global.PantaDb.Where("id = ?", id.Id).Preload("Authority").First(&user).Error
	return
}

// @title    SetUserAuthority
// @description   set the authority of a certain user, 设置一个用户的权限
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     authorityId     string
// @return    err             error

func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = global.PantaDb.Where("uuid = ?", uuid).First(&model.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

// @title    SetUserAuthority
// @description   set the authority of a certain user, 设置一个用户的权限
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     authorityId     string
// @return    err             error

func DeleteUser(id float64) (err error) {
	var user model.SysUser
	err = global.PantaDb.Where("id = ?", id).Delete(&user).Error
	return err
}

func DeleteUserList(ids []int) (err error) {
	err = global.PantaDb.Delete(&[]model.SysUser{}, "id in (?)", ids).Error
	return err
}

// @title    UploadHeaderImg
// @description   upload avatar, 用户头像上传更新地址
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     filePath        string
// @return    err             error
// @return    userInter       *SysUser

func UploadHeaderImg(uuid uuid.UUID, filePath string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	err = global.PantaDb.Where("uuid = ?", uuid).First(&user).Update("header_img", filePath).First(&user).Error
	return err, &user
}
