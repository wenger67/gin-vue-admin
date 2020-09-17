package request

import "panta/model"

type HealthChangeSearch struct{
    model.HealthChange
    PageInfo
}