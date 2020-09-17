package request

import "panta/model"

type LiftChangeSearch struct{
    model.LiftChange
    PageInfo
}