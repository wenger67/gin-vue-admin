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

// @Tags AdDeviceData
// @Summary 创建AdDeviceData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceData true "创建AdDeviceData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDeviceData/createAdDeviceData [post]
func CreateAdDeviceData(c *gin.Context) {
	var adDeviceData model.AdDeviceData
	_ = c.ShouldBindJSON(&adDeviceData)
	err := service.CreateAdDeviceData(adDeviceData)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AdDeviceData
// @Summary 删除AdDeviceData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceData true "删除AdDeviceData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDeviceData/deleteAdDeviceData [delete]
func DeleteAdDeviceData(c *gin.Context) {
	var adDeviceData model.AdDeviceData
	_ = c.ShouldBindJSON(&adDeviceData)
	err := service.DeleteAdDeviceData(adDeviceData)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdDeviceData
// @Summary 批量删除AdDeviceData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdDeviceData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDeviceData/deleteAdDeviceDataByIds [delete]
func DeleteAdDeviceDataByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteAdDeviceDataByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdDeviceData
// @Summary 更新AdDeviceData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceData true "更新AdDeviceData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adDeviceData/updateAdDeviceData [put]
func UpdateAdDeviceData(c *gin.Context) {
	var adDeviceData model.AdDeviceData
	_ = c.ShouldBindJSON(&adDeviceData)
	err := service.UpdateAdDeviceData(&adDeviceData)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AdDeviceData
// @Summary 用id查询AdDeviceData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceData true "用id查询AdDeviceData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adDeviceData/findAdDeviceData [get]
func FindAdDeviceData(c *gin.Context) {
	var adDeviceData model.AdDeviceData
	_ = c.ShouldBindQuery(&adDeviceData)
	err, readDeviceData := service.GetAdDeviceData(adDeviceData.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"readDeviceData": readDeviceData}, c)
	}
}

// @Tags AdDeviceData
// @Summary 分页获取AdDeviceData列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AdDeviceDataSearch true "分页获取AdDeviceData列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDeviceData/getAdDeviceDataList [get]
func GetAdDeviceDataList(c *gin.Context) {
	var pageInfo request.AdDeviceDataSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetAdDeviceDataInfoList(pageInfo)
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
