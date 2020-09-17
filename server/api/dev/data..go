package dev

import (
	"fmt"
	"panta/global"
	"panta/global/response"
	"panta/model"
	"panta/service/dev"
	"github.com/gin-gonic/gin"
)

func CreateRunningData(c *gin.Context) {
	var adDeviceData model.AdDeviceData
	_ = c.ShouldBindJSON(&adDeviceData)
	global.PantaLog.Debug(adDeviceData);
	err := dev.CreateRunningData(adDeviceData)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
