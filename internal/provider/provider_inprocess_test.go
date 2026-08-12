// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package provider

import (
	"context"
	"maps"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	frameworktypes "github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	tfclient "github.com/oracle/terraform-provider-oci/internal/client"
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
	assertSDKv2ResourceMapsIsolated(t, firstEmbedded.ResourcesMap, secondEmbedded.ResourcesMap)
	assertSDKv2ResourceMapsIsolated(t, firstEmbedded.DataSourcesMap, secondEmbedded.DataSourcesMap)
}

func TestSelectiveInProcessProvider(t *testing.T) {
	const resourceName = "oci_identity_tag_namespace"
	enabledServices := maps.Clone(oci_common.OciSdkEnabledServicesMap)
	first, err := NewSDKv2ProviderForInProcessResources(resourceName, resourceName)
	if err != nil {
		t.Fatalf("construct selective provider: %v", err)
	}
	second, err := NewSDKv2ProviderForInProcessResources(resourceName)
	if err != nil {
		t.Fatalf("construct second selective provider: %v", err)
	}
	if len(first.ResourcesMap) != 1 || first.ResourcesMap[resourceName] == nil {
		t.Fatalf("selective resource map = %v, want only %q", first.ResourcesMap, resourceName)
	}
	if len(first.DataSourcesMap) != 0 {
		t.Fatalf("selective provider retained %d data sources, want 0", len(first.DataSourcesMap))
	}
	if err := first.InternalValidate(); err != nil {
		t.Fatalf("selective provider validation failed: %v", err)
	}
	if first.ResourcesMap[resourceName] == second.ResourcesMap[resourceName] {
		t.Fatal("selective providers share a mutable resource schema")
	}

	if _, err := NewSDKv2ProviderForInProcessResources("oci_missing_resource"); err == nil {
		t.Fatal("selective provider accepted an unknown resource")
	}
	if !maps.Equal(enabledServices, oci_common.OciSdkEnabledServicesMap) {
		t.Fatal("selective schema construction changed OCI SDK enabled services")
	}
}

func TestConfigurationOnlyInProcessProvider(t *testing.T) {
	p := NewSDKv2ProviderForInProcessConfiguration()
	if len(p.Schema) == 0 {
		t.Fatal("configuration-only provider has no provider schema")
	}
	if p.ConfigureFunc == nil {
		t.Fatal("configuration-only provider has no ConfigureFunc")
	}
	if len(p.ResourcesMap) != 0 || len(p.DataSourcesMap) != 0 {
		t.Fatalf("configuration-only provider retained resources=%d dataSources=%d", len(p.ResourcesMap), len(p.DataSourcesMap))
	}
	if err := p.InternalValidate(); err != nil {
		t.Fatalf("configuration-only provider validation failed: %v", err)
	}
}

