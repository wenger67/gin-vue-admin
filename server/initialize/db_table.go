package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

// 注册数据库表专用
func DBTables() {
	db := global.GVA_DB
	db.AutoMigrate(model.SysUser{},
		model.SysAuthority{},
		model.SysApi{},
		model.SysBaseMenu{},
		model.JwtBlacklist{},
		model.SysWorkflow{},
		model.SysWorkflowStepInfo{},
		model.SysDictionary{},
		model.SysDictionaryDetail{},
		model.ExaFileUploadAndDownload{},
		model.ExaFile{},
		model.ExaFileChunk{},
		model.ExaCustomer{},
		model.SysOperationRecord{},
		model.CategorySubject{},
		model.Category{},
		model.Address{},
		model.Lift{},
		model.LiftChange{},
		model.LiftModel{},
		model.LiftRecord{},
		model.LiftTrouble{},
		model.UserLift{},
		model.Company{},
		model.Region{},
		model.AdDevice{},
		model.AdDeviceConfig{},
		model.AdDeviceData{},
		model.AdDeviceEvent{},
	)
	global.GVA_LOG.Debug("register table success")
}
