package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdDeviceDataRouter(Router *gin.RouterGroup) {
	AdDeviceDataRouter := Router.Group("adDeviceData").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		AdDeviceDataRouter.POST("createAdDeviceData", v1.CreateAdDeviceData)   // 新建AdDeviceData
		AdDeviceDataRouter.DELETE("deleteAdDeviceData", v1.DeleteAdDeviceData) // 删除AdDeviceData
		AdDeviceDataRouter.DELETE("deleteAdDeviceDataByIds", v1.DeleteAdDeviceDataByIds) // 批量删除AdDeviceData
		AdDeviceDataRouter.PUT("updateAdDeviceData", v1.UpdateAdDeviceData)    // 更新AdDeviceData
		AdDeviceDataRouter.GET("findAdDeviceData", v1.FindAdDeviceData)        // 根据ID获取AdDeviceData
		AdDeviceDataRouter.GET("getAdDeviceDataList", v1.GetAdDeviceDataList)  // 获取AdDeviceData列表
	}
}
