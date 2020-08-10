import service from '@/utils/request'

// @Tags Lift
// @Summary 创建Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lift true "创建Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lift/createLift [post]
export const createLift = (data) => {
     return service({
         url: "/lift/createLift",
         method: 'post',
         data
     })
 }


// @Tags Lift
// @Summary 删除Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lift true "删除Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lift/deleteLift [delete]
 export const deleteLift = (data) => {
     return service({
         url: "/lift/deleteLift",
         method: 'delete',
         data
     })
 }

// @Tags Lift
// @Summary 删除Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lift/deleteLift [delete]
 export const deleteLiftByIds = (data) => {
     return service({
         url: "/lift/deleteLiftByIds",
         method: 'delete',
         data
     })
 }

// @Tags Lift
// @Summary 更新Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lift true "更新Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lift/updateLift [put]
 export const updateLift = (data) => {
     return service({
         url: "/lift/updateLift",
         method: 'put',
         data
     })
 }


// @Tags Lift
// @Summary 用id查询Lift
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lift true "用id查询Lift"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lift/findLift [get]
 export const findLift = (params) => {
     return service({
         url: "/lift/findLift",
         method: 'get',
         params
     })
 }


// @Tags Lift
// @Summary 分页获取Lift列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Lift列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lift/getLiftList [get]
 export const getLiftList = (params) => {
     return service({
         url: "/lift/getLiftList",
         method: 'get',
         params
     })
 }