package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCompanyRouter(Router *gin.RouterGroup) {
	CompanyRouter := Router.Group("company").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		CompanyRouter.POST("createCompany", v1.CreateCompany)   // 新建Company
		CompanyRouter.DELETE("deleteCompany", v1.DeleteCompany) // 删除Company
		CompanyRouter.DELETE("deleteCompanyByIds", v1.DeleteCompanyByIds) // 批量删除Company
		CompanyRouter.PUT("updateCompany", v1.UpdateCompany)    // 更新Company
		CompanyRouter.GET("findCompany", v1.FindCompany)        // 根据ID获取Company
		CompanyRouter.GET("getCompanyList", v1.GetCompanyList)  // 获取Company列表
	}
}
