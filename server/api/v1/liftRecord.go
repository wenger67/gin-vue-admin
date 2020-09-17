package v1

import (
	"fmt"
	"panta/global/response"
	"panta/model"
	"panta/model/request"
	resp "panta/model/response"
	"panta/service"
	"github.com/gin-gonic/gin"
	"time"
)

// @Tags LiftRecord
// @Summary 创建LiftRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftRecord true "创建LiftRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftRecord/createLiftRecord [post]
func CreateLiftRecord(c *gin.Context) {
	var create request.LiftRecordCreate
	_ = c.ShouldBindJSON(&create)

	liftRecord := model.LiftRecord{LiftId: create.LiftId, CategoryId: create.CategoryId,
		StartTime: time.Now(), WorkerId: create.WorkerId}

	err := service.CreateLiftRecord(liftRecord)
	// TODO send message
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func FillLiftRecord(c *gin.Context) {
	var fill request.LiftRecordFill
	_ = c.ShouldBindJSON(&fill)

	err := service.FillLiftRecord(fill)
	// TODO send message
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func ReviewLiftRecord(c *gin.Context) {
	var review request.LiftRecordReview
	_ = c.ShouldBindJSON(&review)

	err := service.ReviewLiftRecord(review)
	// TODO send message
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags LiftRecord
// @Summary 删除LiftRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftRecord true "删除LiftRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftRecord/deleteLiftRecord [delete]
func DeleteLiftRecord(c *gin.Context) {
	var liftRecord model.LiftRecord
	_ = c.ShouldBindJSON(&liftRecord)
	err := service.DeleteLiftRecord(liftRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags LiftRecord
// @Summary 批量删除LiftRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LiftRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftRecord/deleteLiftRecordByIds [delete]
func DeleteLiftRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteLiftRecordByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags LiftRecord
// @Summary 更新LiftRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftRecord true "更新LiftRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /liftRecord/updateLiftRecord [put]
func UpdateLiftRecord(c *gin.Context) {
	var liftRecord model.LiftRecord
	_ = c.ShouldBindJSON(&liftRecord)
	err := service.UpdateLiftRecord(&liftRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags LiftRecord
// @Summary 用id查询LiftRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftRecord true "用id查询LiftRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /liftRecord/findLiftRecord [get]
func FindLiftRecord(c *gin.Context) {
	var liftRecord model.LiftRecord
	_ = c.ShouldBindQuery(&liftRecord)
	err, reliftRecord := service.GetLiftRecord(liftRecord.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reliftRecord": reliftRecord}, c)
	}
}

// @Tags LiftRecord
// @Summary 分页获取LiftRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LiftRecordSearch true "分页获取LiftRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftRecord/getLiftRecordList [get]
func GetLiftRecordList(c *gin.Context) {
	var pageInfo request.LiftRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetLiftRecordInfoList(pageInfo)
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
