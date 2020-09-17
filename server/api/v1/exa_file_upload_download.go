package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"strings"
	"syscall"
	"time"
)

// @Tags ExaFileUploadAndDownload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/upload [post]
func UploadFile(c *gin.Context) {
	noSave := c.DefaultQuery("noSave", "0")
	var storage string
	storage, _ = c.GetQuery("storage")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), c)
	} else {
		var err error
		var filePath string
		var key string
		if storage == "qiniu" {
			// 文件上传后拿到文件路径
			err, filePath, key = utils.Upload(header)
		} else {
			err, filePath, key = UploadLocal(c, header)
		}

		if err != nil {
			response.FailWithMessage(fmt.Sprintf("接收返回值失败，%v", err), c)
		} else {
			// 修改数据库后得到修改后的user并且返回供前端使用
			var file model.ExaFileUploadAndDownload
			file.Url = filePath
			file.Name = header.Filename
			s := strings.Split(file.Name, ".")
			file.Tag = s[len(s)-1]
			file.Key = key
			if noSave == "0" {
				err = service.Upload(file)
			}
			if err != nil {
				response.FailWithMessage(fmt.Sprintf("修改数据库链接失败，%v", err), c)
			} else {
				response.OkDetailed(resp.ExaFileResponse{File: file}, "上传成功", c)
			}
		}
	}
}

func UploadFiles(c *gin.Context)  {
	noSave := c.DefaultQuery("noSave", "0")
	var storage string
	storage, _ = c.GetQuery("storage")

	form, err := c.MultipartForm()
	if form == nil {
		global.PantaLog.Debug(form, err)
	}
	files := form.File["files"]
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), c)
	} else {
		fileList := make([]model.ExaFileUploadAndDownload, 0)
		for _, fileHeader := range files {
			var err error
			var filePath string
			var key string
			if storage == "qiniu" {
				// 文件上传后拿到文件路径
				err, filePath, key = utils.Upload(fileHeader)
			} else {
				err, filePath, key = UploadLocal(c, fileHeader)
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
					fileList = append(fileList, file)
				}
			}
		}
		if len(fileList) == 0 {
			response.FailWithMessage("upload file list all failed", c)
		} else {
			response.OkDetailed(resp.ExaFilesResponse{Files: fileList}, "上传成功", c)
		}
	}
}

func UploadLocal(c *gin.Context, header *multipart.FileHeader) (err error, path string, key string) {
	dir, _ := os.Getwd()
	lastIndex := strings.LastIndex(header.Filename, ".")
	fileName := header.Filename[:lastIndex] + time.Now().String()
	fileNameByte, _ := syscall.ByteSliceFromString(fileName)
	newName := utils.MD5V(fileNameByte)

	suffix := header.Filename[lastIndex:]
	dst :=  dir + "/resource/upload/" + newName + suffix
	err = c.SaveUploadedFile(header, dst)

	url := "http://127.0.0.1:8888/upload/" + newName +suffix
	// TODO save database
	return err, url, "local"
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /fileUploadAndDownload/deleteFile [post]
func DeleteFile(c *gin.Context) {
	var file model.ExaFileUploadAndDownload
	_ = c.ShouldBindJSON(&file)
	err, f := service.FindFile(file.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		err = utils.DeleteFile(f.Key)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)

		} else {
			err = service.DeleteFile(f)
			if err != nil {
				response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
			} else {
				response.OkWithMessage("删除成功", c)
			}
		}
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取文件户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileUploadAndDownload/getFileList [post]
func GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := service.GetFileRecordInfoList(pageInfo)
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
