package request

import "gin-vue-admin/model"

type LiftModelSearch struct{
    model.LiftModel
    PageInfo
}