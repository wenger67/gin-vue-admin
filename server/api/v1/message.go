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

// @Tags Message
// @Summary 创建Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Message true "创建Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /message/createMessage [post]
func CreateMessage(c *gin.Context) {
	var message model.Message
	_ = c.ShouldBindJSON(&message)
	err := service.CreateMessage(message)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Message
// @Summary 删除Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Message true "删除Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /message/deleteMessage [delete]
func DeleteMessage(c *gin.Context) {
	var message model.Message
	_ = c.ShouldBindJSON(&message)
	err := service.DeleteMessage(message)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Message
// @Summary 批量删除Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /message/deleteMessageByIds [delete]
func DeleteMessageByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteMessageByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Message
// @Summary 更新Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Message true "更新Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /message/updateMessage [put]
func UpdateMessage(c *gin.Context) {
	var message model.Message
	_ = c.ShouldBindJSON(&message)
	err := service.UpdateMessage(&message)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Message
// @Summary 用id查询Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Message true "用id查询Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /message/findMessage [get]
func FindMessage(c *gin.Context) {
	var message model.Message
	_ = c.ShouldBindQuery(&message)
	err, remessage := service.GetMessage(message.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"remessage": remessage}, c)
	}
}

// @Tags Message
// @Summary 分页获取Message列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MessageSearch true "分页获取Message列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /message/getMessageList [get]
func GetMessageList(c *gin.Context) {
	var pageInfo request.MessageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetMessageInfoList(pageInfo)
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
