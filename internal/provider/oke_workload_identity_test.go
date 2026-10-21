// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package provider

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	tfclient "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

func TestOKETokenFileRotationAndFailures(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "token")
	p := okeTokenFile{path: path}
	for _, token := range []string{"first-jwt", "second-jwt"} {
		next := filepath.Join(dir, "next")
		if err := os.WriteFile(next, []byte(token+"\n"), 0600); err != nil {
			t.Fatal(err)
		}
		if err := os.Rename(next, path); err != nil {
			t.Fatal(err)
		}
		if got, err := p.ServiceAccountToken(); err != nil || got != token {
			t.Fatalf("rotation: %q %v", got, err)
		}
	}
	for _, content := range []string{"", strings.Repeat("x", 128*1024+1)} {
		if err := os.WriteFile(path, []byte(content), 0600); err != nil {
			t.Fatal(err)
		}
		if got, err := p.ServiceAccountToken(); err == nil || got != "" {
			t.Fatal("accepted invalid file")
		}
	}
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
	if got, err := p.ServiceAccountToken(); err == nil || got != "" || strings.Contains(err.Error(), path) {
		t.Fatalf("unsafe read failure: %q %v", got, err)
	}
}

func TestOKETokenSourceValidation(t *testing.T) {
	path := filepath.Join(t.TempDir(), "token")
	callback := InProcessOptions{OKEWorkloadIdentityTokenProvider: serviceAccountTokenFunc(func() (string, error) { return "callback", nil })}
	for _, tc := range []struct {
		options    InProcessOptions
		auth, path string
	}{
		{InProcessOptions{}, "ApiKey", path},
		{InProcessOptions{}, "OKEWorkloadIdentity", "relative/token"},
		{callback, "OKEWorkloadIdentity", path},
	} {
		if _, err := tc.options.withOKETokenPath(tc.auth, tc.path); err == nil {
			t.Fatal("accepted incompatible token source")
		}
	}
	if _, err := (InProcessOptions{}).withOKETokenPath("OKEWorkloadIdentity", path); err != nil {
		t.Fatal(err)
	}
}

func TestOKECLITokenSourceIsExclusiveForBothProviders(t *testing.T) {
	setupOKEEnvironment(t)
	path := filepath.Join(t.TempDir(), "missing-token")
	for _, framework := range []bool{false, true} {
		clients := &tfclient.OracleClients{Configuration: map[string]string{}}
		var p auth.ConfigurationProviderWithClaimAccess
		if framework {
			fw := New().(*ociPluginProvider) // Normal CLI constructor, not embedded.
			fw.auth, fw.region, fw.okeWorkloadIdentityTokenPath, fw.tenancyOcid = "OKEWorkloadIdentity", "us-ashburn-1", path, "ambient-tenancy"
			got, err := fw._GetSdkConfigProvider(clients)
			if err != nil {
				t.Fatal(err)
			}
			p = got.(auth.ConfigurationProviderWithClaimAccess)
		} else {
			d := schema.TestResourceDataRaw(t, SchemaMap(), map[string]any{"auth": "OKEWorkloadIdentity", "region": "us-ashburn-1", "tenancy_ocid": "ambient-tenancy", globalvar.OKEWorkloadIdentityTokenPathAttrName: path})
			got, err := getSdkConfigProvider(d, clients, false)
			if err != nil {
				t.Fatal(err)
			}
			p = got.(auth.ConfigurationProviderWithClaimAccess)
		}
		if region, _ := p.Region(); region != "us-ashburn-1" {
			t.Fatalf("wrong client region %q", region)
		}
		if value, err := p.TenancyOCID(); err == nil || value != "" || strings.Contains(err.Error(), path) {
			t.Fatalf("unsafe ambient fallback: %q %v", value, err)
		}
	}
}

