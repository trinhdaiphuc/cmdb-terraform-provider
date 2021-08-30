package controller

import (
	"github.com/whatvn/denny"
	"terraform-provider-cmdb/cmdb/model"
)

func GetConfig(ctx *denny.Context) {
	name, ok := ctx.GetQuery("name")
	if !ok {
		ctx.Status(400)
		return
	}
	config := model.GetAllocatedConfig(name)
	ctx.JSON(200, config)
}
