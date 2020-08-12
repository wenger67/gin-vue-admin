package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdDeviceConfigRouter(Router *gin.RouterGroup) {
	AdDeviceConfigRouter := Router.Group("adDeviceConfig").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		AdDeviceConfigRouter.POST("createAdDeviceConfig", v1.CreateAdDeviceConfig)   // 新建AdDeviceConfig
		AdDeviceConfigRouter.DELETE("deleteAdDeviceConfig", v1.DeleteAdDeviceConfig) // 删除AdDeviceConfig
		AdDeviceConfigRouter.DELETE("deleteAdDeviceConfigByIds", v1.DeleteAdDeviceConfigByIds) // 批量删除AdDeviceConfig
		AdDeviceConfigRouter.PUT("updateAdDeviceConfig", v1.UpdateAdDeviceConfig)    // 更新AdDeviceConfig
		AdDeviceConfigRouter.GET("findAdDeviceConfig", v1.FindAdDeviceConfig)        // 根据ID获取AdDeviceConfig
		AdDeviceConfigRouter.GET("getAdDeviceConfigList", v1.GetAdDeviceConfigList)  // 获取AdDeviceConfig列表
	}
}
