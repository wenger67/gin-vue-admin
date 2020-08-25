import service from '@/utils/request'

// @Tags LiftRecord
// @Summary 创建LiftRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftRecord true "创建LiftRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftRecord/createLiftRecord [post]
export const createLiftRecord = (data) => {
     return service({
         url: "/liftRecord/createLiftRecord",
         method: 'post',
         data
     })
 }

 export const fillLiftRecord = (data) => {
  return service({
      url: "/liftRecord/fillLiftRecord",
      method: 'post',
      data
  })
}

export const reviewLiftRecord = (data) => {
  return service({
      url: "/liftRecord/reviewLiftRecord",
      method: 'post',
      data
  })
}


// @Tags LiftRecord
// @Summary 删除LiftRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftRecord true "删除LiftRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftRecord/deleteLiftRecord [delete]
 export const deleteLiftRecord = (data) => {
     return service({
         url: "/liftRecord/deleteLiftRecord",
         method: 'delete',
         data
     })
 }

// @Tags LiftRecord
// @Summary 删除LiftRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LiftRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liftRecord/deleteLiftRecord [delete]
 export const deleteLiftRecordByIds = (data) => {
     return service({
         url: "/liftRecord/deleteLiftRecordByIds",
         method: 'delete',
         data
     })
 }

// @Tags LiftRecord
// @Summary 更新LiftRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftRecord true "更新LiftRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /liftRecord/updateLiftRecord [put]
 export const updateLiftRecord = (data) => {
     return service({
         url: "/liftRecord/updateLiftRecord",
         method: 'put',
         data
     })
 }


// @Tags LiftRecord
// @Summary 用id查询LiftRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LiftRecord true "用id查询LiftRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /liftRecord/findLiftRecord [get]
 export const findLiftRecord = (params) => {
     return service({
         url: "/liftRecord/findLiftRecord",
         method: 'get',
         params
     })
 }


// @Tags LiftRecord
// @Summary 分页获取LiftRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取LiftRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liftRecord/getLiftRecordList [get]
 export const getLiftRecordList = (params) => {
     return service({
         url: "/liftRecord/getLiftRecordList",
         method: 'get',
         params
     })
 }