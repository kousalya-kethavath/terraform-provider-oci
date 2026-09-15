// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package oci_test

import (
	"context"
	"sync"
	"testing"

	frameworkprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/oracle/terraform-provider-oci/oci"
)

func TestProviderReturnsFreshSDKv2Instances(t *testing.T) {
	const count = 8
	providers := make([]any, count)

	var wg sync.WaitGroup
	for i := range count {
		wg.Go(func() {
			providers[i] = oci.Provider()
		})
	}
	wg.Wait()

	for i := range providers {
		if providers[i] == nil {
			t.Fatalf("Provider() result %d is nil", i)
		}
		for j := range i {
			if providers[i] == providers[j] {
				t.Fatalf("Provider() calls %d and %d returned the same SDKv2 instance", i, j)
			}
		}
	}
}

func TestNewReturnsFreshFrameworkInstances(t *testing.T) {
	first := oci.New()
	second := oci.New()
	if first == nil || second == nil {
		t.Fatal("New() returned a nil Plugin Framework provider")
	}
	if first == second {
		t.Fatal("New() returned a shared Plugin Framework provider instance")
	}

	var firstMetadata frameworkprovider.MetadataResponse
	var secondMetadata frameworkprovider.MetadataResponse
	first.Metadata(context.Background(), frameworkprovider.MetadataRequest{}, &firstMetadata)
	second.Metadata(context.Background(), frameworkprovider.MetadataRequest{}, &secondMetadata)
	if firstMetadata.TypeName != "oci" || secondMetadata.TypeName != "oci" {
		t.Fatalf("unexpected Framework provider type names: %q, %q", firstMetadata.TypeName, secondMetadata.TypeName)
	}
}

func TestSelectiveSDKv2Constructors(t *testing.T) {
	const resourceName = "oci_identity_tag_namespace"
	p, err := oci.ProviderForResources(resourceName)
	if err != nil {
		t.Fatalf("ProviderForResources: %v", err)
	}
	if len(p.ResourcesMap) != 1 || p.ResourcesMap[resourceName] == nil || len(p.DataSourcesMap) != 0 {
		t.Fatalf("unexpected selective provider maps: resources=%d dataSources=%d", len(p.ResourcesMap), len(p.DataSourcesMap))
	}

	configuration := oci.ProviderForConfiguration()
	if len(configuration.Schema) == 0 || configuration.ConfigureFunc == nil {
		t.Fatal("ProviderForConfiguration returned an unusable provider")
	}
	if len(configuration.ResourcesMap) != 0 || len(configuration.DataSourcesMap) != 0 {
		t.Fatalf("configuration-only provider retained resources=%d dataSources=%d", len(configuration.ResourcesMap), len(configuration.DataSourcesMap))
	}
}
