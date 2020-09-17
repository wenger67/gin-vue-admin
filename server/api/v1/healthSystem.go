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

// @Tags HealthSystem
// @Summary 创建HealthSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthSystem true "创建HealthSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /healthSystem/createHealthSystem [post]
func CreateHealthSystem(c *gin.Context) {
	var healthSystem model.HealthSystem
	_ = c.ShouldBindJSON(&healthSystem)
	err := service.CreateHealthSystem(healthSystem)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags HealthSystem
// @Summary 删除HealthSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthSystem true "删除HealthSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /healthSystem/deleteHealthSystem [delete]
func DeleteHealthSystem(c *gin.Context) {
	var healthSystem model.HealthSystem
	_ = c.ShouldBindJSON(&healthSystem)
	err := service.DeleteHealthSystem(healthSystem)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags HealthSystem
// @Summary 批量删除HealthSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HealthSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /healthSystem/deleteHealthSystemByIds [delete]
func DeleteHealthSystemByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteHealthSystemByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags HealthSystem
// @Summary 更新HealthSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthSystem true "更新HealthSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /healthSystem/updateHealthSystem [put]
func UpdateHealthSystem(c *gin.Context) {
	var healthSystem model.HealthSystem
	_ = c.ShouldBindJSON(&healthSystem)
	err := service.UpdateHealthSystem(&healthSystem)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags HealthSystem
// @Summary 用id查询HealthSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthSystem true "用id查询HealthSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /healthSystem/findHealthSystem [get]
func FindHealthSystem(c *gin.Context) {
	var healthSystem model.HealthSystem
	_ = c.ShouldBindQuery(&healthSystem)
	err, rehealthSystem := service.GetHealthSystem(healthSystem.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rehealthSystem": rehealthSystem}, c)
	}
}

// @Tags HealthSystem
// @Summary 分页获取HealthSystem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.HealthSystemSearch true "分页获取HealthSystem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /healthSystem/getHealthSystemList [get]
func GetHealthSystemList(c *gin.Context) {
	var pageInfo request.HealthSystemSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetHealthSystemInfoList(pageInfo)
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
