package request

import "gin-vue-admin/model"

type RegionSearch struct{
    model.Region
    PageInfo
    GroupKey string `json:"groupKey" form:"groupKey"`
}