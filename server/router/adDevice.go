package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdDeviceRouter(Router *gin.RouterGroup) {
	AdDeviceRouter := Router.Group("adDevice").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		AdDeviceRouter.POST("createAdDevice", v1.CreateAdDevice)   // 新建AdDevice
		AdDeviceRouter.DELETE("deleteAdDevice", v1.DeleteAdDevice) // 删除AdDevice
		AdDeviceRouter.DELETE("deleteAdDeviceByIds", v1.DeleteAdDeviceByIds) // 批量删除AdDevice
		AdDeviceRouter.PUT("updateAdDevice", v1.UpdateAdDevice)    // 更新AdDevice
		AdDeviceRouter.GET("findAdDevice", v1.FindAdDevice)        // 根据ID获取AdDevice
		AdDeviceRouter.GET("getAdDeviceList", v1.GetAdDeviceList)  // 获取AdDevice列表
	}
}
