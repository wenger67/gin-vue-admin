package request

import "panta/model"

type MessageSearch struct{
    model.Message
    PageInfo
}