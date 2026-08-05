// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

func TestProviderConstructorsReturnFreshInstances(t *testing.T) {
	first := Provider()
	second := Provider()
	if first == second {
		t.Fatal("Provider returned a shared SDKv2 provider")
	}
	firstEmbedded := NewSDKv2ProviderForInProcess()
	secondEmbedded := NewSDKv2ProviderForInProcess()
	if firstEmbedded == secondEmbedded {
		t.Fatal("NewSDKv2ProviderForInProcess returned a shared SDKv2 provider")
	}
}

func TestValidateInProcessProviderConfig(t *testing.T) {
	defaults := schema.TestResourceDataRaw(t, SchemaMap(), map[string]interface{}{})
	if err := validateInProcessProviderConfig(defaults); err != nil {
		t.Fatalf("default provider configuration was rejected: %v", err)
	}

	configured := schema.TestResourceDataRaw(t, SchemaMap(), map[string]interface{}{
		globalvar.DisableAutoRetriesAttrName: true,
	})
	if err := validateInProcessProviderConfig(configured); err == nil {
		t.Fatal("process-global retry option was accepted for in-process use")
	}
}

func TestInternalAndEmbeddedSDKv2SchemasRemainCompatible(t *testing.T) {
	cli := Provider()
	embedded := NewSDKv2ProviderForInProcess()
	if len(cli.Schema) != len(embedded.Schema) {
		t.Fatalf("provider schema size differs: CLI=%d embedded=%d", len(cli.Schema), len(embedded.Schema))
	}
	if len(cli.ResourcesMap) != len(embedded.ResourcesMap) {
		t.Fatalf("resource schema size differs: CLI=%d embedded=%d", len(cli.ResourcesMap), len(embedded.ResourcesMap))
	}
	if len(cli.DataSourcesMap) != len(embedded.DataSourcesMap) {
		t.Fatalf("data-source schema size differs: CLI=%d embedded=%d", len(cli.DataSourcesMap), len(embedded.DataSourcesMap))
	}
}
