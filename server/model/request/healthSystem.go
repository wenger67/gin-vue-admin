package request

import "panta/model"

type HealthSystemSearch struct{
    model.HealthSystem
    PageInfo
}