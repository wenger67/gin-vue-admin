package initialize

import (
	_ "panta/docs"
	"panta/global"
	"panta/middleware"
	"panta/router"
	"panta/router/dev"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.PantaLog.Debug("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	global.PantaLog.Debug("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.PantaLog.Debug("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	router.InitUserRouter(ApiGroup)                  // 注册用户路由
	router.InitBaseRouter(ApiGroup)                  // 注册基础功能路由 不做鉴权
	router.InitMenuRouter(ApiGroup)                  // 注册menu路由
	router.InitAuthorityRouter(ApiGroup)             // 注册角色路由
	router.InitApiRouter(ApiGroup)                   // 注册功能api路由
	router.InitFileUploadAndDownloadRouter(ApiGroup) // 文件上传下载功能路由
	router.InitWorkflowRouter(ApiGroup)              // 工作流相关路由
	router.InitCasbinRouter(ApiGroup)                // 权限相关路由
	router.InitJwtRouter(ApiGroup)                   // jwt相关路由
	router.InitSystemRouter(ApiGroup)                // system相关路由
	router.InitCustomerRouter(ApiGroup)              // 客户路由
	router.InitAutoCodeRouter(ApiGroup)              // 创建自动化代码
	router.InitSysDictionaryDetailRouter(ApiGroup)   // 字典详情管理
	router.InitSysDictionaryRouter(ApiGroup)         // 字典管理
	router.InitSysOperationRecordRouter(ApiGroup)    // 操作记录
	router.InitCategorySubjectRouter(ApiGroup)
	router.InitCategoriesRouter(ApiGroup)
	router.InitCompanyRouter(ApiGroup)
	router.InitRegionRouter(ApiGroup)
	router.InitAddressRouter(ApiGroup)
	router.InitLiftRouter(ApiGroup)
	router.InitLiftModelRouter(ApiGroup)
	router.InitLiftChangeRouter(ApiGroup)
	router.InitLiftRecordRouter(ApiGroup)
	router.InitLiftTroubleRouter(ApiGroup)
	router.InitUserLiftRouter(ApiGroup)
	router.InitAdDeviceConfigRouter(ApiGroup)
	router.InitAdDeviceDataRouter(ApiGroup)
	router.InitAdDeviceEventRouter(ApiGroup)
	router.InitAdDeviceRouter(ApiGroup)
	router.InitMessageRouter(ApiGroup)
	router.InitHealthChangeRouter(ApiGroup)
	router.InitHealthSystemRouter(ApiGroup)

	DevGroup := Router.Group("dev")
	dev.InitDeviceRouter(DevGroup)
	dev.InitDataRouter(DevGroup)

	global.PantaLog.Info("router register success")
	return Router
}
