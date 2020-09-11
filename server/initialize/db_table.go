package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"os"
)

// 注册数据库表专用
func DBTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(model.SysUser{},
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
		model.Device{},
		model.AdDeviceConfig{},
		model.AdDeviceData{},
		model.AdDeviceEvent{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", err)
		os.Exit(0)
	}
	global.GVA_LOG.Debug("register table success")
}
