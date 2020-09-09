package request

import "gin-vue-admin/model"

type HealthSystemSearch struct{
    model.HealthSystem
    PageInfo
}