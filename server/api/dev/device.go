package dev

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"panta/global/response"
	"panta/model"
	"panta/model/request"
	"panta/service"
	"panta/service/dev"
	"panta/utils/enum"
	"strconv"
	"time"
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


func CreateEvent(c *gin.Context) {
	noSave := c.DefaultQuery("noSave", "0")
	deviceId := c.Request.PostFormValue("deviceId")
	typeId := c.Request.PostFormValue("typeId")
	storage := c.Request.PostFormValue("storage")

	form, _ := c.MultipartForm()
	images := form.File["images"]
	videos := form.File["videos"]

	dId, _ := strconv.Atoi(deviceId)
	tId, _ := strconv.Atoi(typeId)

	// create event then updated with content
	event := model.AdDeviceEvent{DeviceID: uint(dId), TypeId: uint(tId)}
	err := dev.CreateEvent(&event)
	if err != nil {
		response.FailWithMessage("create device event failed", c)
		return
	}
	// create content about the event
	content := request.DeviceEventContent{}

	date := time.Now().Format("2006_01_02")
	additionPath := "/upload/" + date + "/device_" + strconv.Itoa(int(event.DeviceID)) + "/type_" +
				strconv.Itoa(int(event.TypeId)) + "/event_" + strconv.Itoa(int(event.ID)) + "/"
	_, imageItems := service.HandleFiles(c, images, noSave, storage, enum.FileBelongEvent, event.ID, additionPath)
	item := request.DeviceEventContentItem{Brief: "start", Images: imageItems, CreatedAt: time.Now()}
	content.Items = append(content.Items, item)
	_, videoItems := service.HandleFiles(c, videos, noSave, storage, enum.FileBelongEvent, event.ID, additionPath)
	content.Videos = videoItems

	temp, _ := json.Marshal(content);
	event.Content = string(temp)
	err = dev.UpdateEvent(&event)
	if err != nil {
		response.FailWithMessage("update event failed", c)
	} else {
		response.OkDetailed(event, "上传成功", c)
	}
}

