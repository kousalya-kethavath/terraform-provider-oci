// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package vault

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/oracle/terraform-provider-oci/internal/client"
)

func TestReadReturnsLazyClientInitializationError(t *testing.T) {
	dataSource := &VaultSecretVersionDataSource{}
	request := datasource.ConfigureRequest{
		ProviderData: &client.OracleClients{SdkClientMap: make(map[string]interface{})},
	}
	response := &datasource.ConfigureResponse{}

	dataSource.Configure(t.Context(), request, response)
	if response.Diagnostics.HasError() {
		t.Fatalf("Configure returned diagnostics: %v", response.Diagnostics)
	}

	readResponse := &datasource.ReadResponse{}
	dataSource.Read(t.Context(), datasource.ReadRequest{}, readResponse)
	if !readResponse.Diagnostics.HasError() {
		t.Fatal("Read did not return the lazy Vault client initialization error")
	}
}

func TestConfigureRejectsUnexpectedProviderData(t *testing.T) {
	dataSource := &VaultSecretVersionDataSource{}
	request := datasource.ConfigureRequest{ProviderData: "unexpected"}
	response := &datasource.ConfigureResponse{}

	dataSource.Configure(t.Context(), request, response)

	if !response.Diagnostics.HasError() {
		t.Fatal("Configure accepted an unexpected provider data type")
	}
}
