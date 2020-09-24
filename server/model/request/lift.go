package request

import "panta/model"

type LiftSearch struct{
    model.Lift
    PageInfo
    RankInfo
}