package controller

import (
	"github.com/whatvn/denny"
	"terraform-provider-cmdb/cmdb/model"
)

func GetHistory(ctx *denny.Context) {
	name, ok := ctx.GetQuery("name")
	if !ok {
		ctx.Status(400)
		return
	}
	history := model.GetHistory(name)
	ctx.JSON(200, history)
}
