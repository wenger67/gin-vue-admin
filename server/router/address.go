package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAddressRouter(Router *gin.RouterGroup) {
	AddressRouter := Router.Group("address").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		AddressRouter.POST("createAddress", v1.CreateAddress)   // 新建Address
		AddressRouter.DELETE("deleteAddress", v1.DeleteAddress) // 删除Address
		AddressRouter.DELETE("deleteAddressByIds", v1.DeleteAddressByIds) // 批量删除Address
		AddressRouter.PUT("updateAddress", v1.UpdateAddress)    // 更新Address
		AddressRouter.GET("findAddress", v1.FindAddress)        // 根据ID获取Address
		AddressRouter.GET("getAddressList", v1.GetAddressList)  // 获取Address列表
	}
}
