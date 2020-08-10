import service from '@/utils/request'

// @Tags LiftModel
// @Summary 创建LiftModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftModel true "创建LiftModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftModel/createLiftModel [post]
export const createLiftModel = (data) => {
     return service({
         url: "/liftModel/createLiftModel",
         method: 'post',
         data
     })
 }


// @Tags LiftModel
// @Summary 删除LiftModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftModel true "删除LiftModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftModel/deleteLiftModel [delete]
 export const deleteLiftModel = (data) => {
     return service({
         url: "/liftModel/deleteLiftModel",
         method: 'delete',
         data
     })
 }

// @Tags LiftModel
// @Summary 删除LiftModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LiftModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftModel/deleteLiftModel [delete]
 export const deleteLiftModelByIds = (data) => {
     return service({
         url: "/liftModel/deleteLiftModelByIds",
         method: 'delete',
         data
     })
 }

// @Tags LiftModel
// @Summary 更新LiftModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftModel true "更新LiftModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /liftModel/updateLiftModel [put]
 export const updateLiftModel = (data) => {
     return service({
         url: "/liftModel/updateLiftModel",
         method: 'put',
         data
     })
 }


// @Tags LiftModel
// @Summary 用id查询LiftModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftModel true "用id查询LiftModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /liftModel/findLiftModel [get]
 export const findLiftModel = (params) => {
     return service({
         url: "/liftModel/findLiftModel",
         method: 'get',
         params
     })
 }


// @Tags LiftModel
// @Summary 分页获取LiftModel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取LiftModel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftModel/getLiftModelList [get]
 export const getLiftModelList = (params) => {
     return service({
         url: "/liftModel/getLiftModelList",
         method: 'get',
         params
     })
 }