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

// @Tags Categories
// @Summary 创建Categories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category true "创建Categories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /categories/createCategories [post]
func CreateCategories(c *gin.Context) {
	var categories model.Category
	_ = c.ShouldBindJSON(&categories)
	err := service.CreateCategories(categories)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Categories
// @Summary 删除Categories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category true "删除Categories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /categories/deleteCategories [delete]
func DeleteCategories(c *gin.Context) {
	var categories model.Category
	_ = c.ShouldBindJSON(&categories)
	err := service.DeleteCategories(categories)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Categories
// @Summary 批量删除Categories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Categories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /categories/deleteCategoriesByIds [delete]
func DeleteCategoriesByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteCategoriesByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Categories
// @Summary 更新Categories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category true "更新Categories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /categories/updateCategories [put]
func UpdateCategories(c *gin.Context) {
	var categories model.Category
	_ = c.ShouldBindJSON(&categories)
	err := service.UpdateCategories(&categories)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Categories
// @Summary 用id查询Categories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category true "用id查询Categories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /categories/findCategories [get]
func FindCategories(c *gin.Context) {
	var categories model.Category
	_ = c.ShouldBindQuery(&categories)
	err, recategories := service.GetCategories(categories.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"recategories": recategories}, c)
	}
}

// @Tags Categories
// @Summary 分页获取Categories列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CategorySearch true "分页获取Categories列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /categories/getCategoriesList [get]
func GetCategoriesList(c *gin.Context) {
	var sp request.CategorySearch
	_ = c.ShouldBindQuery(&sp)
	err, list, total := service.GetCategoriesInfoList(sp)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     sp.Page,
			PageSize: sp.PageSize,
		}, c)
	}
}
