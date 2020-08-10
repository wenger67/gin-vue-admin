package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitLiftRouter(Router *gin.RouterGroup) {
	LiftRouter := Router.Group("lift").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		LiftRouter.POST("createLift", v1.CreateLift)   // 新建Lift
		LiftRouter.DELETE("deleteLift", v1.DeleteLift) // 删除Lift
		LiftRouter.DELETE("deleteLiftByIds", v1.DeleteLiftByIds) // 批量删除Lift
		LiftRouter.PUT("updateLift", v1.UpdateLift)    // 更新Lift
		LiftRouter.GET("findLift", v1.FindLift)        // 根据ID获取Lift
		LiftRouter.GET("getLiftList", v1.GetLiftList)  // 获取Lift列表
	}
}
