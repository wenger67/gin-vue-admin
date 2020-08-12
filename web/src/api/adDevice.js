import service from '@/utils/request'

// @Tags AdDevice
// @Summary 创建AdDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDevice true "创建AdDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDevice/createAdDevice [post]
export const createAdDevice = (data) => {
     return service({
         url: "/adDevice/createAdDevice",
         method: 'post',
         data
     })
 }


// @Tags AdDevice
// @Summary 删除AdDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDevice true "删除AdDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDevice/deleteAdDevice [delete]
 export const deleteAdDevice = (data) => {
     return service({
         url: "/adDevice/deleteAdDevice",
         method: 'delete',
         data
     })
 }

// @Tags AdDevice
// @Summary 删除AdDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDevice/deleteAdDevice [delete]
 export const deleteAdDeviceByIds = (data) => {
     return service({
         url: "/adDevice/deleteAdDeviceByIds",
         method: 'delete',
         data
     })
 }

// @Tags AdDevice
// @Summary 更新AdDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDevice true "更新AdDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adDevice/updateAdDevice [put]
 export const updateAdDevice = (data) => {
     return service({
         url: "/adDevice/updateAdDevice",
         method: 'put',
         data
     })
 }


// @Tags AdDevice
// @Summary 用id查询AdDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDevice true "用id查询AdDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adDevice/findAdDevice [get]
 export const findAdDevice = (params) => {
     return service({
         url: "/adDevice/findAdDevice",
         method: 'get',
         params
     })
 }


// @Tags AdDevice
// @Summary 分页获取AdDevice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AdDevice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDevice/getAdDeviceList [get]
 export const getAdDeviceList = (params) => {
     return service({
         url: "/adDevice/getAdDeviceList",
         method: 'get',
         params
     })
 }