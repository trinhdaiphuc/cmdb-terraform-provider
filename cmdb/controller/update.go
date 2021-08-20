package controller

import "github.com/whatvn/denny"

func UpdateName(ctx *denny.Context) {
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
	newName := putAllocatedName(name, value)
	ctx.JSON(200, newName)
}
