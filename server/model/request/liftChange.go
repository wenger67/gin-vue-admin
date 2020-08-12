package request

import "gin-vue-admin/model"

type LiftChangeSearch struct{
    model.LiftChange
    PageInfo
}