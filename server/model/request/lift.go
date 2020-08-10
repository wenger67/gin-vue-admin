package request

import "gin-vue-admin/model"

type LiftSearch struct{
    model.Lift
    PageInfo
}