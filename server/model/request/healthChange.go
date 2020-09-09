package request

import "gin-vue-admin/model"

type HealthChangeSearch struct{
    model.HealthChange
    PageInfo
}