package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	MessageRouter := Router.Group("message").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		MessageRouter.POST("createMessage", v1.CreateMessage)   // 新建Message
		MessageRouter.DELETE("deleteMessage", v1.DeleteMessage) // 删除Message
		MessageRouter.DELETE("deleteMessageByIds", v1.DeleteMessageByIds) // 批量删除Message
		MessageRouter.PUT("updateMessage", v1.UpdateMessage)    // 更新Message
		MessageRouter.GET("findMessage", v1.FindMessage)        // 根据ID获取Message
		MessageRouter.GET("getMessageList", v1.GetMessageList)  // 获取Message列表
	}
}
