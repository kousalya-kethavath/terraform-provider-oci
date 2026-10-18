// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package provider

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

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
