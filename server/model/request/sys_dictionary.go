package request

import "panta/model"

type SysDictionarySearch struct{
    model.SysDictionary
    PageInfo
}