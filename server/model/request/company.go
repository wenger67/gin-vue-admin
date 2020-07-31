package request

import "gin-vue-admin/model"

type CompanySearch struct{
    model.Company
    PageInfo
    ID int
}