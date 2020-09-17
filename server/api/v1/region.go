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

// @Tags Region
// @Summary 创建Region
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Region true "创建Region"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /region/createRegion [post]
func CreateRegion(c *gin.Context) {
	var region model.Region
	_ = c.ShouldBindJSON(&region)
	err := service.CreateRegion(region)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Region
// @Summary 删除Region
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Region true "删除Region"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /region/deleteRegion [delete]
func DeleteRegion(c *gin.Context) {
	var region model.Region
	_ = c.ShouldBindJSON(&region)
	err := service.DeleteRegion(region)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Region
// @Summary 批量删除Region
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Region"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /region/deleteRegionByIds [delete]
func DeleteRegionByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteRegionByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Region
// @Summary 更新Region
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Region true "更新Region"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /region/updateRegion [put]
func UpdateRegion(c *gin.Context) {
	var region model.Region
	_ = c.ShouldBindJSON(&region)
	err := service.UpdateRegion(&region)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Region
// @Summary 用id查询Region
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Region true "用id查询Region"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /region/findRegion [get]
func FindRegion(c *gin.Context) {
	var region model.Region
	_ = c.ShouldBindQuery(&region)
	err, reregion := service.GetRegion(region.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reregion": reregion}, c)
	}
}

// @Tags Region
// @Summary 分页获取Region列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.RegionSearch true "分页获取Region列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /region/getRegionList [get]
func GetRegionList(c *gin.Context) {
	var pageInfo request.RegionSearch
	_ = c.ShouldBindQuery(&pageInfo)

	err, list, total := service.GetRegionInfoList(pageInfo)
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
