package request

import (
	"panta/model"
)

type AdDeviceSearch struct{
    model.Device
    PageInfo
}
