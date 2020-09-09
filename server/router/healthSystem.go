package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitHealthSystemRouter(Router *gin.RouterGroup) {
	HealthSystemRouter := Router.Group("healthSystem").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		HealthSystemRouter.POST("createHealthSystem", v1.CreateHealthSystem)   // 新建HealthSystem
		HealthSystemRouter.DELETE("deleteHealthSystem", v1.DeleteHealthSystem) // 删除HealthSystem
		HealthSystemRouter.DELETE("deleteHealthSystemByIds", v1.DeleteHealthSystemByIds) // 批量删除HealthSystem
		HealthSystemRouter.PUT("updateHealthSystem", v1.UpdateHealthSystem)    // 更新HealthSystem
		HealthSystemRouter.GET("findHealthSystem", v1.FindHealthSystem)        // 根据ID获取HealthSystem
		HealthSystemRouter.GET("getHealthSystemList", v1.GetHealthSystemList)  // 获取HealthSystem列表
	}
}
