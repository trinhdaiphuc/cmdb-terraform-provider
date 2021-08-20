package controller

import (
	"github.com/whatvn/denny"
	"time"
)

// AllocatedName holds metadata of a name allocated to a deployed resource
// within one of our theoretical deployment environments
type AllocatedName struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

var (
	names = make(map[string]*AllocatedName)
)

func putAllocatedName(name, value string) *AllocatedName {
	v, ok := names[name]
	if ok {
		v.UpdatedAt = time.Now().String()
		names[name] = v
		return v
	}
	allocatedName := &AllocatedName{
		Name:      name,
		Value:     value,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	names[name] = allocatedName
	return allocatedName
}

func CreateName(ctx *denny.Context) {
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
