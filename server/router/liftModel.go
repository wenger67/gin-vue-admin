package router

import (
	"panta/api/v1"
	"panta/middleware"
	"github.com/gin-gonic/gin"
)

func InitLiftModelRouter(Router *gin.RouterGroup) {
	LiftModelRouter := Router.Group("liftModel").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		LiftModelRouter.POST("createLiftModel", v1.CreateLiftModel)   // 新建LiftModel
		LiftModelRouter.DELETE("deleteLiftModel", v1.DeleteLiftModel) // 删除LiftModel
		LiftModelRouter.DELETE("deleteLiftModelByIds", v1.DeleteLiftModelByIds) // 批量删除LiftModel
		LiftModelRouter.PUT("updateLiftModel", v1.UpdateLiftModel)    // 更新LiftModel
		LiftModelRouter.GET("findLiftModel", v1.FindLiftModel)        // 根据ID获取LiftModel
		LiftModelRouter.GET("getLiftModelList", v1.GetLiftModelList)  // 获取LiftModel列表
	}
}
