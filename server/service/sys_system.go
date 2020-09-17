package service

import (
	"panta/config"
	"panta/global"
	"panta/model"
	"panta/utils"
)

// @title    GetSystemConfig
// @description   读取配置文件
// @auth                     （2020/04/05  20:22）
// @return    err             error
// @return    conf            Server

func GetSystemConfig() (err error, conf config.Server) {
	return nil, global.PantaConfig
}

// @title    SetSystemConfig
// @description   set system config, 设置配置文件
// @auth                    （2020/04/05  20:22）
// @param     system         model.System
// @return    err            error

func SetSystemConfig(system model.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.PantaVp.Set(k, v)
	}
	err = global.PantaVp.WriteConfig()
	return err
}
