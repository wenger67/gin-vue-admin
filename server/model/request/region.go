package request

import "panta/model"

type RegionSearch struct{
    model.Region
    PageInfo
    GroupKey string `json:"groupKey" form:"groupKey"`
}