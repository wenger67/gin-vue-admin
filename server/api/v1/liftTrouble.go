package v1

import (
	"fmt"
	"panta/global/response"
	"panta/model"
	"panta/model/request"
	resp "panta/model/response"
	"panta/service"
	"panta/utils/enum"
	"github.com/gin-gonic/gin"
	"time"
)

// @Tags LiftTrouble
// @Summary 创建LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftTrouble true "创建LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftTrouble/createLiftTrouble [post]
func CreateLiftTrouble(c *gin.Context) {
	var liftTrouble model.LiftTrouble
	_ = c.ShouldBindJSON(&liftTrouble)
	if liftTrouble.Progress == 0 {
		liftTrouble.Progress = 1 // goto step2
	}
	liftTrouble.StartTime = time.Now()
	err := service.CreateLiftTrouble(liftTrouble)
	// TODO send message
	_, lift := service.GetLift(liftTrouble.LiftId)
	if liftTrouble.FromCategoryId == uint(enum.LiftUrgentButtonTrouble) {
		// urgent button
		err = service.CreateMessage(model.Message{FromId: enum.SystemNotifyUserId, TargetId: lift.Maintainer.Admin.ID,
			Content: "触发电梯救援按钮", TypeId: uint(enum.MessageUrgent)})
		err = service.CreateMessage(model.Message{FromId: enum.SystemNotifyUserId, TargetId: lift.Owner.Admin.ID,
			Content: "触发电梯救援按钮", TypeId: uint(enum.MessageUrgent)})
		_, list, _ := service.GetUserInfoList(request.SearchUserParams{PageInfo:request.PageInfo{Page: 1,
			PageSize: 9999}, CompanyId: 15})
		users := list.([]model.SysUser)
		for _, val := range users {
			if val.ID == uint(enum.SystemNotifyUserId) {
				continue
			}
			err = service.CreateMessage(model.Message{FromId: enum.SystemNotifyUserId, TargetId: val.ID,
				Content: "触发电梯救援按钮", TypeId: uint(enum.MessageUrgent)})
		}
	}
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags LiftTrouble
// @Summary 删除LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftTrouble true "删除LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftTrouble/deleteLiftTrouble [delete]
func DeleteLiftTrouble(c *gin.Context) {
	var liftTrouble model.LiftTrouble
	_ = c.ShouldBindJSON(&liftTrouble)
	err := service.DeleteLiftTrouble(liftTrouble)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags LiftTrouble
// @Summary 批量删除LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftTrouble/deleteLiftTroubleByIds [delete]
func DeleteLiftTroubleByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteLiftTroubleByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags LiftTrouble
// @Summary 更新LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftTrouble true "更新LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /liftTrouble/updateLiftTrouble [put]
func UpdateLiftTrouble(c *gin.Context) {
	var params request.LiftTroubleUpdate
	_ = c.ShouldBindJSON(&params)
	err := service.UpdateLiftTrouble(&params)
	// TODO send message
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags LiftTrouble
// @Summary 用id查询LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftTrouble true "用id查询LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /liftTrouble/findLiftTrouble [get]
func FindLiftTrouble(c *gin.Context) {
	var liftTrouble model.LiftTrouble
	_ = c.ShouldBindQuery(&liftTrouble)
	err, reliftTrouble := service.GetLiftTrouble(liftTrouble.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reliftTrouble": reliftTrouble}, c)
	}
}

// @Tags LiftTrouble
// @Summary 分页获取LiftTrouble列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LiftTroubleSearch true "分页获取LiftTrouble列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftTrouble/getLiftTroubleList [get]
func GetLiftTroubleList(c *gin.Context) {
	var pageInfo request.LiftTroubleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetLiftTroubleInfoList(pageInfo)
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
