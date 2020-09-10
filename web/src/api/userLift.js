import service from '@/utils/request'

// @Tags UserLifts
// @Summary 创建UserLifts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserLifts true "创建UserLifts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userLift/createUserLifts [post]
export const createUserLift = (data) => {
     return service({
         url: "/userLift/createUserLift",
         method: 'post',
         data
     })
 }


// @Tags UserLifts
// @Summary 删除UserLifts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserLifts true "删除UserLifts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userLift/deleteUserLifts [delete]
 export const deleteUserLift = (data) => {
     return service({
         url: "/userLift/deleteUserLift",
         method: 'delete',
         data
     })
 }

// @Tags UserLifts
// @Summary 删除UserLifts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserLifts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userLift/deleteUserLifts [delete]
 export const deleteUserLiftByIds = (data) => {
     return service({
         url: "/userLift/deleteUserLiftByIds",
         method: 'delete',
         data
     })
 }

// @Tags UserLifts
// @Summary 更新UserLifts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserLifts true "更新UserLifts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userLift/updateUserLifts [put]
 export const updateUserLift = (data) => {
     return service({
         url: "/userLift/updateUserLift",
         method: 'put',
         data
     })
 }


// @Tags UserLifts
// @Summary 用id查询UserLifts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserLifts true "用id查询UserLifts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userLift/findUserLifts [get]
 export const findUserLift = (params) => {
     return service({
         url: "/userLift/findUserLift",
         method: 'get',
         params
     })
 }


// @Tags UserLifts
// @Summary 分页获取UserLifts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取UserLifts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userLift/getUserLiftsList [get]
 export const getUserLiftList = (params) => {
     return service({
         url: "/userLift/getUserLiftList",
         method: 'get',
         params
     })
 }