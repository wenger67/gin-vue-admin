import service from '@/utils/request'

// @Tags AdDeviceConfig
// @Summary 创建AdDeviceConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceConfig true "创建AdDeviceConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDeviceConfig/createAdDeviceConfig [post]
export const createAdDeviceConfig = (data) => {
     return service({
         url: "/adDeviceConfig/createAdDeviceConfig",
         method: 'post',
         data
     })
 }


// @Tags AdDeviceConfig
// @Summary 删除AdDeviceConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceConfig true "删除AdDeviceConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDeviceConfig/deleteAdDeviceConfig [delete]
 export const deleteAdDeviceConfig = (data) => {
     return service({
         url: "/adDeviceConfig/deleteAdDeviceConfig",
         method: 'delete',
         data
     })
 }

// @Tags AdDeviceConfig
// @Summary 删除AdDeviceConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdDeviceConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adDeviceConfig/deleteAdDeviceConfig [delete]
 export const deleteAdDeviceConfigByIds = (data) => {
     return service({
         url: "/adDeviceConfig/deleteAdDeviceConfigByIds",
         method: 'delete',
         data
     })
 }

// @Tags AdDeviceConfig
// @Summary 更新AdDeviceConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceConfig true "更新AdDeviceConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adDeviceConfig/updateAdDeviceConfig [put]
 export const updateAdDeviceConfig = (data) => {
     return service({
         url: "/adDeviceConfig/updateAdDeviceConfig",
         method: 'put',
         data
     })
 }


// @Tags AdDeviceConfig
// @Summary 用id查询AdDeviceConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdDeviceConfig true "用id查询AdDeviceConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adDeviceConfig/findAdDeviceConfig [get]
 export const findAdDeviceConfig = (params) => {
     return service({
         url: "/adDeviceConfig/findAdDeviceConfig",
         method: 'get',
         params
     })
 }


// @Tags AdDeviceConfig
// @Summary 分页获取AdDeviceConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AdDeviceConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adDeviceConfig/getAdDeviceConfigList [get]
 export const getAdDeviceConfigList = (params) => {
     return service({
         url: "/adDeviceConfig/getAdDeviceConfigList",
         method: 'get',
         params
     })
 }