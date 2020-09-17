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

// @Tags AdDeviceConfig
// @Summary 创建AdDeviceConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceConfig true "创建AdDeviceConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDeviceConfig/createAdDeviceConfig [post]
func CreateAdDeviceConfig(c *gin.Context) {
	var adDeviceConfig model.AdDeviceConfig
	_ = c.ShouldBindJSON(&adDeviceConfig)
	err := service.CreateAdDeviceConfig(adDeviceConfig)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AdDeviceConfig
// @Summary 删除AdDeviceConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceConfig true "删除AdDeviceConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDeviceConfig/deleteAdDeviceConfig [delete]
func DeleteAdDeviceConfig(c *gin.Context) {
	var adDeviceConfig model.AdDeviceConfig
	_ = c.ShouldBindJSON(&adDeviceConfig)
	err := service.DeleteAdDeviceConfig(adDeviceConfig)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdDeviceConfig
// @Summary 批量删除AdDeviceConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdDeviceConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDeviceConfig/deleteAdDeviceConfigByIds [delete]
func DeleteAdDeviceConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteAdDeviceConfigByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdDeviceConfig
// @Summary 更新AdDeviceConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceConfig true "更新AdDeviceConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adDeviceConfig/updateAdDeviceConfig [put]
func UpdateAdDeviceConfig(c *gin.Context) {
	var adDeviceConfig model.AdDeviceConfig
	_ = c.ShouldBindJSON(&adDeviceConfig)
	err := service.UpdateAdDeviceConfig(&adDeviceConfig)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AdDeviceConfig
// @Summary 用id查询AdDeviceConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceConfig true "用id查询AdDeviceConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adDeviceConfig/findAdDeviceConfig [get]
func FindAdDeviceConfig(c *gin.Context) {
	var adDeviceConfig model.AdDeviceConfig
	_ = c.ShouldBindQuery(&adDeviceConfig)
	err, readDeviceConfig := service.GetAdDeviceConfig(adDeviceConfig.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"readDeviceConfig": readDeviceConfig}, c)
	}
}

// @Tags AdDeviceConfig
// @Summary 分页获取AdDeviceConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AdDeviceConfigSearch true "分页获取AdDeviceConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDeviceConfig/getAdDeviceConfigList [get]
func GetAdDeviceConfigList(c *gin.Context) {
	var pageInfo request.AdDeviceConfigSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetAdDeviceConfigInfoList(pageInfo)
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
