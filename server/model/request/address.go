package request

import "gin-vue-admin/model"

type AddressSearch struct{
    model.Address
    PageInfo
}