package controller

import "github.com/whatvn/denny"

func deleteAllocatedName(name string) {
	if _, ok := names[name]; ok {
		delete(names, name)
	}
}

func DeleteName(ctx *denny.Context) {
	name, ok := ctx.GetQuery("name")
	if !ok {
		ctx.Status(400)
		return
	}

	deleteAllocatedName(name)
	ctx.Status(200)
}
