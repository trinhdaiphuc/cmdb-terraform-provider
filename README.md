# Terraform provider

## Table of content

1. [Introduction](#introduction)
2. [What Providers Do](#what-providers-do)
3. [Where Providers Come From](#where-providers-come-from)
4. [How to Use Providers](#how-to-use-providers)

## Introduction

Terraform relies on plugins called "providers" to interact with cloud providers, SaaS providers, and other APIs.

Terraform configurations must declare which providers they require so that Terraform can install and use them.
Additionally, some providers require configuration (like endpoint URLs or cloud regions) before they can be used.

- Resources: are the most important element in the Terraform language. Each resource block describes one or more
  infrastructure objects, such as virtual networks, compute instances, or higher-level components such as DNS records.

- Data sources: allow Terraform use information defined outside of Terraform, defined by another separate Terraform
  configuration, or modified by functions.

## What Providers Do

Each provider adds a set of resource types and/or data sources that Terraform can manage.

Every resource type is implemented by a provider; without providers, Terraform can't manage any kind of infrastructure.

Most providers configure a specific infrastructure platform (either cloud or self-hosted). Providers can also offer
local utilities for tasks like generating random numbers for unique resource names.

## Where Providers Come From

Providers are distributed separately from Terraform itself, and each provider has its own release cadence and version
numbers.

The Terraform Registry is the main directory of publicly available Terraform providers, and hosts providers for most
major infrastructure platforms.

## How to Use Providers

To use resources from a given provider, you need to include some information about it in your configuration. See the
following pages for details:

[Provider Requirements](#provider-requirements) documents how to declare providers so Terraform can install them.

[Provider Configuration](#provider-configuration) documents how to configure settings for providers.

[Dependency Lock File](#dependency-lock-file) documents an additional HCL file that can be included with a
configuration, which tells Terraform to always use a specific set of provider versions.

### Provider Requirements

Terraform relies on plugins called "providers" to interact with remote systems.

#### Requiring Providers

Each Terraform module must declare which providers it requires, so that Terraform can install and use them. Provider
requirements are declared in a `required_providers` block.

A provider requirement consists of a local name, a source location, and a version constraint:

```terraform
terraform {
  required_providers {
    cmdb = {
      version = "0.3"
      source = "zalopay.com.vn/top/cmdb"
    }
  }
}
```

The `required_providers` block must be nested inside the top-level terraform block (which can also contain other
settings)
.

Each argument in the `required_providers` block enables one provider. The key determines the provider's local name (its
unique identifier within this module), and the value is an object with the following elements:

- source - the global source address for the provider you intend to use, such as hashicorp/aws.

- version - a version constraint specifying which subset of available provider versions the module is compatible with.

#### Names and Addresses

Each provider has two identifiers:

- A unique source address, which is only used when requiring a provider.
- A local name, which is used everywhere else in a Terraform module.

##### Local Names

Local names are module-specific, and are assigned when requiring a provider. Local names must be unique per-module.

Outside of the `required_providers` block, Terraform configurations always refer to providers by their local names. For
example, the following configuration declares `cmdb` as the local name for `zalopay.com.vn/top/cmdb`, then uses that
local name when configuring the provider:

```terraform
terraform {
  required_providers {
    cmdb = {
      version = "0.3"
      source = "zalopay.com.vn/top/cmdb"
    }
  }
}

provider "cmdb" {
  # ...
}
```

##### Source Addresses

A provider's source address is its global identifier. It also specifies the primary location where Terraform can
download it.

Source addresses consist of three parts delimited by slashes (/), as follows:

`[<HOSTNAME>/]<NAMESPACE>/<TYPE>`

- Hostname (optional): The hostname of the Terraform registry that distributes the provider. If omitted, this defaults
  to registry.terraform.io, the hostname of the public Terraform Registry.

- Namespace: An organizational namespace within the specified registry. For the public Terraform Registry and for
  Terraform Cloud's private registry, this represents the organization that publishes the provider. This field may have
  other meanings for other registry hosts.

- Type: A short name for the platform or system the provider manages. Must be unique within a particular namespace on a
  particular registry host.

#### Version Constraints

Each provider plugin has its own set of available versions, allowing the functionality of the provider to evolve over
time. Each provider dependency you declare should have a version constraint given in the version argument so Terraform
can select a single version per provider that all modules are compatible with.

Each module should at least declare the minimum provider version it is known to work with, using the >= version
constraint syntax:

```terraform
terraform {
  required_providers {
    mycloud = {
      source = "hashicorp/aws"
      version = ">= 1.0"
    }
  }
}
```

The ~> operator is a convenient shorthand for allowing only patch releases within a specific minor release:

```terraform
terraform {
  required_providers {
    mycloud = {
      source = "hashicorp/aws"
      version = "~> 1.0.4"
    }
  }
}
```

### Provider Configuration

Provider configurations belong in the root module of a Terraform configuration. A provider configuration is created
using a `provider` block:

```terraform
provider "google" {
  project = "acme-app"
  region = "us-central1"
}
```

The name given in the block header (`"google"` in this example) is the local name of the provider to configure. This
provider should already be included in a `required_providers` block.

The body of the block (between `{` and `}`) contains configuration arguments for the provider. Most arguments in this
section are defined by the provider itself; in this example both project and region are specific to the google provider.

### Dependency Lock File

A Terraform configuration may refer to two different kinds of external dependency that come from outside of its own
codebase:

- Providers, which are plugins for Terraform that extend it with support for interacting with various external systems.

- Modules, which allow splitting out groups of Terraform configuration constructs (written in the Terraform language)
  into reusable abstractions.

Both of these dependency types can be published and updated independently from Terraform itself and from the
configurations that depend on them. For that reason, Terraform must determine which versions of those dependencies are
potentially compatible with the current configuration and which versions are currently selected for use.

## Terraform Custom Provider

Interact with APIs using Terraform providers. Use a provider as a bridge between Terraform and a target API. Then,
extend Terraform by developing a custom Terraform provider.

Later in the track, you will re-create the Cmdb provider based on
the [Terraform Plugin SDK v2](https://github.com/hashicorp/terraform-plugin-sdk).

### Terraform plugins

Terraform is comprised of Terraform Core and Terraform Plugins.

![image](./images/core-plugins-api.png)

- Terraform Core reads the configuration and builds the resource dependency graph.
- Terraform Plugins (providers and provisioners) bridge Terraform Core and their respective target APIs. Terraform
  provider plugins implement resources via basic CRUD (create, read, update, and delete) APIs to communicate with third
  party services.

Upon `terraform plan` or `terraform apply`, Terraform Core asks the Terraform provider to perform an action via a RPC
interface. The provider attempts to fulfill the request by invoking a CRUD operation against the target API's client
library. This process enforces a clear separation of concerns. Providers are able to serve as an abstraction of a client
library.

### Setup and Implement Read

#### Prerequisites

- A Golang 1.15+ installed and configured.
- The Terraform 0.14+ CLI installed locally.
- Docker and Docker Compose to run an instance of Cmdb locally.

#### Set up your development environment

- Run docker-compose up to spin up a local instance of Cmdb on port :8080.

```shell
docker-compose up
```

- Verify that Cmdb is running by sending a request to its health check endpoint.

```shell
curl localhost:8080/health
```

##### Explore main.go file

Open `main.go` in the root of the repository. The contents of the main function consume the Plugin SDK's plugin library
which facilitates the RPC communication between Terraform Core and the plugin.

```go
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/trinhdaiphuc/terraform-provider-cmdb/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return provider.Provider()
		},
	})
}
```

_Notice the ProviderFunc returns a *provider.Provider from the terraform-provider-cmdb/provider package._

##### Explore provider schema

The provider/provider.go file currently defines an cmdb provider.

```go
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
```

The helper/schema library is part of Terraform Core. It abstracts many of the complexities and ensures consistency
between providers. The *schema.Provider type can accept:

- The resources it supports (ResourcesMap and DataSourcesMap)
- Configuration keys (properties in *schema.Schema{})
- Any callbacks to configure (ConfigureContextFunc). This function retrieves the `api_version` and `host` from the
  provider schema to connect to create a client connect to cmdb and configure your provider.

#### Implement Create

- Define `config` resource

To create a Cmdb config, you would send a POST request to the /api/v1/configs endpoint with a config item.

```shell
curl -X POST localhost:8080/api/v1/configs -d "name=db.host&value=localhost"

{"name":"db.host","value":"localhost","createdAt":"2021-08-31T16:26:45+07:00","updatedAt":"2021-08-31T16:26:45+07:00"}
```

```go
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfig() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConfigCreate,
		ReadContext:   resourceConfigRead,
		UpdateContext: resourceConfigUpdate,
		DeleteContext: resourceConfigDelete,
		Schema:        map[string]*schema.Schema{},
	}
}

func resourceConfigCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		diags diag.Diagnostics
		// ...
	)

	// ...

	return diags
}

func resourceConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		diags diag.Diagnostics
		// ...
	)
	// ...
	return diags
}

func resourceConfigUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// ...

	return resourceConfigRead(ctx, d, m)
}

func resourceConfigDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		diags diag.Diagnostics
		// ...
	)

	// ...

	return diags
}
```

Look at file provider/resource.go in the provider directory. As a general convention, Terraform providers put each
resource in their own file, named after the resource, prefixed with `resource_`.

- Define `config` schema

Replace the line `Schema: map[string]*schema.Schema{}`, in your resourceOrder function with the following schema. The
order resource schema should resemble the request body.

```shell
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
```


