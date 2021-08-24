# Terraform provider

## Table of content

1. [Performance test's infrastructure](#performance-tests-infrastructure)
2. [Performance test's infrastructure](#performance-tests-infrastructure)
3. [Performance test's technical](#performance-tests-technical)

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

Provider Requirements documents how to declare providers so Terraform can install them.

Provider Configuration documents how to configure settings for providers.

Dependency Lock File documents an additional HCL file that can be included with a configuration, which tells Terraform
to always use a specific set of provider versions.