package request

import "panta/model"

type LiftTroubleSearch struct{
    model.LiftTrouble
    PageInfo
    RankInfo
}

type LiftTroubleUpdate struct {
    model.LiftTrouble
}