package dev

import (
	"encoding/json"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/service"
	"gin-vue-admin/service/dev"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
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
	imageItems := handleFiles(c, images, noSave, storage, event)
	item := request.DeviceEventContentItem{Brief: "start", Images: imageItems, CreatedAt: time.Now()}
	content.Items = append(content.Items, item)
	videoItems := handleFiles(c, videos, noSave, storage, event)
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

func handleFiles(c *gin.Context, files []*multipart.FileHeader, noSave string, storage string, event model.AdDeviceEvent) (list []string){
	for _, fileHeader := range files {
		var err error
		var filePath string
		var key string
		if storage == "qiniu" {
			// 文件上传后拿到文件路径
			err, filePath, key = utils.Upload(fileHeader)
		} else {
			// upload local
			dir, _ := os.Getwd()
			fileName := fileHeader.Filename
			date := time.Now().Format("2006_01_02")
			dateTime := time.Now().Format("2006_01_02_15_04_05")
			prefix := dir + "/resource"
			additionPath := "/upload/" + date + "/" + strconv.Itoa(int(event.DeviceID)) + "/" +
				strconv.Itoa(int(event.TypeId)) + "/" + strconv.Itoa(int(event.ID)) + "/"
			newFileName := utils.MD5V([]byte(dateTime)) + "_" + fileName
			err = os.MkdirAll(prefix + additionPath, os.ModePerm)
			if err != nil {
				global.PantaLog.Warning("create dir failed", err)
			}
			dst := prefix + additionPath + newFileName
			err = c.SaveUploadedFile(fileHeader, dst)
			if err == nil {
				filePath = "http://127.0.0.1:8888" + additionPath + newFileName
			} else {
				global.PantaLog.Warning("save file failed ", err)
			}
		}

		if err != nil {
			global.PantaLog.Warning("get file path failed " , fileHeader.Filename)
		} else {
			// 修改数据库后得到修改后的user并且返回供前端使用
			var file model.ExaFileUploadAndDownload
			file.Url = filePath
			file.Name = fileHeader.Filename
			s := strings.Split(file.Name, ".")
			file.Tag = s[len(s)-1]
			file.Key = key
			if noSave == "0" {
				err = service.Upload(file)
			}
			if err != nil {
				global.PantaLog.Warning(fmt.Sprintf("修改数据库链接失败，%v", err))
			} else {
				list = append(list, filePath)
			}
		}
	}
	return
}