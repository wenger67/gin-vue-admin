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

// @Tags AdDevice
// @Summary 创建AdDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDevice true "创建AdDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDevice/createAdDevice [post]
func CreateAdDevice(c *gin.Context) {
	var dCreate model.AdDevice
	_ = c.ShouldBindJSON(&dCreate)
	// TODO verify params
	err := service.CreateAdDevice(dCreate)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AdDevice
// @Summary 删除AdDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDevice true "删除AdDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDevice/deleteAdDevice [delete]
func DeleteAdDevice(c *gin.Context) {
	var adDevice model.AdDevice
	_ = c.ShouldBindJSON(&adDevice)
	err := service.DeleteAdDevice(adDevice)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdDevice
// @Summary 批量删除AdDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDevice/deleteAdDeviceByIds [delete]
func DeleteAdDeviceByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteAdDeviceByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdDevice
// @Summary 更新AdDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDevice true "更新AdDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adDevice/updateAdDevice [put]
func UpdateAdDevice(c *gin.Context) {
	var adDevice model.AdDevice
	_ = c.ShouldBindJSON(&adDevice)
	err := service.UpdateAdDevice(&adDevice)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AdDevice
// @Summary 用id查询AdDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDevice true "用id查询AdDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adDevice/findAdDevice [get]
func FindAdDevice(c *gin.Context) {
	var adDevice model.AdDevice
	_ = c.ShouldBindQuery(&adDevice)
	err, readDevice := service.GetAdDevice(adDevice.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"readDevice": readDevice}, c)
	}
}

// @Tags AdDevice
// @Summary 分页获取AdDevice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AdDeviceSearch true "分页获取AdDevice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDevice/getAdDeviceList [get]
func GetAdDeviceList(c *gin.Context) {
	var pageInfo request.AdDeviceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetAdDeviceInfoList(pageInfo)
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
