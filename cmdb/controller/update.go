package controller

import (
	"github.com/whatvn/denny"
	"terraform-provider-cmdb/cmdb/model"
)

func UpdateConfig(ctx *denny.Context) {
	name, ok := ctx.GetPostForm("name")
	if !ok {
		ctx.Status(400)
		return
	}
	value, ok := ctx.GetPostForm("value")
	if !ok {
		ctx.Status(400)
		return
	}
	config := model.PutAllocatedConfig(name, value)
	ctx.JSON(200, config)
}
