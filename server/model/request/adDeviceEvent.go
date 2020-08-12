package request

import "gin-vue-admin/model"

type AdDeviceEventSearch struct{
    model.AdDeviceEvent
    PageInfo
}