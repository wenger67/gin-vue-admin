package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitHealthChangeRouter(Router *gin.RouterGroup) {
	HealthChangeRouter := Router.Group("healthChange").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		HealthChangeRouter.POST("createHealthChange", v1.CreateHealthChange)   // 新建HealthChange
		HealthChangeRouter.DELETE("deleteHealthChange", v1.DeleteHealthChange) // 删除HealthChange
		HealthChangeRouter.DELETE("deleteHealthChangeByIds", v1.DeleteHealthChangeByIds) // 批量删除HealthChange
		HealthChangeRouter.PUT("updateHealthChange", v1.UpdateHealthChange)    // 更新HealthChange
		HealthChangeRouter.GET("findHealthChange", v1.FindHealthChange)        // 根据ID获取HealthChange
		HealthChangeRouter.GET("getHealthChangeList", v1.GetHealthChangeList)  // 获取HealthChange列表
	}
}
