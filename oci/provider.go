// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

// Package oci exposes the supported Go construction boundary for embedding
// Terraform Provider OCI in another Go process.
package oci

import (
	frameworkprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	internalprovider "github.com/oracle/terraform-provider-oci/internal/provider"
)

// Provider returns a fresh SDKv2 provider instance configured for safe
// in-process embedding.
func Provider() *schema.Provider {
	return internalprovider.NewSDKv2ProviderForInProcess()
}

// New returns a fresh Plugin Framework provider instance configured for safe
// in-process embedding.
func New() frameworkprovider.Provider {
	return internalprovider.NewFrameworkProviderForInProcess()
}
