package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitLiftTroubleRouter(Router *gin.RouterGroup) {
	LiftTroubleRouter := Router.Group("liftTrouble").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		LiftTroubleRouter.POST("createLiftTrouble", v1.CreateLiftTrouble)   // 新建LiftTrouble
		LiftTroubleRouter.DELETE("deleteLiftTrouble", v1.DeleteLiftTrouble) // 删除LiftTrouble
		LiftTroubleRouter.DELETE("deleteLiftTroubleByIds", v1.DeleteLiftTroubleByIds) // 批量删除LiftTrouble
		LiftTroubleRouter.PUT("updateLiftTrouble", v1.UpdateLiftTrouble)    // 更新LiftTrouble
		LiftTroubleRouter.GET("findLiftTrouble", v1.FindLiftTrouble)        // 根据ID获取LiftTrouble
		LiftTroubleRouter.GET("getLiftTroubleList", v1.GetLiftTroubleList)  // 获取LiftTrouble列表
	}
}
