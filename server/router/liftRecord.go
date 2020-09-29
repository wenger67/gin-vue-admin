package router

import (
	"panta/api/v1"
	"panta/middleware"
	"github.com/gin-gonic/gin"
)

func InitLiftRecordRouter(Router *gin.RouterGroup) {
	LiftRecordRouter := Router.Group("liftRecord").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		LiftRecordRouter.POST("createLiftRecord", v1.CreateLiftRecord)   // 新建LiftRecord
		LiftRecordRouter.PUT("updateLiftRecord", v1.UpdateLiftRecord)    // 更新LiftRecord

		LiftRecordRouter.DELETE("deleteLiftRecord", v1.DeleteLiftRecord) // 删除LiftRecord
		LiftRecordRouter.DELETE("deleteLiftRecordByIds", v1.DeleteLiftRecordByIds) // 批量删除LiftRecord
		LiftRecordRouter.GET("findLiftRecord", v1.FindLiftRecord)        // 根据ID获取LiftRecord
		LiftRecordRouter.GET("getLiftRecordList", v1.GetLiftRecordList)  // 获取LiftRecord列表
	}
}
