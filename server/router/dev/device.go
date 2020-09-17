package dev

import (
	"panta/api/dev"
	"panta/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeviceRouter(Router *gin.RouterGroup) {
	DeviceRouter := Router.Group("device").Use(middleware.OperationRecord())
	{
		DeviceRouter.GET("findDevice", dev.FindDevice)
		DeviceRouter.POST("createEvent", dev.CreateEvent);
	}
}
