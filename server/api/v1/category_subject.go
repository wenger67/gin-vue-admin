package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
)

func CreateSubject(c *gin.Context) {
	var subject model.CategorySubject
	_ = c.ShouldBindJSON(&subject)
	SubjectVerify := utils.Rules{
		"SubjectName": {utils.NotEmpty()},
	}
	SubjectVerifyErr := utils.Verify(subject, SubjectVerify)
	if SubjectVerifyErr != nil {
		response.FailWithMessage(SubjectVerifyErr.Error(), c)
		return
	}

	err := service.CreateSubject(subject)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("create category subject failed, %v", err), c)
	} else {
		response.OkWithMessage("create success", c)
	}
}

func DeleteSubject(c *gin.Context) {
	var subject model.CategorySubject
	_ = c.ShouldBindJSON(&subject)
	SubjectVerify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	SubjectVerifyErr := utils.Verify(subject.ID, SubjectVerify)
	if SubjectVerifyErr != nil {
		response.FailWithMessage(SubjectVerifyErr.Error(), c)
		return
	}
	err := service.DeleteSubject(subject)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("delete subject failed, %v", err), c)
	} else {
		response.OkWithMessage("delete success", c)
	}
}

func GetSubjectList(c *gin.Context) {
	var sp request.SearchSubjectParams
	_ = c.ShouldBindJSON(&sp)
	PageVerifyErr := utils.Verify(sp.PageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}

	err, list, total := service.GetSubjectList(sp.CategorySubject, sp.PageInfo, sp.OrderKey, sp.Desc)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("get subject list failed, %v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List: list,
			Total: total,
			Page: sp.PageInfo.Page,
			PageSize: sp.PageInfo.PageSize,
		}, c)
	}
}

func GetSubjectById(c *gin.Context) {
	var id request.GetById
	_ = c.ShouldBindJSON(&id)
	IdVerifyErr := utils.Verify(id, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err, subject := service.GetSubjectById(id.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("get subject failed, %v", err), c)
	} else {
		response.OkWithData(resp.SubjectResponse{Subject: subject}, c)
	}
}

func UpdateSubject(c *gin.Context) {
	var subject model.CategorySubject
	_ = c.ShouldBindJSON(&subject)
	SubjectVerify := utils.Rules{"SubjectName": {utils.NotEmpty()}}
	SubjectVerifyErr := utils.Verify(subject, SubjectVerify)
	if SubjectVerifyErr != nil {
		response.FailWithMessage(SubjectVerifyErr.Error(), c)
		return
	}
	err := service.UpdateSubject(subject)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("update subject failed, %v", err), c)
	} else {
		response.OkWithMessage("update subject success", c)
	}
}

func GetAllSubjects(c *gin.Context) {
	err, subjects := service.GetAllSubjects()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("get all subjects failed , %v", err), c)
	} else {
		response.OkWithData(resp.SubjectListResponse{SubjectList: subjects}, c)
	}
}