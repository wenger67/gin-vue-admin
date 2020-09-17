package request

import (
    "panta/model"
    "time"
)

type AdDeviceEventSearch struct{
    model.AdDeviceEvent
    PageInfo
}

type DeviceEventCreate struct {
    model.AdDeviceEvent
    Storage string `json:"storage" form:"storage"`
}

type DeviceEventContent struct {
    Items []DeviceEventContentItem `json:"items" form:"items"`
    Videos []string `json:"videos" form:"videos"`
}

type DeviceEventContentItem struct {
    Brief string `json:"brief" form:"brief"`
    Images []string `json:"images" form:"images"`
    CreatedAt time.Time `json:"createdAt" form:"createdAt"`
}