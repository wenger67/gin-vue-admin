package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils/enum"
	"github.com/gin-gonic/gin"
)

// @Tags UserLift
// @Summary 创建UserLift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserLift true "创建UserLift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userLift/createUserLift [post]
func CreateUserLift(c *gin.Context) {
	var userLift model.UserLift
	_ = c.ShouldBindJSON(&userLift)
	err := service.CreateUserLift(userLift)
	// TODO send message
	categoryId := userLift.CategoryId
	_, lift := service.GetLift(userLift.LiftId)

	switch categoryId {
	case uint(enum.UserLiftMaintain):
		// lift add maintain worker
		_ = service.CreateMessage(model.Message{FromId:enum.SystemNotifyUserId, TargetId: lift.Maintainer.Admin.ID,
			Content: "电梯" + lift.NickName + "新增维保人员", TypeId: uint(enum.MessageNewMaintainWorker)})
		_ = service.CreateMessage(model.Message{FromId:enum.SystemNotifyUserId, TargetId: lift.Owner.Admin.ID,
			Content: "电梯" + lift.NickName + "新增维保人员", TypeId: uint(enum.MessageNewMaintainWorker)})
		break
	case uint(enum.UserLiftCheck):
		_ = service.CreateMessage(model.Message{FromId:enum.SystemNotifyUserId, TargetId: lift.Checker.Admin.ID,
			Content:"电梯" + lift.NickName + "新增维保人员", TypeId: uint(enum.MessageNewCheckWorker)})
		_ = service.CreateMessage(model.Message{FromId:enum.SystemNotifyUserId, TargetId: lift.Owner.Admin.ID,
			Content: "电梯" + lift.NickName + "新增维保人员", TypeId: uint(enum.MessageNewCheckWorker)})
		break
	case uint(enum.UserLiftInstall):
		_ = service.CreateMessage(model.Message{FromId: enum.SystemNotifyUserId, TargetId: lift.Installer.Admin.ID,
			Content: "电梯" + lift.NickName + "新增安装人员", TypeId: uint(enum.MessageNewInstallWorker)})
		_ = service.CreateMessage(model.Message{FromId: enum.SystemNotifyUserId, TargetId: lift.Owner.Admin.ID,
			Content: "电梯" + lift.NickName + "新增安装人员", TypeId: uint(enum.MessageNewInstallWorker)})
		break
	case uint(enum.UserLiftManage):
		_ = service.CreateMessage(model.Message{FromId: enum.SystemNotifyUserId, TargetId: lift.Owner.Admin.ID,
			Content: "电梯" + lift.NickName + "新增管理人员", TypeId: uint(enum.MessageNewInstallWorker)})
		break
	case uint(enum.UserLiftSupervise):
		_ = service.CreateMessage(model.Message{FromId: enum.SystemNotifyUserId, TargetId: lift.Owner.Admin.ID,
			Content: "电梯" + lift.NickName + "新增监督人员", TypeId: uint(enum.MessageNewInstallWorker)})
		break
	}

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags UserLift
// @Summary 删除UserLift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserLift true "删除UserLift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userLift/deleteUserLift [delete]
func DeleteUserLift(c *gin.Context) {
	var userLift model.UserLift
	_ = c.ShouldBindJSON(&userLift)
	err := service.DeleteUserLift(userLift)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags UserLift
// @Summary 批量删除UserLift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserLift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userLift/deleteUserLiftByIds [delete]
func DeleteUserLiftByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteUserLiftByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags UserLift
// @Summary 更新UserLift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserLift true "更新UserLift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userLift/updateUserLift [put]
func UpdateUserLift(c *gin.Context) {
	var userLift model.UserLift
	_ = c.ShouldBindJSON(&userLift)
	err := service.UpdateUserLift(&userLift)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags UserLift
// @Summary 用id查询UserLift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserLift true "用id查询UserLift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userLift/findUserLift [get]
func FindUserLift(c *gin.Context) {
	var userLift model.UserLift
	_ = c.ShouldBindQuery(&userLift)
	err, reuserLift := service.GetUserLift(userLift.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reuserLift": reuserLift}, c)
	}
}

// @Tags UserLift
// @Summary 分页获取UserLift列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UserLiftSearch true "分页获取UserLift列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userLift/getUserLiftList [get]
func GetUserLiftList(c *gin.Context) {
	var pageInfo request.UserLiftSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetUserLiftList(pageInfo)
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
