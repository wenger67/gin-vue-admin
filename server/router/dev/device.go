package dev

import (
	"gin-vue-admin/api/dev"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeviceRouter(Router *gin.RouterGroup) {
	DeviceRouter := Router.Group("device").Use(middleware.OperationRecord())
	{
		DeviceRouter.GET("findDevice", dev.FindDevice)
	}
}
