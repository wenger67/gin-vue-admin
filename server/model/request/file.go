package request

import "panta/model"

type FileListRequest struct {
	PageInfo
	model.ExaFileUploadAndDownload
}

type DeleteFileRequest struct {
	model.ExaFileUploadAndDownload
	Storage string `json:"storage"`
}