// Build protocol configuration so unknown values reach the real Configure
// entry points without being reduced to Go zero values by test fixtures.
func okeTokenPathTestConfig(typ tftypes.Type, tokenPath any) tftypes.Value {
	attributes := typ.(tftypes.Object).AttributeTypes
	values := make(map[string]tftypes.Value, len(attributes))
	for name, attrType := range attributes {
		values[name] = tftypes.NewValue(attrType, nil)
	}
	values[globalvar.AuthAttrName] = tftypes.NewValue(tftypes.String, globalvar.AuthOKEWorkloadIdentity)
	values[globalvar.RegionAttrName] = tftypes.NewValue(tftypes.String, "us-ashburn-1")
	values[globalvar.OKEWorkloadIdentityTokenPathAttrName] = tftypes.NewValue(tftypes.String, tokenPath)
	return tftypes.NewValue(typ, values)
}

func TestOKETokenPathConfigureValues(t *testing.T) {
	// If a regression selects default pod auth, stop SDK initialization before
	// it can contact OKE. The expected diagnostic must identify the unknown path.
	t.Setenv(auth.ResourcePrincipalVersionEnvVar, "")
	for _, tc := range []struct {
		name       string
		value      any
		diagnostic string
	}{
		{"unknown", tftypes.UnknownValue, "oke_workload_identity_token_path must be known"},
		{"null retains default pod identity", nil, "OCI_RESOURCE_PRINCIPAL_VERSION"},
		{"known selects explicit token source", "/run/oci/test.jwt", "cannot initialize OKE workload identity with the configured ServiceAccount token"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, inProcess := range []bool{false, true} {
				t.Run(fmt.Sprintf("SDKv2/inProcess=%t", inProcess), func(t *testing.T) {
					p := newSDKv2ProviderWithMaps(inProcess, nil, nil)
					server := p.GRPCProvider()
					schemaResp, err := server.GetProviderSchema(t.Context(), &tfprotov5.GetProviderSchemaRequest{})
					if err != nil {
						t.Fatal(err)
					}
					typ := schemaResp.Provider.ValueType()
					raw, err := tfprotov5.NewDynamicValue(typ, okeTokenPathTestConfig(typ, tc.value))
					if err != nil {
						t.Fatal(err)
					}
					response, err := server.ConfigureProvider(t.Context(), &tfprotov5.ConfigureProviderRequest{TerraformVersion: "1.4.6", Config: &raw})
					if err != nil {
						t.Fatal(err)
					}
					if len(response.Diagnostics) != 1 || response.Diagnostics[0].Severity != tfprotov5.DiagnosticSeverityError || !strings.Contains(response.Diagnostics[0].Summary+response.Diagnostics[0].Detail, tc.diagnostic) {
						t.Fatalf("unexpected Configure diagnostics: %v", response.Diagnostics)
					}
					if p.Meta() != nil {
						t.Fatal("unknown path configured provider clients")
					}
				})
				t.Run(fmt.Sprintf("Framework/inProcess=%t", inProcess), func(t *testing.T) {
					p := &ociPluginProvider{inProcess: inProcess}
					var schemaResp fwprovider.SchemaResponse
					p.Schema(t.Context(), fwprovider.SchemaRequest{}, &schemaResp)
					typ := schemaResp.Schema.Type().TerraformType(t.Context())
					var response fwprovider.ConfigureResponse
					p.Configure(t.Context(), fwprovider.ConfigureRequest{TerraformVersion: "1.4.6", Config: tfsdk.Config{Schema: schemaResp.Schema, Raw: okeTokenPathTestConfig(typ, tc.value)}}, &response)
					if len(response.Diagnostics) != 1 || !response.Diagnostics.HasError() || !strings.Contains(response.Diagnostics[0].Detail(), tc.diagnostic) {
						t.Fatalf("unexpected Configure diagnostics: %v", response.Diagnostics)
					}
					if response.ResourceData != nil || response.DataSourceData != nil {
						t.Fatal("unknown path configured provider clients")
					}
				})
			}
		})
	}
}
