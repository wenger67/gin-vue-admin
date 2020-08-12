package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitLiftChangeRouter(Router *gin.RouterGroup) {
	LiftChangeRouter := Router.Group("liftChange").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		LiftChangeRouter.POST("createLiftChange", v1.CreateLiftChange)   // 新建LiftChange
		LiftChangeRouter.DELETE("deleteLiftChange", v1.DeleteLiftChange) // 删除LiftChange
		LiftChangeRouter.DELETE("deleteLiftChangeByIds", v1.DeleteLiftChangeByIds) // 批量删除LiftChange
		LiftChangeRouter.PUT("updateLiftChange", v1.UpdateLiftChange)    // 更新LiftChange
		LiftChangeRouter.GET("findLiftChange", v1.FindLiftChange)        // 根据ID获取LiftChange
		LiftChangeRouter.GET("getLiftChangeList", v1.GetLiftChangeList)  // 获取LiftChange列表
	}
}
