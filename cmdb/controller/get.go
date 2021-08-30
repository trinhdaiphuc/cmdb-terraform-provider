package controller

import (
	"github.com/trinhdaiphuc/terraform-provider-cmdb/cmdb/model"
	"github.com/whatvn/denny"
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
