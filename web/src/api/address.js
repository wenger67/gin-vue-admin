import service from '@/utils/request'

// @Tags Address
// @Summary 创建Address
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Address true "创建Address"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /address/createAddress [post]
export const createAddress = (data) => {
     return service({
         url: "/address/createAddress",
         method: 'post',
         data
     })
 }


// @Tags Address
// @Summary 删除Address
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Address true "删除Address"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /address/deleteAddress [delete]
 export const deleteAddress = (data) => {
     return service({
         url: "/address/deleteAddress",
         method: 'delete',
         data
     })
 }

// @Tags Address
// @Summary 删除Address
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Address"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /address/deleteAddress [delete]
 export const deleteAddressByIds = (data) => {
     return service({
         url: "/address/deleteAddressByIds",
         method: 'delete',
         data
     })
 }

// @Tags Address
// @Summary 更新Address
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Address true "更新Address"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /address/updateAddress [put]
 export const updateAddress = (data) => {
     return service({
         url: "/address/updateAddress",
         method: 'put',
         data
     })
 }


// @Tags Address
// @Summary 用id查询Address
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Address true "用id查询Address"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /address/findAddress [get]
 export const findAddress = (params) => {
     return service({
         url: "/address/findAddress",
         method: 'get',
         params
     })
 }


// @Tags Address
// @Summary 分页获取Address列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Address列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /address/getAddressList [get]
 export const getAddressList = (params) => {
     return service({
         url: "/address/getAddressList",
         method: 'get',
         params
     })
 }