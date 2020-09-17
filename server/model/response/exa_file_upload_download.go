package response

import "panta/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}

type ExaFilesResponse struct {
	Files []model.ExaFileUploadAndDownload `json:"files"`
}
