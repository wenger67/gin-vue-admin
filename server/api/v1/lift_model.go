package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags LiftModel
// @Summary 创建LiftModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftModel true "创建LiftModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftModel/createLiftModel [post]
func CreateLiftModel(c *gin.Context) {
	var liftModel model.LiftModel
	_ = c.ShouldBindJSON(&liftModel)
	err := service.CreateLiftModel(liftModel)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags LiftModel
// @Summary 删除LiftModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftModel true "删除LiftModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftModel/deleteLiftModel [delete]
func DeleteLiftModel(c *gin.Context) {
	var liftModel model.LiftModel
	_ = c.ShouldBindJSON(&liftModel)
	err := service.DeleteLiftModel(liftModel)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags LiftModel
// @Summary 批量删除LiftModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LiftModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftModel/deleteLiftModelByIds [delete]
func DeleteLiftModelByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteLiftModelByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags LiftModel
// @Summary 更新LiftModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftModel true "更新LiftModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /liftModel/updateLiftModel [put]
func UpdateLiftModel(c *gin.Context) {
	var liftModel model.LiftModel
	_ = c.ShouldBindJSON(&liftModel)
	err := service.UpdateLiftModel(&liftModel)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags LiftModel
// @Summary 用id查询LiftModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftModel true "用id查询LiftModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /liftModel/findLiftModel [get]
func FindLiftModel(c *gin.Context) {
	var liftModel model.LiftModel
	_ = c.ShouldBindQuery(&liftModel)
	err, reliftModel := service.GetLiftModel(liftModel.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reliftModel": reliftModel}, c)
	}
}

// @Tags LiftModel
// @Summary 分页获取LiftModel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LiftModelSearch true "分页获取LiftModel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftModel/getLiftModelList [get]
func GetLiftModelList(c *gin.Context) {
	var pageInfo request.LiftModelSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetLiftModelInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
