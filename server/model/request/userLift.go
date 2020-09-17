package request

import "panta/model"

type UserLiftSearch struct{
    model.UserLift
    PageInfo
}