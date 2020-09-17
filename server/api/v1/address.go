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

// @Tags Address
// @Summary 创建Address
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Address true "创建Address"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /address/createAddress [post]
func CreateAddress(c *gin.Context) {
	var address model.Address
	_ = c.ShouldBindJSON(&address)
	err := service.CreateAddress(address)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Address
// @Summary 删除Address
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Address true "删除Address"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /address/deleteAddress [delete]
func DeleteAddress(c *gin.Context) {
	var address model.Address
	_ = c.ShouldBindJSON(&address)
	err := service.DeleteAddress(address)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Address
// @Summary 批量删除Address
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Address"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /address/deleteAddressByIds [delete]
func DeleteAddressByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteAddressByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Address
// @Summary 更新Address
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Address true "更新Address"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /address/updateAddress [put]
func UpdateAddress(c *gin.Context) {
	var address model.Address
	_ = c.ShouldBindJSON(&address)
	err := service.UpdateAddress(&address)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Address
// @Summary 用id查询Address
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Address true "用id查询Address"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /address/findAddress [get]
func FindAddress(c *gin.Context) {
	var address model.Address
	_ = c.ShouldBindQuery(&address)
	err, readdress := service.GetAddress(address.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"readdress": readdress}, c)
	}
}

// @Tags Address
// @Summary 分页获取Address列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddressSearch true "分页获取Address列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /address/getAddressList [get]
func GetAddressList(c *gin.Context) {
	var pageInfo request.AddressSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetAddressInfoList(pageInfo)
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
