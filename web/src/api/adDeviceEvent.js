import service from '@/utils/request'

// @Tags AdDeviceEvent
// @Summary 创建AdDeviceEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceEvent true "创建AdDeviceEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDeviceEvent/createAdDeviceEvent [post]
export const createAdDeviceEvent = (data) => {
     return service({
         url: "/adDeviceEvent/createAdDeviceEvent",
         method: 'post',
         data
     })
 }


// @Tags AdDeviceEvent
// @Summary 删除AdDeviceEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceEvent true "删除AdDeviceEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDeviceEvent/deleteAdDeviceEvent [delete]
 export const deleteAdDeviceEvent = (data) => {
     return service({
         url: "/adDeviceEvent/deleteAdDeviceEvent",
         method: 'delete',
         data
     })
 }

// @Tags AdDeviceEvent
// @Summary 删除AdDeviceEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdDeviceEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDeviceEvent/deleteAdDeviceEvent [delete]
 export const deleteAdDeviceEventByIds = (data) => {
     return service({
         url: "/adDeviceEvent/deleteAdDeviceEventByIds",
         method: 'delete',
         data
     })
 }

// @Tags AdDeviceEvent
// @Summary 更新AdDeviceEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceEvent true "更新AdDeviceEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adDeviceEvent/updateAdDeviceEvent [put]
 export const updateAdDeviceEvent = (data) => {
     return service({
         url: "/adDeviceEvent/updateAdDeviceEvent",
         method: 'put',
         data
     })
 }


// @Tags AdDeviceEvent
// @Summary 用id查询AdDeviceEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceEvent true "用id查询AdDeviceEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adDeviceEvent/findAdDeviceEvent [get]
 export const findAdDeviceEvent = (params) => {
     return service({
         url: "/adDeviceEvent/findAdDeviceEvent",
         method: 'get',
         params
     })
 }


// @Tags AdDeviceEvent
// @Summary 分页获取AdDeviceEvent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AdDeviceEvent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDeviceEvent/getAdDeviceEventList [get]
 export const getAdDeviceEventList = (params) => {
     return service({
         url: "/adDeviceEvent/getAdDeviceEventList",
         method: 'get',
         params
     })
 }