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
