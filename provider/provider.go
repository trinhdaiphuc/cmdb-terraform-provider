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
	Hostname   types.String `tfsdk:"hostname"`
	Protocol   types.String `tfsdk:"protocol"`
	Port       types.String `tfsdk:"port"`
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
			"hostname": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
			"protocol": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
		},
	}, nil
}

func (p *provider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, []*tfprotov6.Diagnostic) {
	return map[string]tfsdk.ResourceType{
		"hashicups_order": resourceConfigType{},
	}, nil
}

func (p *provider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, []*tfprotov6.Diagnostic) {
	return map[string]tfsdk.DataSourceType{
	}, nil
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	// Retrieve provider data from configuration
	var (
		config                               providerData
		hostname, apiVersion, protocol, port string
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

	// User must provide a hostname to the provider
	if config.Hostname.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityWarning,
			Summary:  "Unable to create client",
			Detail:   "Cannot use unknown value as hostname",
		})
		return
	}

	if config.Hostname.Null {
		hostname = os.Getenv("CMDB_HOSTNAME")
	} else {
		hostname = config.Hostname.Value
	}

	if hostname == "" {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			// Error vs warning - empty value must stop execution
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Unable to find hostname",
			Detail:   "Hostname cannot be an empty string",
		})
	}

	// User must provide a protocol to the provider
	if config.Protocol.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityWarning,
			Summary:  "Unable to create client",
			Detail:   "Cannot use unknown value as protocol",
		})
		return
	}

	if config.Protocol.Null {
		apiVersion = os.Getenv("CMDB_PROTOCOL")
	} else {
		apiVersion = config.Protocol.Value
	}

	if protocol == "" {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			// Error vs warning - empty value must stop execution
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Unable to find protocol",
			Detail:   "Protocol cannot be an empty string",
		})
	}

	// User must provide a port to the provider
	if config.Port.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityWarning,
			Summary:  "Unable to create client",
			Detail:   "Cannot use unknown value as port",
		})
		return
	}

	if config.Port.Null {
		hostname = os.Getenv("CMDB_PORT")
	} else {
		hostname = config.Port.Value
	}

	if port == "" {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			// Error vs warning - empty value must stop execution
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Unable to find port",
			Detail:   "Port cannot be an empty string",
		})
	}

	cli := NewClient(port, protocol, hostname, apiVersion)
	p.client = cli
	p.configured = true
}