func TestCloneSDKv2ResourceIsolatesMutableStructures(t *testing.T) {
	timeout := time.Minute
	source := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"field": {
				Type:          schema.TypeString,
				ConflictsWith: []string{"other"},
				ExactlyOneOf:  []string{"field", "other"},
				AtLeastOneOf:  []string{"field", "other"},
				RequiredWith:  []string{"other"},
			},
		},
		SchemaFunc: func() map[string]*schema.Schema {
			return map[string]*schema.Schema{"dynamic": {Type: schema.TypeString}}
		},
		StateUpgraders:                 make([]schema.StateUpgrader, 1),
		ValidateRawResourceConfigFuncs: make([]schema.ValidateRawResourceConfigFunc, 1),
		Importer:                       &schema.ResourceImporter{},
		Timeouts:                       &schema.ResourceTimeout{Create: &timeout},
	}

	cloned := cloneSDKv2Resource(source)
	if cloned == source || cloned.Schema["field"] == source.Schema["field"] {
		t.Fatal("resource schema was not cloned")
	}
	if &cloned.Schema["field"].ConflictsWith[0] == &source.Schema["field"].ConflictsWith[0] ||
		&cloned.Schema["field"].ExactlyOneOf[0] == &source.Schema["field"].ExactlyOneOf[0] ||
		&cloned.Schema["field"].AtLeastOneOf[0] == &source.Schema["field"].AtLeastOneOf[0] ||
		&cloned.Schema["field"].RequiredWith[0] == &source.Schema["field"].RequiredWith[0] {
		t.Fatal("schema constraint slices were not cloned")
	}
	if &cloned.StateUpgraders[0] == &source.StateUpgraders[0] ||
		&cloned.ValidateRawResourceConfigFuncs[0] == &source.ValidateRawResourceConfigFuncs[0] {
		t.Fatal("resource slices were not cloned")
	}
	if cloned.Importer == source.Importer || cloned.Timeouts == source.Timeouts || cloned.Timeouts.Create == source.Timeouts.Create {
		t.Fatal("resource pointer fields were not cloned")
	}

	firstDynamic := cloned.SchemaFunc()
	secondDynamic := cloned.SchemaFunc()
	if firstDynamic["dynamic"] == secondDynamic["dynamic"] {
		t.Fatal("SchemaFunc returned shared schemas")
	}

	cloned.Schema["field"].ConflictsWith[0] = "changed"
	*cloned.Timeouts.Create = 2 * time.Minute
	if source.Schema["field"].ConflictsWith[0] != "other" || *source.Timeouts.Create != time.Minute {
		t.Fatal("mutating a cloned resource affected its source")
	}
}

func TestClonedSDKv2ResourceReturnsLazyClientInitializationErrors(t *testing.T) {
	resource := cloneSDKv2Resource(&schema.Resource{
		Importer: &schema.ResourceImporter{
			StateContext: func(_ context.Context, _ *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				meta.(*tfclient.OracleClients).GetClient("oci_missing.Client")
				return nil, nil
			},
		},
		Create: func(_ *schema.ResourceData, meta interface{}) error {
			meta.(*tfclient.OracleClients).GetClient("oci_missing.Client")
			return nil
		},
		ReadContext: func(_ context.Context, _ *schema.ResourceData, meta interface{}) diag.Diagnostics {
			meta.(*tfclient.OracleClients).GetClient("oci_missing.Client")
			return nil
		},
	})
	clients := &tfclient.OracleClients{SdkClientMap: make(map[string]interface{})}

	err := resource.Create(nil, clients)
	if err == nil || !strings.Contains(err.Error(), "provider clients are not configured") {
		t.Fatalf("Create lazy-client error = %v", err)
	}
	diagnostics := resource.ReadContext(t.Context(), nil, clients)
	if !diagnostics.HasError() || !strings.Contains(diagnostics[0].Summary, "provider clients are not configured") {
		t.Fatalf("ReadContext lazy-client diagnostics = %v", diagnostics)
	}
	_, err = resource.Importer.StateContext(t.Context(), nil, clients)
	if err == nil || !strings.Contains(err.Error(), "provider clients are not configured") {
		t.Fatalf("import lazy-client error = %v", err)
	}
}

func TestClonedSDKv2ResourceDoesNotMaskUnrelatedPanics(t *testing.T) {
	resource := cloneSDKv2Resource(&schema.Resource{
		Create: func(*schema.ResourceData, interface{}) error {
			panic("programming error")
		},
	})

	defer func() {
		if recovered := recover(); recovered != "programming error" {
			t.Fatalf("recovered panic = %v", recovered)
		}
	}()
	_ = resource.Create(nil, nil)
}

func assertSDKv2ResourceMapsIsolated(t *testing.T, first, second map[string]*schema.Resource) {
	t.Helper()
	if len(first) != len(second) {
		t.Fatalf("resource map sizes differ: first=%d second=%d", len(first), len(second))
	}
	for name, firstResource := range first {
		secondResource, ok := second[name]
		if !ok {
			t.Fatalf("resource %q is missing from the second provider", name)
		}
		assertSDKv2ResourcesIsolated(t, name, firstResource, secondResource)
	}
}

