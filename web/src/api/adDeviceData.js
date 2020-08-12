import service from '@/utils/request'

// @Tags AdDeviceData
// @Summary 创建AdDeviceData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceData true "创建AdDeviceData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDeviceData/createAdDeviceData [post]
export const createAdDeviceData = (data) => {
     return service({
         url: "/adDeviceData/createAdDeviceData",
         method: 'post',
         data
     })
 }


// @Tags AdDeviceData
// @Summary 删除AdDeviceData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceData true "删除AdDeviceData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDeviceData/deleteAdDeviceData [delete]
 export const deleteAdDeviceData = (data) => {
     return service({
         url: "/adDeviceData/deleteAdDeviceData",
         method: 'delete',
         data
     })
 }

// @Tags AdDeviceData
// @Summary 删除AdDeviceData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdDeviceData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDeviceData/deleteAdDeviceData [delete]
 export const deleteAdDeviceDataByIds = (data) => {
     return service({
         url: "/adDeviceData/deleteAdDeviceDataByIds",
         method: 'delete',
         data
     })
 }

// @Tags AdDeviceData
// @Summary 更新AdDeviceData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceData true "更新AdDeviceData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adDeviceData/updateAdDeviceData [put]
 export const updateAdDeviceData = (data) => {
     return service({
         url: "/adDeviceData/updateAdDeviceData",
         method: 'put',
         data
     })
 }


// @Tags AdDeviceData
// @Summary 用id查询AdDeviceData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceData true "用id查询AdDeviceData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adDeviceData/findAdDeviceData [get]
 export const findAdDeviceData = (params) => {
     return service({
         url: "/adDeviceData/findAdDeviceData",
         method: 'get',
         params
     })
 }


// @Tags AdDeviceData
// @Summary 分页获取AdDeviceData列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AdDeviceData列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDeviceData/getAdDeviceDataList [get]
 export const getAdDeviceDataList = (params) => {
     return service({
         url: "/adDeviceData/getAdDeviceDataList",
         method: 'get',
         params
     })
 }