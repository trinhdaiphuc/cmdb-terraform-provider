package prodiver

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

func resourceName() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNameCreate,
		ReadContext:   resourceNameRead,
		UpdateContext: resourceNameUpdate,
		DeleteContext: resourceNameDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"config": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
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

func resourceNameCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		diags diag.Diagnostics
		cli   = m.(*Client)
		item  = d.Get("config").(map[string]interface{})
	)

	name := item["name"].(string)
	value := item["value"].(string)
	_, err := cli.CreateName(name, value)
	if err != nil {
		diag.FromErr(err)
	}
	d.SetId(name)
	return diags
}

func resourceNameRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		diags diag.Diagnostics
		cli   = m.(*Client)
	)

	name := d.Id()
	resp, err := cli.GetName(name)
	if err != nil {
		diag.FromErr(err)
	}
	d.Set("config", flattenOrderItems(resp))
	return diags
}

func resourceNameUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		cli  = m.(*Client)
		item = d.Get("config").(map[string]interface{})
	)

	name := d.Id()

	if d.HasChange("config") {
		dname := item["name"].(string)
		if name != dname {
			return diag.Errorf("Invalid name")
		}
		value := item["value"].(string)
		_, err := cli.UpdateName(name, value)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC3339))
	}

	return resourceNameRead(ctx, d, m)
}

func resourceNameDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		diags diag.Diagnostics
		cli   = m.(*Client)
	)

	name := d.Id()
	err := cli.DeleteName(name)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func flattenOrderItems(config *Config) interface{} {
	if config != nil {
		cf := make(map[string]interface{})
		cf["name"] = config.Name
		cf["value"] = config.Value
		cf["createdAt"] = config.CreatedAt
		cf["updatedAt"] = config.UpdateAt
		return cf
	}

	return nil
}
