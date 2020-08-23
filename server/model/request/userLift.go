package request

import "gin-vue-admin/model"

type UserLiftSearch struct{
    model.UserLift
    PageInfo
}