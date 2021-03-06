package service

import (
	"errors"
	"panta/global"
	"panta/model"
	"panta/model/request"
	"panta/model/response"
	"gorm.io/gorm"
	"strconv"
)

// @title    CreateAuthority
// @description   创建一个角色
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
// @return    authority       model.SysAuthority

func CreateAuthority(auth model.SysAuthority) (err error, authority model.SysAuthority) {
	var authorityBox model.SysAuthority
	notHas := errors.Is(global.PantaDb.Where("authority_id = ?", auth.AuthorityId).Find(&authorityBox).Error, gorm.ErrRecordNotFound)
	if !notHas {
		return errors.New("存在相同角色id"), auth
	}
	err = global.PantaDb.Create(&auth).Error
	return err, auth
}

// @title    CopyAuthority
// @description   复制一个角色
// @auth                     （2020/04/05  20:22）
// @param     copyInfo        response.SysAuthorityCopyResponse
// @return                    error
// @return    authority       model.SysAuthority

func CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (err error, authority model.SysAuthority) {
	var authorityBox model.SysAuthority
	notHas := errors.Is(global.PantaDb.Where("authority_id = ?", copyInfo.Authority.AuthorityId).Find(&authorityBox).
		Error, gorm.ErrRecordNotFound)
	if !notHas {
		return errors.New("存在相同角色id"), authority
	}
	copyInfo.Authority.Children = []model.SysAuthority{}
	err, menus := GetMenuAuthority(copyInfo.OldAuthorityId)
	var baseMenu []model.SysBaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = global.PantaDb.Create(&copyInfo.Authority).Error

	paths := GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = DeleteAuthority(&copyInfo.Authority)
	}
	return err, copyInfo.Authority
}

// @title    UpdateAuthority
// @description   更改一个角色
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
// @return    authority       model.SysAuthority

func UpdateAuthority(auth model.SysAuthority) (err error, authority model.SysAuthority) {
	err = global.PantaDb.Where("authority_id = ?", auth.AuthorityId).First(&model.SysAuthority{}).Updates(&auth).Error
	return err, auth
}

// @title    DeleteAuthority
// @description   删除角色
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
// 删除角色

func DeleteAuthority(auth *model.SysAuthority) (err error) {
	err = global.PantaDb.Where("authority_id = ?", auth.AuthorityId).Find(&model.SysUser{}).Error
	if err == nil {
		err = errors.New("此角色有用户正在使用禁止删除")
		return
	}
	err = global.PantaDb.Where("parent_id = ?", auth.AuthorityId).Find(&model.SysAuthority{}).Error
	if err == nil {
		err = errors.New("此角色存在子角色不允许删除")
		return
	}
	db := global.PantaDb.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth).Unscoped().Delete(auth)
	if len(auth.SysBaseMenus) > 0 {
		err = db.Association("SysBaseMenus").Delete(auth.SysBaseMenus)
	} else {
		err = db.Error
	}
	ClearCasbin(0, auth.AuthorityId)
	return err
}

// @title    GetInfoList
// @description   删除文件切片记录
// @auth                     （2020/04/05  20:22）
// @param     info            request.PaveInfo
// @return                    error
// 分页获取数据

func GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.PantaDb
	var authority []model.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = 0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = findChildrenAuthority(&authority[k])
		}
	}
	return err, authority, total
}

// @title    GetAuthorityInfo
// @description   获取所有角色信息
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
// @param     authority       model.SysAuthority

func GetAuthorityInfo(auth model.SysAuthority) (err error, sa model.SysAuthority) {
	err = global.PantaDb.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

// @title    SetDataAuthority
// @description   设置角色资源权限
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error

func SetDataAuthority(auth model.SysAuthority) error {
	var s model.SysAuthority
	global.PantaDb.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	return global.PantaDb.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)

}

// @title    SetMenuAuthority
// @description   菜单与角色绑定
// @auth                     （2020/04/05  20:22）
// @param     auth            *model.SysAuthority
// @return                    error

func SetMenuAuthority(auth *model.SysAuthority) error {
	var s model.SysAuthority
	global.PantaDb.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.PantaDb.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

// @title    findChildrenAuthority
// @description   查询子角色
// @auth                     （2020/04/05  20:22）
// @param     auth            *model.SysAuthority
// @return                    error

func findChildrenAuthority(authority *model.SysAuthority) (err error) {
	err = global.PantaDb.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
