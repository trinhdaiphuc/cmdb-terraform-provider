package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHistory() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHistoryRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configVersions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"createdAt": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updatedAt": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceHistoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cli := m.(*Client)

	// // Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	his := d.Get("name").(string)
	resp, err := cli.GetHistory(his)
	if err != nil {
		diag.FromErr(err)
	}

	histories := make([]map[string]interface{}, 0)

	for _, v := range resp.VersionConfigs {
		config := make(map[string]interface{})

		config["name"] = v.Name
		config["value"] = v.Value
		config["createdAt"] = v.CreatedAt
		config["updatedAt"] = v.UpdateAt

		histories = append(histories, config)
	}

	if err := d.Set("history", histories); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(his)

	return diags
}
