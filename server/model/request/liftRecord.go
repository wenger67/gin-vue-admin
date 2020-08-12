package request

import "gin-vue-admin/model"

type LiftRecordSearch struct{
    model.LiftRecord
    PageInfo
}