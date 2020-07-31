import service from '@/utils/request'

// @Tags Region
// @Summary 创建Region
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Region true "创建Region"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /region/createRegion [post]
export const createRegion = (data) => {
     return service({
         url: "/region/createRegion",
         method: 'post',
         data
     })
 }


// @Tags Region
// @Summary 删除Region
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Region true "删除Region"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /region/deleteRegion [delete]
 export const deleteRegion = (data) => {
     return service({
         url: "/region/deleteRegion",
         method: 'delete',
         data
     })
 }

// @Tags Region
// @Summary 删除Region
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Region"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /region/deleteRegion [delete]
 export const deleteRegionByIds = (data) => {
     return service({
         url: "/region/deleteRegionByIds",
         method: 'delete',
         data
     })
 }

// @Tags Region
// @Summary 更新Region
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Region true "更新Region"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /region/updateRegion [put]
 export const updateRegion = (data) => {
     return service({
         url: "/region/updateRegion",
         method: 'put',
         data
     })
 }


// @Tags Region
// @Summary 用id查询Region
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Region true "用id查询Region"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /region/findRegion [get]
 export const findRegion = (params) => {
     return service({
         url: "/region/findRegion",
         method: 'get',
         params
     })
 }


// @Tags Region
// @Summary 分页获取Region列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Region列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /region/getRegionList [get]
 export const getRegionList = (params) => {
     return service({
         url: "/region/getRegionList",
         method: 'get',
         params
     })
 }