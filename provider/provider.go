package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: providerConfigure,
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CMDB_API_VERSION", "v1"),
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CMDB_HOST", "http://localhost:8080"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cmdb_config": resourceConfig(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cmdb_config": dataSourceHistory(),
		},
	}
}

// providerConfigure parses the config into the Terraform provider meta object
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var (
		diags      diag.Diagnostics
		apiVersion = d.Get("api_version").(string)
		host       = d.Get("host").(string)
	)
	cli := NewClient(host, apiVersion)
	return cli, diags
}
