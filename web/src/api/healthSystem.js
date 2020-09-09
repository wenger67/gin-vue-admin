import service from '@/utils/request'

// @Tags HealthSystem
// @Summary 创建HealthSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthSystem true "创建HealthSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /healthSystem/createHealthSystem [post]
export const createHealthSystem = (data) => {
     return service({
         url: "/healthSystem/createHealthSystem",
         method: 'post',
         data
     })
 }


// @Tags HealthSystem
// @Summary 删除HealthSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthSystem true "删除HealthSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /healthSystem/deleteHealthSystem [delete]
 export const deleteHealthSystem = (data) => {
     return service({
         url: "/healthSystem/deleteHealthSystem",
         method: 'delete',
         data
     })
 }

// @Tags HealthSystem
// @Summary 删除HealthSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HealthSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /healthSystem/deleteHealthSystem [delete]
 export const deleteHealthSystemByIds = (data) => {
     return service({
         url: "/healthSystem/deleteHealthSystemByIds",
         method: 'delete',
         data
     })
 }

// @Tags HealthSystem
// @Summary 更新HealthSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthSystem true "更新HealthSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /healthSystem/updateHealthSystem [put]
 export const updateHealthSystem = (data) => {
     return service({
         url: "/healthSystem/updateHealthSystem",
         method: 'put',
         data
     })
 }


// @Tags HealthSystem
// @Summary 用id查询HealthSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HealthSystem true "用id查询HealthSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /healthSystem/findHealthSystem [get]
 export const findHealthSystem = (params) => {
     return service({
         url: "/healthSystem/findHealthSystem",
         method: 'get',
         params
     })
 }


// @Tags HealthSystem
// @Summary 分页获取HealthSystem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取HealthSystem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /healthSystem/getHealthSystemList [get]
 export const getHealthSystemList = (params) => {
     return service({
         url: "/healthSystem/getHealthSystemList",
         method: 'get',
         params
     })
 }