package request

import "panta/model"

type CompanySearch struct{
    model.Company
    PageInfo
    ID int
}