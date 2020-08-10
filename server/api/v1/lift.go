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

// @Tags Lift
// @Summary 创建Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lift true "创建Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lift/createLift [post]
func CreateLift(c *gin.Context) {
	var lift model.Lift
	_ = c.ShouldBindJSON(&lift)
	err := service.CreateLift(lift)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Lift
// @Summary 删除Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lift true "删除Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lift/deleteLift [delete]
func DeleteLift(c *gin.Context) {
	var lift model.Lift
	_ = c.ShouldBindJSON(&lift)
	err := service.DeleteLift(lift)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Lift
// @Summary 批量删除Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lift/deleteLiftByIds [delete]
func DeleteLiftByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteLiftByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Lift
// @Summary 更新Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lift true "更新Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lift/updateLift [put]
func UpdateLift(c *gin.Context) {
	var lift model.Lift
	_ = c.ShouldBindJSON(&lift)
	err := service.UpdateLift(&lift)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Lift
// @Summary 用id查询Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lift true "用id查询Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lift/findLift [get]
func FindLift(c *gin.Context) {
	var lift model.Lift
	_ = c.ShouldBindQuery(&lift)
	err, relift := service.GetLift(lift.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"relift": relift}, c)
	}
}

// @Tags Lift
// @Summary 分页获取Lift列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LiftSearch true "分页获取Lift列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lift/getLiftList [get]
func GetLiftList(c *gin.Context) {
	var pageInfo request.LiftSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetLiftInfoList(pageInfo)
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
