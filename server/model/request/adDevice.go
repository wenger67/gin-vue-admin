package request

import (
	"gin-vue-admin/model"
)

type AdDeviceSearch struct{
    model.AdDevice
    PageInfo
}
