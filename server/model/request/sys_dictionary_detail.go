package request

import "panta/model"

type SysDictionaryDetailSearch struct{
    model.SysDictionaryDetail
    PageInfo
}