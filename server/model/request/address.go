package request

import "panta/model"

type AddressSearch struct{
    model.Address
    PageInfo
}