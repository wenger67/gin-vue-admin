package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCategorySubjectRouter(Router *gin.RouterGroup) {
	CategorySubjectRouter := Router.Group("subject").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		CategorySubjectRouter.POST("createSubject", v1.CreateSubject)
		CategorySubjectRouter.POST("deleteSubject", v1.DeleteSubject)
		CategorySubjectRouter.POST("getSubjectList", v1.GetSubjectList)
		CategorySubjectRouter.POST("getSubjectById", v1.GetSubjectById)
		CategorySubjectRouter.POST("updateSubject", v1.UpdateSubject)
		CategorySubjectRouter.POST("getAllSubjects", v1.GetAllSubjects)
	}
}