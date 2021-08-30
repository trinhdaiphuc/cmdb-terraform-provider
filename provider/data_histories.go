package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

type dataSourceHistoriesType struct{}

func (r dataSourceHistoriesType) GetSchema(_ context.Context) (schema.Schema, []*tfprotov6.Diagnostic) {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"history": {
				// When Computed is true, the provider will set value --
				// the user cannot define the value
				Computed: true,
				Attributes: schema.SingleNestedAttributes(map[string]schema.Attribute{
					"name": {
						Type:     types.StringType,
						Required: true,
					},
					"configVersions": {
						Attributes: schema.ListNestedAttributes(map[string]schema.Attribute{
							"name": {
								Type:     types.StringType,
								Computed: true,
							},
							"value": {
								Type:     types.StringType,
								Computed: true,
							},
							"createdAt": {
								Type:     types.StringType,
								Computed: true,
							},
							"updatedAt": {
								Type:     types.StringType,
								Computed: true,
							},
						}, schema.ListNestedAttributesOptions{}),
					},
				}),
			},
		},
	}, nil
}

func (r dataSourceHistoriesType) NewDataSource(ctx context.Context, p tfsdk.Provider) (tfsdk.DataSource, []*tfprotov6.Diagnostic) {
	return dataSourceHistories{
		p: *(p.(*provider)),
	}, nil
}

type dataSourceHistories struct {
	p provider
}

func (r dataSourceHistories) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	// Declare struct that this function will set to this data source's state
	var his History

	err := req.Config.Get(ctx, &his)

	histories, err := r.p.client.GetHistory(his.Name)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error retrieving histories",
		})
		return
	}

	// Sample debug message
	// To view this message, set the TF_LOG environment variable to DEBUG
	// 		`export TF_LOG=DEBUG`
	// To hide debug message, unset the environment variable
	// 		`unset TF_LOG`
	fmt.Fprintf(stderr, "[DEBUG]-Resource State:%+v", histories)

	// Set state
	err = resp.State.Set(ctx, &histories)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error reading histories",
			Detail:   fmt.Sprintf("An unexpected error was encountered while reading the datasource_history: %+v", err.Error()),
		})
		return
	}
}
