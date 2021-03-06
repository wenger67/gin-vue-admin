package v1

import (
	"fmt"
	"panta/global"
	"panta/global/response"
	"panta/model"
	"panta/model/request"
	resp "panta/model/response"
	"panta/service"
	"panta/utils"
	"panta/utils/enum"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Tags Lift
// @Summary 创建Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lift true "创建Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lift/createLift [post]
func CreateLift(c *gin.Context) {
	var lift model.Lift
	_ = c.ShouldBindJSON(&lift)
	err := service.CreateLift(&lift)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		// health system initial
		timeDim := utils.CalTimeDimensionInitVal(lift.InstallTime)
		health := model.HealthSystem{LiftId: int(lift.ID), TimeDimension: timeDim,
			HumanDimension: utils.HealthDimensionInitialValue, InnerDimension: utils.HealthDimensionInitialValue,
			MaintainDimension: utils.HealthDimensionInitialValue, SensorDimension: utils.HealthDimensionInitialValue}
		err = service.CreateHealthSystem(health)
		// health change initial
		err = service.CreateHealthChange(model.HealthChange{LiftId: lift.ID,
			DimensionId: uint(enum.HealthTimeDimension),
			Content: "Initial", Score: timeDim})
		err = service.CreateHealthChange(model.HealthChange{LiftId: lift.ID,
			DimensionId: uint(enum.HealthHumanDimension),
			Content: "Initial", Score: utils.HealthDimensionInitialValue})
		err = service.CreateHealthChange(model.HealthChange{LiftId: lift.ID,
			DimensionId: uint(enum.HealthInnerDimension),
			Content: "Initial", Score: utils.HealthDimensionInitialValue})
		err = service.CreateHealthChange(model.HealthChange{LiftId: lift.ID,
			DimensionId: uint(enum.HealthMaintainDimension),
			Content: "Initial", Score: utils.HealthDimensionInitialValue})
		err = service.CreateHealthChange(model.HealthChange{LiftId: lift.ID,
			DimensionId: uint(enum.HealthSensorDimension),
			Content: "Initial", Score: utils.HealthDimensionInitialValue})
		// TODO send message
		if err != nil {
			response.FailWithMessage("init health system failed for "+strconv.Itoa(int(lift.ID)), c)
		} else {
			response.OkWithMessage("创建成功", c)
		}
	}
}

// @Tags Lift
// @Summary 删除Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lift true "删除Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lift/deleteLift [delete]
func DeleteLift(c *gin.Context) {
	var lift model.Lift
	_ = c.ShouldBindJSON(&lift)
	err := service.DeleteLift(lift)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		// delete relative health system
		ids := []int{int(lift.ID)}
		err = service.DeleteHealthSystemByLiftIds(request.IdsReq{Ids: ids})
		err = service.DeleteHealthChangeByLiftIds(request.IdsReq{Ids: ids})
		if err != nil {
			response.FailWithMessage("delete relative health system failed "+err.Error(), c)
		} else {
			response.OkWithMessage("删除成功", c)
		}
	}
}

// @Tags Lift
// @Summary 批量删除Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lift/deleteLiftByIds [delete]
func DeleteLiftByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteLiftByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Lift
// @Summary 更新Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lift true "更新Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lift/updateLift [put]
func UpdateLift(c *gin.Context) {
	var lift model.Lift
	_ = c.ShouldBindJSON(&lift)

	changes := ""                         // change columns
	_, origin := service.GetLift(lift.ID) // save origin data
	err := service.UpdateLift(&lift)      // update data
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		_, lift = service.GetLift(origin.ID) // get updated data to compare
		// TODO lift info changes
		if origin.NickName != lift.NickName {
			changes += "nick name changes from [" + origin.NickName + "] to [" + lift.NickName + "], "
		}
		if origin.InstallerId != lift.InstallerId {
			changes += "installer changes from [" + origin.Installer.FullName + "] to [" + lift.Installer.FullName + "], "
		}
		if origin.MaintainerId != lift.MaintainerId {
			changes += "maintainer changes from [" + origin.Maintainer.FullName + "] to [" + lift.Maintainer.FullName + "], "
		}
		if origin.CheckerId != lift.CheckerId {
			changes += "checker changes from [" + origin.Checker.FullName + "] to [" + lift.Checker.FullName + "], "
		}
		if origin.OwnerId != lift.OwnerId {
			changes += "owner changes from [" + origin.Owner.FullName + "] to [" + lift.Owner.FullName + "], "
		}
		if origin.CheckTime != lift.CheckTime {
			changes += "check time changes from [" + origin.CheckTime.String() + "] to [" + lift.CheckTime.String() + "], "
		}
		if origin.LiftModelId != lift.LiftModelId {
			changes += "lift model changes from [" + origin.LiftModel.Brand + "] to [" + lift.LiftModel.Brand + "], "
		}
		if origin.CategoryId != lift.CategoryId {
			changes += "lift category changes from [" + origin.Category.CategoryName + "] to [" + lift.Category.CategoryName + "], "
		}
		if origin.AddressId != lift.AddressId {
			changes += "lift address changes from [" + origin.Address.AddressName + "] to [" + lift.Address.AddressName + "], "
		}
		if origin.FloorCount != lift.FloorCount {
			changes += "lift floor count changes from [" + strconv.Itoa(int(origin.FloorCount)) + "] to [" + strconv.Itoa(int(lift.FloorCount)) + "], "
		}
		if origin.Cell != lift.Cell {
			changes += "lift cell changes from [" + strconv.Itoa(int(origin.Cell)) + "] to [" + strconv.Itoa(int(lift.Cell)) + "], "
		}
		if origin.Building != lift.Building {
			changes += "lift building changes from [" + origin.Building + "] to [" + lift.Building + "], "
		}
		if origin.AdDevice.ID != lift.AdDevice.ID {
			changes += "lift smart device changes from [" + strconv.Itoa(int(origin.AdDevice.ID)) + "] to [" + strconv.Itoa(int(origin.AdDevice.ID)) + "], "
		}

		if changes != "" {
			var liftChange model.LiftChange
			liftChange.LiftId = int(lift.ID)
			liftChange.Content = changes
			err := service.CreateLiftChange(liftChange) // save changes
			// TODO send message

			if err != nil {
				global.PantaLog.Warning("create lift change failed.", err.Error())
			}
		} else {
			global.PantaLog.Debug("nothing changed")
		}
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Lift
// @Summary 用id查询Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lift true "用id查询Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lift/findLift [get]
func FindLift(c *gin.Context) {
	var lift model.Lift
	_ = c.ShouldBindQuery(&lift)
	err, relift := service.GetLift(lift.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"relift": relift}, c)
	}
}

// @Tags Lift
// @Summary 分页获取Lift列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LiftSearch true "分页获取Lift列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lift/getLiftList [get]
func GetLiftList(c *gin.Context) {
	var pageInfo request.LiftSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetLiftInfoList(pageInfo)
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
