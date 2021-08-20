package prodiver

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
			"hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CMDB_HOST", "localhost"),
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CMDB_PROTOCOL", "http"),
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CMDB_PORT", 8080),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cmdb_name": resourceName(),
		},
	}
}

// providerConfigure parses the config into the Terraform provider meta object
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var (
		diags      diag.Diagnostics
		apiVersion = d.Get("api_version").(string)
		protocol   = d.Get("protocol").(string)
		hostname   = d.Get("hostname").(string)
		port       = d.Get("port").(int)
	)
	cli := NewClient(port, protocol, hostname, apiVersion)
	return cli, diags
}
