package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRegionRouter(Router *gin.RouterGroup) {
	RegionRouter := Router.Group("region").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		RegionRouter.POST("createRegion", v1.CreateRegion)   // 新建Region
		RegionRouter.DELETE("deleteRegion", v1.DeleteRegion) // 删除Region
		RegionRouter.DELETE("deleteRegionByIds", v1.DeleteRegionByIds) // 批量删除Region
		RegionRouter.PUT("updateRegion", v1.UpdateRegion)    // 更新Region
		RegionRouter.GET("findRegion", v1.FindRegion)        // 根据ID获取Region
		RegionRouter.GET("getRegionList", v1.GetRegionList)  // 获取Region列表
	}
}
