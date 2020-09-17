package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gorm.io/gorm"
)

// @title    DeleteBaseMenu
// @description   删除基础路由
// @auth                     （2020/04/05  20:22）
// @param     id              float64
// @return    err             error

func DeleteBaseMenu(id float64) (err error) {
	err = global.PantaDb.Where("parent_id = ?", id).First(&model.SysBaseMenu{}).Error
	if err != nil {
		var menu model.SysBaseMenu
		db := global.PantaDb.Preload("SysAuthoritys").Where("id = ?", id).First(&menu).Delete(&menu)
		if len(menu.SysAuthoritys) > 0 {
			err = db.Association("SysAuthoritys").Delete(menu.SysAuthoritys)
		} else {
			err = db.Error
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

// @title    UpdateBaseMenu
// @description   更新路由
// @auth                     （2020/04/05  20:22）
// @param     menu            model.SysBaseMenu
// @return    err             errorgetMenu

func UpdateBaseMenu(menu model.SysBaseMenu) (err error) {
	var oldMenu model.SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort
	db := global.PantaDb.Where("id = ?", menu.ID).Find(&oldMenu)
	if oldMenu.Name != menu.Name {
		notSame := errors.Is(global.PantaDb.Where("id <> ? AND name = ?", menu.ID,
			menu.Name).First(&model.SysBaseMenu{}).Error, gorm.ErrRecordNotFound)
		if !notSame {
			global.PantaLog.Debug("存在相同name修改失败")
			return errors.New("存在相同name修改失败")
		}
	}
	err = db.Updates(upDateMap).Error
	global.PantaLog.Debug("菜单修改时候，关联菜单err:%v", err)
	return err
}

// @title    GetBaseMenuById
// @description   get current menus, 返回当前选中menu
// @auth                     （2020/04/05  20:22）
// @param     id              float64
// @return    err             error

func GetBaseMenuById(id float64) (err error, menu model.SysBaseMenu) {
	err = global.PantaDb.Where("id = ?", id).First(&menu).Error
	return
}
