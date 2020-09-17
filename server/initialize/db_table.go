package initialize

import (
	"panta/global"
	"panta/model"
	"os"
)

// 注册数据库表专用
func DBTables() {
	db := global.PantaDb
	err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		AutoMigrate(model.SysUser{},
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
		model.HealthSystem{},
		model.HealthChange{},
		model.Message{},
	)
	if err != nil {
		global.PantaLog.Error("register table failed", err)
		os.Exit(0)
	}
	global.PantaLog.Debug("register table success")
}
