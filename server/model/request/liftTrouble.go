package request

import "gin-vue-admin/model"

type LiftTroubleSearch struct{
    model.LiftTrouble
    PageInfo
}