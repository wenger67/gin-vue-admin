package dev

import (
	"gin-vue-admin/api/dev"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDataRouter(Router *gin.RouterGroup) {
	DataRouter := Router.Group("data").Use(middleware.OperationRecord())
	{
		DataRouter.POST("createRunningData", dev.CreateRunningData)
	}
}