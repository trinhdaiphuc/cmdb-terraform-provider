package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"terraform-provider-cmdb/provider"
)

func main() {
	tfsdk.Serve(context.Background(), provider.New, tfsdk.ServeOpts{
		Name: "cmdb",
	})
}
