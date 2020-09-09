package request

import "gin-vue-admin/model"

type MessageSearch struct{
    model.Message
    PageInfo
}