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

// @Tags AdDeviceEvent
// @Summary 创建AdDeviceEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceEvent true "创建AdDeviceEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDeviceEvent/createAdDeviceEvent [post]
func CreateAdDeviceEvent(c *gin.Context) {
	var adDeviceEvent model.AdDeviceEvent
	_ = c.ShouldBindJSON(&adDeviceEvent)
	err := service.CreateAdDeviceEvent(adDeviceEvent)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AdDeviceEvent
// @Summary 删除AdDeviceEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceEvent true "删除AdDeviceEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDeviceEvent/deleteAdDeviceEvent [delete]
func DeleteAdDeviceEvent(c *gin.Context) {
	var adDeviceEvent model.AdDeviceEvent
	_ = c.ShouldBindJSON(&adDeviceEvent)
	err := service.DeleteAdDeviceEvent(adDeviceEvent)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdDeviceEvent
// @Summary 批量删除AdDeviceEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdDeviceEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDeviceEvent/deleteAdDeviceEventByIds [delete]
func DeleteAdDeviceEventByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteAdDeviceEventByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdDeviceEvent
// @Summary 更新AdDeviceEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceEvent true "更新AdDeviceEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adDeviceEvent/updateAdDeviceEvent [put]
func UpdateAdDeviceEvent(c *gin.Context) {
	var adDeviceEvent model.AdDeviceEvent
	_ = c.ShouldBindJSON(&adDeviceEvent)
	err := service.UpdateAdDeviceEvent(&adDeviceEvent)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AdDeviceEvent
// @Summary 用id查询AdDeviceEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceEvent true "用id查询AdDeviceEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adDeviceEvent/findAdDeviceEvent [get]
func FindAdDeviceEvent(c *gin.Context) {
	var adDeviceEvent model.AdDeviceEvent
	_ = c.ShouldBindQuery(&adDeviceEvent)
	err, readDeviceEvent := service.GetAdDeviceEvent(adDeviceEvent.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"readDeviceEvent": readDeviceEvent}, c)
	}
}

// @Tags AdDeviceEvent
// @Summary 分页获取AdDeviceEvent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AdDeviceEventSearch true "分页获取AdDeviceEvent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDeviceEvent/getAdDeviceEventList [get]
func GetAdDeviceEventList(c *gin.Context) {
	var pageInfo request.AdDeviceEventSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetAdDeviceEventInfoList(pageInfo)
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
