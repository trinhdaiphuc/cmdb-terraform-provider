package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"time"
)

type resourceConfigType struct{}

type resourceConfig struct {
	p provider
}

// NewResource instance
func (r resourceConfigType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, []*tfprotov6.Diagnostic) {
	return resourceConfig{
		p: *(p.(*provider)),
	}, nil
}

func (r resourceConfigType) GetSchema(_ context.Context) (schema.Schema, []*tfprotov6.Diagnostic) {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": {
				Type: types.StringType,
				// When Computed is true, the provider will set value --
				// the user cannot define the value
				Computed: true,
			},
			"last_updated": {
				Type:     types.StringType,
				Computed: true,
			},
			"config": {
				Required: true,
				Attributes: schema.SingleNestedAttributes(map[string]schema.Attribute{
					"name": {
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						Type:     types.StringType,
						Required: true,
					},
					"createdAt": {
						Type:     types.StringType,
						Computed: true,
					},
					"updatedAt": {
						Type:     types.StringType,
						Computed: true,
					},
				}),
			},
		},
	}, nil
}

func (r resourceConfig) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Provider not configured",
			Detail:   "The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		})
		return
	}

	// Retrieve values from plan
	var plan ConfigResource
	err := req.Plan.Get(ctx, &plan)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error reading plan",
			Detail:   "An unexpected error was encountered while reading the plan: " + err.Error(),
		})
		return
	}

	// Create new config
	config, err := r.p.client.CreateConfig(plan.Config)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error creating config",
			Detail:   "Could not create config, unexpected error: " + err.Error(),
		})
		return
	}

	// Generate resource state struct
	var result = ConfigResource{
		ID:          types.String{Value: config.Name},
		Config:      config,
		LastUpdated: types.String{Value: time.Now().Format(time.RFC3339)},
	}

	err = resp.State.Set(ctx, result)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error setting state",
			Detail:   "Could not set state, unexpected error: " + err.Error(),
		})
		return
	}
}

func (r resourceConfig) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// Get current state
	var state ConfigResource
	err := req.State.Get(ctx, &state)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error reading state",
			Detail:   "An unexpected error was encountered while reading the state: " + err.Error(),
		})
		return
	}

	// Get config from API and then update what is in state from what the API returns
	configName := state.ID.Value

	// Get order current value
	config, err := r.p.client.GetConfig(configName)

	// Map response body to resource schema attribute
	state.Config = config

	// Set state
	err = resp.State.Set(ctx, &state)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error setting state",
			Detail:   "Unexpected error encountered trying to set new state: " + err.Error(),
		})
		return
	}
}

func (r resourceConfig) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Get plan values
	var plan ConfigResource
	err := req.Plan.Get(ctx, &plan)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error reading plan",
			Detail:   "An unexpected error was encountered while reading the plan: " + err.Error(),
		})
		return
	}

	// Get current state
	var state ConfigResource
	err = req.State.Get(ctx, &state)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error reading prior state",
			Detail:   "An unexpected error was encountered while reading the prior state: " + err.Error(),
		})
		return
	}

	// Get config from API and then update what is in state from what the API returns
	configName := state.ID.Value
	plan.Config.Name = configName

	// Update order by calling API
	config, err := r.p.client.UpdateConfig(plan.Config)

	// Generate resource state struct
	var result = ConfigResource{
		ID:          types.String{Value: configName},
		Config:      config,
		LastUpdated: types.String{Value: string(time.Now().Format(time.RFC850))},
	}

	// Set state
	err = resp.State.Set(ctx, result)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error setting state",
			Detail:   "Could not set state, unexpected error: " + err.Error(),
		})
		return
	}
}

func (r resourceConfig) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state ConfigResource
	err := req.State.Get(ctx, &state)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error reading configuration",
			Detail:   "An unexpected error was encountered while reading the configuration: " + err.Error(),
		})
		return
	}

	// Get order ID from state
	configName := state.ID.Value

	// Delete order by calling API
	err = r.p.client.DeleteConfig(configName)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error deleting order",
			Detail:   "Could not delete configName " + configName + ": " + err.Error(),
		})
		return
	}

	// Remove resource from state
	resp.State.RemoveResource(ctx)
}
