package request

import "panta/model"

type LiftModelSearch struct{
    model.LiftModel
    PageInfo
}