func assertSDKv2ResourcesIsolated(t *testing.T, path string, first, second *schema.Resource) {
	t.Helper()
	if first == nil || second == nil {
		if first != second {
			t.Fatalf("resource %q differs between provider instances", path)
		}
		return
	}
	if first == second {
		t.Fatalf("resource %q is shared between provider instances", path)
	}
	assertSDKv2SchemaMapsIsolated(t, path, first.Schema, second.Schema)

	if first.Importer != nil && first.Importer == second.Importer {
		t.Fatalf("resource %q importer is shared between provider instances", path)
	}
	if first.Timeouts != nil && first.Timeouts == second.Timeouts {
		t.Fatalf("resource %q timeouts are shared between provider instances", path)
	}
	if first.Identity != nil && first.Identity == second.Identity {
		t.Fatalf("resource %q identity is shared between provider instances", path)
	}
}

func assertSDKv2SchemaMapsIsolated(t *testing.T, path string, first, second map[string]*schema.Schema) {
	t.Helper()
	if len(first) != len(second) {
		t.Fatalf("schema map sizes differ for %q: first=%d second=%d", path, len(first), len(second))
	}
	for name, firstSchema := range first {
		secondSchema, ok := second[name]
		if !ok {
			t.Fatalf("schema %q is missing from the second provider", path+"."+name)
		}
		assertSDKv2SchemasIsolated(t, path+"."+name, firstSchema, secondSchema)
	}
}

func assertSDKv2SchemasIsolated(t *testing.T, path string, first, second *schema.Schema) {
	t.Helper()
	if first == nil || second == nil {
		if first != second {
			t.Fatalf("schema %q differs between provider instances", path)
		}
		return
	}
	if first == second {
		t.Fatalf("schema %q is shared between provider instances", path)
	}

	switch firstElement := first.Elem.(type) {
	case *schema.Schema:
		secondElement, ok := second.Elem.(*schema.Schema)
		if !ok {
			t.Fatalf("schema element type differs for %q", path)
		}
		assertSDKv2SchemasIsolated(t, path+"[]", firstElement, secondElement)
	case *schema.Resource:
		secondElement, ok := second.Elem.(*schema.Resource)
		if !ok {
			t.Fatalf("schema element type differs for %q", path)
		}
		assertSDKv2ResourcesIsolated(t, path+"[]", firstElement, secondElement)
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

func TestFrameworkInProcessProviderRejectsIgnoreDefinedTags(t *testing.T) {
	p := &ociPluginProvider{inProcess: true}
	value := frameworktypes.ListValueMust(
		frameworktypes.StringType,
		[]attr.Value{frameworktypes.StringValue("Oracle-Tags.CreatedBy")},
	)
	if diags := p.setIgnoreDefinedTags(t.Context(), value); diags.HasError() {
		t.Fatalf("decode ignore_defined_tags: %v", diags)
	}
	if len(p.ignoreDefinedTags) != 1 || p.ignoreDefinedTags[0] != "Oracle-Tags.CreatedBy" {
		t.Fatalf("unexpected decoded ignore_defined_tags: %v", p.ignoreDefinedTags)
	}

	err := p.validateInProcessProviderConfig()
	if err == nil {
		t.Fatal("Framework in-process provider accepted ignore_defined_tags")
	}
	if !strings.Contains(err.Error(), globalvar.DefinedTagsToIgnore) {
		t.Fatalf("validation error does not identify ignore_defined_tags: %v", err)
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

func BenchmarkInProcessProviderConstruction(b *testing.B) {
	b.Run("full", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_ = NewSDKv2ProviderForInProcess()
		}
	})
	b.Run("identity-resource", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			if _, err := NewSDKv2ProviderForInProcessResources("oci_identity_tag_namespace"); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("configuration-only", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_ = NewSDKv2ProviderForInProcessConfiguration()
		}
	})
}
