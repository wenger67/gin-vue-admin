package dev

import (
	"panta/api/dev"
	"panta/middleware"
	"github.com/gin-gonic/gin"
)

func InitDataRouter(Router *gin.RouterGroup) {
	DataRouter := Router.Group("data").Use(middleware.OperationRecord())
	{
		DataRouter.POST("createRunningData", dev.CreateRunningData)
	}
}