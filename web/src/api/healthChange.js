import service from '@/utils/request'

// @Tags HealthChange
// @Summary 创建HealthChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthChange true "创建HealthChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /healthChange/createHealthChange [post]
export const createHealthChange = (data) => {
     return service({
         url: "/healthChange/createHealthChange",
         method: 'post',
         data
     })
 }


// @Tags HealthChange
// @Summary 删除HealthChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthChange true "删除HealthChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /healthChange/deleteHealthChange [delete]
 export const deleteHealthChange = (data) => {
     return service({
         url: "/healthChange/deleteHealthChange",
         method: 'delete',
         data
     })
 }

// @Tags HealthChange
// @Summary 删除HealthChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HealthChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /healthChange/deleteHealthChange [delete]
 export const deleteHealthChangeByIds = (data) => {
     return service({
         url: "/healthChange/deleteHealthChangeByIds",
         method: 'delete',
         data
     })
 }

// @Tags HealthChange
// @Summary 更新HealthChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthChange true "更新HealthChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /healthChange/updateHealthChange [put]
 export const updateHealthChange = (data) => {
     return service({
         url: "/healthChange/updateHealthChange",
         method: 'put',
         data
     })
 }


// @Tags HealthChange
// @Summary 用id查询HealthChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthChange true "用id查询HealthChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /healthChange/findHealthChange [get]
 export const findHealthChange = (params) => {
     return service({
         url: "/healthChange/findHealthChange",
         method: 'get',
         params
     })
 }


// @Tags HealthChange
// @Summary 分页获取HealthChange列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取HealthChange列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /healthChange/getHealthChangeList [get]
 export const getHealthChangeList = (params) => {
     return service({
         url: "/healthChange/getHealthChangeList",
         method: 'get',
         params
     })
 }