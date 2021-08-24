package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

// ConfigResource -
type ConfigResource struct {
	ID          types.String `tfsdk:"id"`
	Config      Config       `tfsdk:"item"`
	LastUpdated types.String `tfsdk:"last_updated"`
}

type Config struct {
	Name      string `json:"name" tfsdk:"name"`
	Value     string `json:"value" tfsdk:"value"`
	CreatedAt string `json:"createdAt" tfsdk:"created_at"`
	UpdateAt  string `json:"updateAt" tfsdk:"update_at"`
}

type RequestParams struct {
	Name  string `http:"name,form"`
	Value string `http:"value,form"`
}
