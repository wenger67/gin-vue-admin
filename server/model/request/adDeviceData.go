package request

import "gin-vue-admin/model"

type AdDeviceDataSearch struct{
    model.AdDeviceData
    PageInfo
}