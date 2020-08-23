import service from '@/utils/request'

// @Tags UserLifts
// @Summary 创建UserLifts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserLifts true "创建UserLifts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userLift/createUserLifts [post]
export const createUserLifts = (data) => {
     return service({
         url: "/userLift/createUserLifts",
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
 export const deleteUserLifts = (data) => {
     return service({
         url: "/userLift/deleteUserLifts",
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
 export const deleteUserLiftsByIds = (data) => {
     return service({
         url: "/userLift/deleteUserLiftsByIds",
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
 export const updateUserLifts = (data) => {
     return service({
         url: "/userLift/updateUserLifts",
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
 export const findUserLifts = (params) => {
     return service({
         url: "/userLift/findUserLifts",
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
 export const getUserLiftsList = (params) => {
     return service({
         url: "/userLift/getUserLiftsList",
         method: 'get',
         params
     })
 }