package main

import (
	"panta/core"
	"panta/global"
	"panta/initialize"
	//"runtime"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	switch global.PantaConfig.System.DbType {
	case "mysql":
		initialize.Mysql()
	// case "sqlite":
	//	initialize.Sqlite()  // sqlite需要gcc支持 windows用户需要自行安装gcc 如需使用打开注释即可
	default:
		initialize.Mysql()
	}
	initialize.DBTables()
	if global.PantaConfig.System.NeedInitData {
		initialize.InitData()
	}
	// 程序结束前关闭数据库链接
	db,_ := global.PantaDb.DB()
	defer db.Close()

	core.RunWindowsServer()
}
