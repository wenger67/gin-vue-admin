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

// @Tags HealthChange
// @Summary 创建HealthChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthChange true "创建HealthChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /healthChange/createHealthChange [post]
func CreateHealthChange(c *gin.Context) {
	var healthChange model.HealthChange
	_ = c.ShouldBindJSON(&healthChange)
	err := service.CreateHealthChange(healthChange)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags HealthChange
// @Summary 删除HealthChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthChange true "删除HealthChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /healthChange/deleteHealthChange [delete]
func DeleteHealthChange(c *gin.Context) {
	var healthChange model.HealthChange
	_ = c.ShouldBindJSON(&healthChange)
	err := service.DeleteHealthChange(healthChange)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags HealthChange
// @Summary 批量删除HealthChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HealthChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /healthChange/deleteHealthChangeByIds [delete]
func DeleteHealthChangeByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteHealthChangeByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags HealthChange
// @Summary 更新HealthChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthChange true "更新HealthChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /healthChange/updateHealthChange [put]
func UpdateHealthChange(c *gin.Context) {
	var healthChange model.HealthChange
	_ = c.ShouldBindJSON(&healthChange)
	err := service.UpdateHealthChange(&healthChange)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags HealthChange
// @Summary 用id查询HealthChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthChange true "用id查询HealthChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /healthChange/findHealthChange [get]
func FindHealthChange(c *gin.Context) {
	var healthChange model.HealthChange
	_ = c.ShouldBindQuery(&healthChange)
	err, rehealthChange := service.GetHealthChange(healthChange.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rehealthChange": rehealthChange}, c)
	}
}

// @Tags HealthChange
// @Summary 分页获取HealthChange列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.HealthChangeSearch true "分页获取HealthChange列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /healthChange/getHealthChangeList [get]
func GetHealthChangeList(c *gin.Context) {
	var pageInfo request.HealthChangeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetHealthChangeInfoList(pageInfo)
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
