package core

import (
	"fmt"
	"panta/global"
	"panta/initialize"
	"panta/websocket"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.PantaConfig.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")
	Router.Static("/upload", "./resource/upload")

	// handle websocket request
	Router.GET("/api/ws/endpoint", websocket.ServeWS)

	//InstallPlugs(Router)
	// end 插件描述

	address := fmt.Sprintf(":%d", global.PantaConfig.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.PantaLog.Debug("server run success on ", address)

	fmt.Printf(`欢迎使用 panta
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.PantaLog.Error(s.ListenAndServe())
}
