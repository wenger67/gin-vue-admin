package request

import "gin-vue-admin/model"

type AdDeviceConfigSearch struct{
    model.AdDeviceConfig
    PageInfo
}