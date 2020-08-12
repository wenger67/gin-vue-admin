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

// @Tags LiftTrouble
// @Summary 创建LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftTrouble true "创建LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftTrouble/createLiftTrouble [post]
func CreateLiftTrouble(c *gin.Context) {
	var liftTrouble model.LiftTrouble
	_ = c.ShouldBindJSON(&liftTrouble)
	err := service.CreateLiftTrouble(liftTrouble)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags LiftTrouble
// @Summary 删除LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftTrouble true "删除LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftTrouble/deleteLiftTrouble [delete]
func DeleteLiftTrouble(c *gin.Context) {
	var liftTrouble model.LiftTrouble
	_ = c.ShouldBindJSON(&liftTrouble)
	err := service.DeleteLiftTrouble(liftTrouble)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags LiftTrouble
// @Summary 批量删除LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftTrouble/deleteLiftTroubleByIds [delete]
func DeleteLiftTroubleByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteLiftTroubleByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags LiftTrouble
// @Summary 更新LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftTrouble true "更新LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /liftTrouble/updateLiftTrouble [put]
func UpdateLiftTrouble(c *gin.Context) {
	var liftTrouble model.LiftTrouble
	_ = c.ShouldBindJSON(&liftTrouble)
	err := service.UpdateLiftTrouble(&liftTrouble)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags LiftTrouble
// @Summary 用id查询LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftTrouble true "用id查询LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /liftTrouble/findLiftTrouble [get]
func FindLiftTrouble(c *gin.Context) {
	var liftTrouble model.LiftTrouble
	_ = c.ShouldBindQuery(&liftTrouble)
	err, reliftTrouble := service.GetLiftTrouble(liftTrouble.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reliftTrouble": reliftTrouble}, c)
	}
}

// @Tags LiftTrouble
// @Summary 分页获取LiftTrouble列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LiftTroubleSearch true "分页获取LiftTrouble列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftTrouble/getLiftTroubleList [get]
func GetLiftTroubleList(c *gin.Context) {
	var pageInfo request.LiftTroubleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetLiftTroubleInfoList(pageInfo)
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
