package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdDeviceEventRouter(Router *gin.RouterGroup) {
	AdDeviceEventRouter := Router.Group("adDeviceEvent").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		AdDeviceEventRouter.POST("createAdDeviceEvent", v1.CreateAdDeviceEvent)   // 新建AdDeviceEvent
		AdDeviceEventRouter.DELETE("deleteAdDeviceEvent", v1.DeleteAdDeviceEvent) // 删除AdDeviceEvent
		AdDeviceEventRouter.DELETE("deleteAdDeviceEventByIds", v1.DeleteAdDeviceEventByIds) // 批量删除AdDeviceEvent
		AdDeviceEventRouter.PUT("updateAdDeviceEvent", v1.UpdateAdDeviceEvent)    // 更新AdDeviceEvent
		AdDeviceEventRouter.GET("findAdDeviceEvent", v1.FindAdDeviceEvent)        // 根据ID获取AdDeviceEvent
		AdDeviceEventRouter.GET("getAdDeviceEventList", v1.GetAdDeviceEventList)  // 获取AdDeviceEvent列表
	}
}
