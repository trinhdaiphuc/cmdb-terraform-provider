package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"os"
)

type provider struct {
	configured bool
	client     *Client
}

// Provider schema struct
type providerData struct {
	ApiVersion types.String `tfsdk:"api_version"`
	Host       types.String `tfsdk:"host"`
}

var stderr = os.Stderr

func New() tfsdk.Provider {
	return &provider{}
}

// GetSchema
func (p *provider) GetSchema(_ context.Context) (schema.Schema, []*tfprotov6.Diagnostic) {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_version": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
			"host": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
		},
	}, nil
}

func (p *provider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, []*tfprotov6.Diagnostic) {
	return map[string]tfsdk.ResourceType{
		"cmdb_config": resourceConfigType{},
	}, nil
}

func (p *provider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, []*tfprotov6.Diagnostic) {
	return map[string]tfsdk.DataSourceType{
		"cmdb_config": dataSourceHistoriesType{},
	}, nil
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	// Retrieve provider data from configuration
	var (
		config           providerData
		host, apiVersion string
	)
	err := req.Config.Get(ctx, &config)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error parsing configuration",
			Detail:   "Error parsing the configuration, this is an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
		})
		return
	}

	// User must provide an api_version to the provider
	if config.ApiVersion.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityWarning,
			Summary:  "Unable to create client",
			Detail:   "Cannot use unknown value as api version",
		})
		return
	}

	if config.ApiVersion.Null {
		apiVersion = os.Getenv("CMDB_API_VERSION")
	} else {
		apiVersion = config.ApiVersion.Value
	}

	if apiVersion == "" {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			// Error vs warning - empty value must stop execution
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Unable to find api_version",
			Detail:   "Api version cannot be an empty string",
		})
	}

	// User must provide a host to the provider
	if config.Host.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityWarning,
			Summary:  "Unable to create client",
			Detail:   "Cannot use unknown value as host",
		})
		return
	}

	if config.Host.Null {
		host = os.Getenv("CMDB_HOST")
	} else {
		host = config.Host.Value
	}

	if host == "" {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			// Error vs warning - empty value must stop execution
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Unable to find host",
			Detail:   "Host cannot be an empty string",
		})
	}

	cli := NewClient(host, apiVersion)
	p.client = cli
	p.configured = true
}
