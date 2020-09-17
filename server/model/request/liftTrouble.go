package request

import "panta/model"

type LiftTroubleSearch struct{
    model.LiftTrouble
    PageInfo
}