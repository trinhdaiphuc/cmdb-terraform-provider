package controller

import (
	"github.com/trinhdaiphuc/terraform-provider-cmdb/cmdb/model"
	"github.com/whatvn/denny"
)

func DeleteConfig(ctx *denny.Context) {
	name, ok := ctx.GetQuery("name")
	if !ok {
		ctx.Status(400)
		return
	}

	model.DeleteAllocatedConfig(name)
	ctx.Status(200)
}
