package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"panta/global"
	"panta/model"
	"panta/model/request"
	"panta/utils"
	"panta/utils/enum"
	"strings"
	"time"
)

// @title    CreateFileRecord
// @description   创建文件上传记录
// @param     file            model.ExaFileUploadAndDownload
// @auth                     （2020/04/05  20:22）
// @return                    error

func CreateFileRecord(file *model.ExaFileUploadAndDownload) error {
	err := global.PantaDb.Create(file).Error
	return err
}

// @title    FindFile
// @description   删除文件切片记录
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error

func FindFile(id uint) (error, model.ExaFileUploadAndDownload) {
	var file model.ExaFileUploadAndDownload
	err := global.PantaDb.Where("id = ?", id).First(&file).Error
	return err, file
}

// @title    DeleteFileDatabaseRecord
// @description   删除文件记录
// @auth                     （2020/04/05  20:22）
// @param     file            model.ExaFileUploadAndDownload
// @return                    error

func DeleteFileDatabaseRecord(file model.ExaFileUploadAndDownload) error {
	err := global.PantaDb.Where("id = ?", file.ID).Unscoped().Delete(file).Error
	return err
}

// @title    GetFileRecordInfoList
// @description   分页获取数据
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return    err             error
// @return    list            error
// @return    total           error

func GetFileRecordInfoList(params request.FileListRequest) (err error, list []model.ExaFileUploadAndDownload, total int64) {
	limit := params.PageSize
	offset := params.PageSize * (params.Page - 1)
	db := global.PantaDb.Model(&model.ExaFileUploadAndDownload{})
	var fileLists []model.ExaFileUploadAndDownload

	err = db.Debug().Count(&total).Error
	err = db.Debug().Where("record_id = ? AND trouble_id = ? AND event_id = ?",
		params.RecordId, params.TroubleId, params.EventId).Limit(limit).Offset(offset).
		Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}

func Cleanup()  {
	_, files, _ := GetFileRecordInfoList(request.FileListRequest{
		PageInfo:request.PageInfo{Page: 1, PageSize: 9999},
		ExaFileUploadAndDownload: model.ExaFileUploadAndDownload{RecordId: 0, TroubleId: 0, EventId: 0}})
	for _, file := range files{
		global.PantaLog.Debug(file.Url)
		// err := utils.DeleteFileInDisk(file.Url)
		// if err != nil {
		// 	global.PantaLog.Warning("delete file in disk failed", err)
		// }
		// err = DeleteFileDatabaseRecord(file)
		// if err != nil {
		// 	global.PantaLog.Warning("delete file in database failed", err)
		// }
	}
}

/**
	1. save files into disk
	2. store file record into database
 */
func HandleFiles(c *gin.Context, files []*multipart.FileHeader, noSave string, storage string,
	belong string, belongId uint, additionPath string) (fileList []model.ExaFileUploadAndDownload, paths []string){
	for _, fileHeader := range files {
		var err error
		var filePath string
		var key string
		if storage == "qiniu" {
			// 文件上传后拿到文件路径
			err, filePath, key = utils.Upload(fileHeader)
		} else {
			// upload local
			root, _ := os.Getwd()
			prefix := root + "/resource"

			fileName := utils.MD5V([]byte(time.Now().Format("2006_01_02_15_04_05"))) + "_" + fileHeader.Filename
			err = os.MkdirAll(prefix + additionPath, os.ModePerm)
			if err != nil {
				global.PantaLog.Warning("create root failed", err)
			}
			dst := prefix + additionPath + fileName
			err = c.SaveUploadedFile(fileHeader, dst)
			if err == nil {
				filePath = "http://127.0.0.1:8888" + additionPath + fileName
			} else {
				global.PantaLog.Warning("save file failed ", err)
			}
		}

		if err != nil {
			global.PantaLog.Warning("get file path failed " , fileHeader.Filename)
		} else {
			var file model.ExaFileUploadAndDownload
			file.Url = filePath
			file.Name = fileHeader.Filename
			s := strings.Split(file.Name, ".")
			file.Tag = s[len(s)-1]
			file.Key = key
			// assign file belong to
			if belong == enum.FileBelongEvent {
				file.EventId = belongId
			} else if belong == enum.FileBelongRecord {
				file.RecordId = belongId
			} else if belong == enum.FileBelongTrouble {
				file.TroubleId = belongId
			}
			if noSave == "0" {
				err = CreateFileRecord(&file)
			}
			if err != nil {
				global.PantaLog.Warning(fmt.Sprintf("修改数据库链接失败，%v", err))
			} else {
				paths = append(paths, filePath)
				fileList = append(fileList, file)
			}
		}
	}
	return
}