package controller

import "github.com/whatvn/denny"

func getAllocatedName(name string) *AllocatedName {
	if v, ok := names[name]; ok {
		return v
	}
	return nil
}

func GetName(ctx *denny.Context) {
	name, ok := ctx.GetQuery("name")
	if !ok {
		ctx.Status(400)
		return
	}
	getName := getAllocatedName(name)
	ctx.JSON(200, getName)
}
