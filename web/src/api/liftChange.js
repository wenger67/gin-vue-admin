import service from '@/utils/request'

// @Tags LiftChange
// @Summary 创建LiftChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftChange true "创建LiftChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftChange/createLiftChange [post]
export const createLiftChange = (data) => {
     return service({
         url: "/liftChange/createLiftChange",
         method: 'post',
         data
     })
 }


// @Tags LiftChange
// @Summary 删除LiftChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftChange true "删除LiftChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftChange/deleteLiftChange [delete]
 export const deleteLiftChange = (data) => {
     return service({
         url: "/liftChange/deleteLiftChange",
         method: 'delete',
         data
     })
 }

// @Tags LiftChange
// @Summary 删除LiftChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LiftChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftChange/deleteLiftChange [delete]
 export const deleteLiftChangeByIds = (data) => {
     return service({
         url: "/liftChange/deleteLiftChangeByIds",
         method: 'delete',
         data
     })
 }

// @Tags LiftChange
// @Summary 更新LiftChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftChange true "更新LiftChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /liftChange/updateLiftChange [put]
 export const updateLiftChange = (data) => {
     return service({
         url: "/liftChange/updateLiftChange",
         method: 'put',
         data
     })
 }


// @Tags LiftChange
// @Summary 用id查询LiftChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftChange true "用id查询LiftChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /liftChange/findLiftChange [get]
 export const findLiftChange = (params) => {
     return service({
         url: "/liftChange/findLiftChange",
         method: 'get',
         params
     })
 }


// @Tags LiftChange
// @Summary 分页获取LiftChange列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取LiftChange列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftChange/getLiftChangeList [get]
 export const getLiftChangeList = (params) => {
     return service({
         url: "/liftChange/getLiftChangeList",
         method: 'get',
         params
     })
 }