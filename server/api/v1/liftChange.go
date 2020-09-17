package v1

import (
	"fmt"
	"panta/global/response"
	"panta/model"
	"panta/model/request"
	resp "panta/model/response"
	"panta/service"
	"github.com/gin-gonic/gin"
)

// @Tags LiftChange
// @Summary 创建LiftChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftChange true "创建LiftChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftChange/createLiftChange [post]
func CreateLiftChange(c *gin.Context) {
	var liftChange model.LiftChange
	_ = c.ShouldBindJSON(&liftChange)
	err := service.CreateLiftChange(liftChange)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags LiftChange
// @Summary 删除LiftChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftChange true "删除LiftChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftChange/deleteLiftChange [delete]
func DeleteLiftChange(c *gin.Context) {
	var liftChange model.LiftChange
	_ = c.ShouldBindJSON(&liftChange)
	err := service.DeleteLiftChange(liftChange)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags LiftChange
// @Summary 批量删除LiftChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LiftChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftChange/deleteLiftChangeByIds [delete]
func DeleteLiftChangeByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteLiftChangeByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags LiftChange
// @Summary 更新LiftChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftChange true "更新LiftChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /liftChange/updateLiftChange [put]
func UpdateLiftChange(c *gin.Context) {
	var liftChange model.LiftChange
	_ = c.ShouldBindJSON(&liftChange)
	err := service.UpdateLiftChange(&liftChange)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags LiftChange
// @Summary 用id查询LiftChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftChange true "用id查询LiftChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /liftChange/findLiftChange [get]
func FindLiftChange(c *gin.Context) {
	var liftChange model.LiftChange
	_ = c.ShouldBindQuery(&liftChange)
	err, reliftChange := service.GetLiftChange(liftChange.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reliftChange": reliftChange}, c)
	}
}

// @Tags LiftChange
// @Summary 分页获取LiftChange列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LiftChangeSearch true "分页获取LiftChange列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftChange/getLiftChangeList [get]
func GetLiftChangeList(c *gin.Context) {
	var pageInfo request.LiftChangeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetLiftChangeInfoList(pageInfo)
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
