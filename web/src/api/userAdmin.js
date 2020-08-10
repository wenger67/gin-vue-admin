import service from '@/utils/request'

// @Tags UserAdmin
// @Summary 创建UserAdmin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAdmin true "创建UserAdmin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAdmin/createUserAdmin [post]
export const createUserAdmin = (data) => {
     return service({
         url: "/userAdmin/createUserAdmin",
         method: 'post',
         data
     })
 }


// @Tags UserAdmin
// @Summary 删除UserAdmin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAdmin true "删除UserAdmin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAdmin/deleteUserAdmin [delete]
 export const deleteUserAdmin = (data) => {
     return service({
         url: "/userAdmin/deleteUserAdmin",
         method: 'delete',
         data
     })
 }

// @Tags UserAdmin
// @Summary 删除UserAdmin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserAdmin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAdmin/deleteUserAdmin [delete]
 export const deleteUserAdminByIds = (data) => {
     return service({
         url: "/userAdmin/deleteUserAdminByIds",
         method: 'delete',
         data
     })
 }

// @Tags UserAdmin
// @Summary 更新UserAdmin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAdmin true "更新UserAdmin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userAdmin/updateUserAdmin [put]
 export const updateUserAdmin = (data) => {
     return service({
         url: "/userAdmin/updateUserAdmin",
         method: 'put',
         data
     })
 }


// @Tags UserAdmin
// @Summary 用id查询UserAdmin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAdmin true "用id查询UserAdmin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userAdmin/findUserAdmin [get]
 export const findUserAdmin = (params) => {
     return service({
         url: "/userAdmin/findUserAdmin",
         method: 'get',
         params
     })
 }


// @Tags UserAdmin
// @Summary 分页获取UserAdmin列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取UserAdmin列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAdmin/getUserAdminList [get]
 export const getUserAdminList = (params) => {
     return service({
         url: "/userAdmin/getUserAdminList",
         method: 'get',
         params
     })
 }