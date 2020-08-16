package dev

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/service/dev"
	"github.com/gin-gonic/gin"
)

func FindDevice(c *gin.Context) {
	var param model.Lift
	_ = c.ShouldBindQuery(&param)
	err, lift := dev.FindLift(param.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"lift": lift}, c)
	}
}
