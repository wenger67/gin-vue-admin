package request

import "panta/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}