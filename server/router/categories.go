package router

import (
	"panta/api/v1"
	"panta/middleware"
	"github.com/gin-gonic/gin"
)

func InitCategoriesRouter(Router *gin.RouterGroup) {
	CategoriesRouter := Router.Group("categories").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		CategoriesRouter.POST("createCategories", v1.CreateCategories)   // 新建Categories
		CategoriesRouter.DELETE("deleteCategories", v1.DeleteCategories) // 删除Categories
		CategoriesRouter.DELETE("deleteCategoriesByIds", v1.DeleteCategoriesByIds) // 批量删除Categories
		CategoriesRouter.PUT("updateCategories", v1.UpdateCategories)    // 更新Categories
		CategoriesRouter.GET("findCategories", v1.FindCategories)        // 根据ID获取Categories
		CategoriesRouter.GET("getCategoriesList", v1.GetCategoriesList)  // 获取Categories列表
	}
}
