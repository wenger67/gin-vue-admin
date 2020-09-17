package router

import (
	"panta/api/v1"
	"panta/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserLiftRouter(Router *gin.RouterGroup) {
	UserLiftRouter := Router.Group("userLift").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		UserLiftRouter.POST("createUserLift", v1.CreateUserLift)   // 新建UserLift
		UserLiftRouter.DELETE("deleteUserLift", v1.DeleteUserLift) // 删除UserLift
		UserLiftRouter.DELETE("deleteUserLiftByIds", v1.DeleteUserLiftByIds) // 批量删除UserLift
		UserLiftRouter.PUT("updateUserLift", v1.UpdateUserLift)    // 更新UserLift
		UserLiftRouter.GET("findUserLift", v1.FindUserLift)        // 根据ID获取UserLift
		UserLiftRouter.GET("getUserLiftList", v1.GetUserLiftList)  // 获取UserLift列表
	}
}
