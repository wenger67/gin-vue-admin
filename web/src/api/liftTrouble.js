import service from '@/utils/request'

// @Tags LiftTrouble
// @Summary 创建LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftTrouble true "创建LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftTrouble/createLiftTrouble [post]
export const createLiftTrouble = (data) => {
     return service({
         url: "/liftTrouble/createLiftTrouble",
         method: 'post',
         data
     })
 }


// @Tags LiftTrouble
// @Summary 删除LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftTrouble true "删除LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftTrouble/deleteLiftTrouble [delete]
 export const deleteLiftTrouble = (data) => {
     return service({
         url: "/liftTrouble/deleteLiftTrouble",
         method: 'delete',
         data
     })
 }

// @Tags LiftTrouble
// @Summary 删除LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftTrouble/deleteLiftTrouble [delete]
 export const deleteLiftTroubleByIds = (data) => {
     return service({
         url: "/liftTrouble/deleteLiftTroubleByIds",
         method: 'delete',
         data
     })
 }

// @Tags LiftTrouble
// @Summary 更新LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftTrouble true "更新LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /liftTrouble/updateLiftTrouble [put]
 export const updateLiftTrouble = (data) => {
     return service({
         url: "/liftTrouble/updateLiftTrouble",
         method: 'put',
         data
     })
 }


// @Tags LiftTrouble
// @Summary 用id查询LiftTrouble
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftTrouble true "用id查询LiftTrouble"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /liftTrouble/findLiftTrouble [get]
 export const findLiftTrouble = (params) => {
     return service({
         url: "/liftTrouble/findLiftTrouble",
         method: 'get',
         params
     })
 }


// @Tags LiftTrouble
// @Summary 分页获取LiftTrouble列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取LiftTrouble列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftTrouble/getLiftTroubleList [get]
 export const getLiftTroubleList = (params) => {
     return service({
         url: "/liftTrouble/getLiftTroubleList",
         method: 'get',
         params
     })
 }