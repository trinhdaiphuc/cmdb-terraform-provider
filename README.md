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

## Terraform Plugin Framework

