// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package provider

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdkterraform "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

type serviceAccountTokenFunc func() (string, error)

func (f serviceAccountTokenFunc) ServiceAccountToken() (string, error) { return f() }

func setupOKEEnvironment(t *testing.T) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	template := &x509.Certificate{SerialNumber: big.NewInt(1), NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), IsCA: true, BasicConstraintsValid: true, KeyUsage: x509.KeyUsageCertSign}
	der, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "ca.crt")
	if err := os.WriteFile(path, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv(auth.ResourcePrincipalVersionEnvVar, "2.2")
	t.Setenv(auth.ResourcePrincipalRegionEnvVar, "us-phoenix-1")
	t.Setenv(auth.KubernetesServiceHostEnvVar, "127.0.0.1")
	t.Setenv(auth.OciKubernetesServiceAccountCertPath, path)
}

func TestInProcessOKEAuthenticationIsolationAndRenewableSupplier(t *testing.T) {
	setupOKEEnvironment(t)
	var first, second atomic.Int64
	makeOptions := func(calls *atomic.Int64) InProcessOptions {
		return InProcessOptions{OKEWorkloadIdentityTokenProvider: serviceAccountTokenFunc(func() (string, error) {
			calls.Add(1)
			return "", fmt.Errorf("sensitive-token-must-not-leak")
		})}
	}
	for _, framework := range []bool{false, true} {
		for i, options := range []InProcessOptions{makeOptions(&first), makeOptions(&second)} {
			var providers []auth.ConfigurationProviderWithClaimAccess
			if framework {
				p := NewFrameworkProviderForInProcessWithOptions(options).(*ociPluginProvider)
				p.auth, p.region = globalvar.AuthOKEWorkloadIdentity, "us-ashburn-1"
				got, err := p._getConfigProviders()
				if err != nil {
					t.Fatal(err)
				}
				providers = append(providers, got[0].(auth.ConfigurationProviderWithClaimAccess))
			} else {
				d := schema.TestResourceDataRaw(t, SchemaMap(), map[string]any{"auth": globalvar.AuthOKEWorkloadIdentity, "region": "us-ashburn-1"})
				got, err := getConfigProvidersForMode(d, strings.ToLower(globalvar.AuthOKEWorkloadIdentity), true, options)
				if err != nil {
					t.Fatal(err)
				}
				providers = append(providers, got[0].(auth.ConfigurationProviderWithClaimAccess))
			}
			p := providers[0]
			if region, err := p.Region(); err != nil || region != "us-ashburn-1" {
				t.Fatalf("region = %q, %v", region, err)
			}
			for range 2 {
				_, err := p.KeyID()
				if err == nil || strings.Contains(err.Error(), "sensitive-token") {
					t.Fatalf("unsafe or missing supplier error: %v", err)
				}
			}
			wantFirst, wantSecond := int64(2), int64(0)
			if framework {
				wantFirst += 2
				wantSecond += 2
			}
			if i == 1 {
				wantSecond += 2
			}
			if first.Load() != wantFirst || second.Load() != wantSecond {
				t.Fatalf("suppliers crossed: first=%d second=%d", first.Load(), second.Load())
			}
		}
	}
}

func TestInProcessOKERejectsIncompatibleAuthentication(t *testing.T) {
	options := InProcessOptions{OKEWorkloadIdentityTokenProvider: serviceAccountTokenFunc(func() (string, error) { t.Fatal("supplier must not be called"); return "", nil })}
	p := NewSDKv2ProviderForInProcessConfigurationWithOptions(options)
	diags := p.Configure(context.Background(), sdkterraform.NewResourceConfigRaw(map[string]any{"auth": "ApiKey"}))
	if !diags.HasError() {
		t.Fatal("SDKv2 accepted custom supplier with ApiKey")
	}
	fw := NewFrameworkProviderForInProcessWithOptions(options).(*ociPluginProvider)
	fw.auth = "ApiKey"
	if _, err := fw._getConfigProviders(); err == nil {
		t.Fatal("Framework accepted custom supplier with ApiKey")
	}
}

func TestInProcessOKEZeroOptionsPreserveDefaultProvider(t *testing.T) {
	setupOKEEnvironment(t)
	p, err := (InProcessOptions{}).okeConfigurationProvider("us-ashburn-1")
	if err != nil {
		t.Fatal(err)
	}
	if region, err := p.Region(); err != nil || region != "us-phoenix-1" {
		t.Fatalf("default OKE region behavior changed: %q %v", region, err)
	}
}

func TestInProcessOKERejectsTypedNilSupplier(t *testing.T) {
	var supplier serviceAccountTokenFunc
	options := InProcessOptions{OKEWorkloadIdentityTokenProvider: supplier}
	for _, authType := range []string{"OKEWorkloadIdentity", "ApiKey"} {
		if err := options.validate(authType); err == nil {
			t.Fatalf("accepted typed nil supplier with %s", authType)
		}
	}
}

func TestInProcessOKEDoesNotFallBackToOtherCredentials(t *testing.T) {
	setupOKEEnvironment(t)
	options := InProcessOptions{OKEWorkloadIdentityTokenProvider: serviceAccountTokenFunc(func() (string, error) { return "", fmt.Errorf("secretJWT") })}
	for _, framework := range []bool{false, true} {
		clients := &tf_client.OracleClients{Configuration: map[string]string{}}
		var configuration auth.ConfigurationProviderWithClaimAccess
		if framework {
			p := NewFrameworkProviderForInProcessWithOptions(options).(*ociPluginProvider)
			p.auth, p.region, p.tenancyOcid = globalvar.AuthOKEWorkloadIdentity, "us-ashburn-1", "fallback-tenancy"
			got, err := p._GetSdkConfigProvider(clients)
			if err != nil {
				t.Fatal(err)
			}
			configuration = got.(auth.ConfigurationProviderWithClaimAccess)
		} else {
			d := schema.TestResourceDataRaw(t, SchemaMap(), map[string]any{"auth": globalvar.AuthOKEWorkloadIdentity, "region": "us-ashburn-1", "tenancy_ocid": "fallback-tenancy"})
			got, err := getSdkConfigProvider(d, clients, true, options)
			if err != nil {
				t.Fatal(err)
			}
			configuration = got.(auth.ConfigurationProviderWithClaimAccess)
		}
		if value, err := configuration.TenancyOCID(); err == nil || value != "" || strings.Contains(err.Error(), "secretJWT") {
			t.Fatalf("fell back to ambient credentials: tenancy=%q error=%v", value, err)
		}
	}
}

type failedExchangeConfiguration struct {
	auth.ConfigurationProviderWithClaimAccess
}

func (failedExchangeConfiguration) KeyID() (string, error) {
	return "secretJWT", fmt.Errorf("secretJWT")
}
func (failedExchangeConfiguration) PrivateRSAKey() (*rsa.PrivateKey, error) {
	return nil, fmt.Errorf("secretJWT")
}
func (failedExchangeConfiguration) TenancyOCID() (string, error) {
	return "secretJWT", fmt.Errorf("secretJWT")
}
func (failedExchangeConfiguration) GetClaim(string) (any, error) {
	return "secretJWT", fmt.Errorf("secretJWT")
}

func TestInProcessOKEExchangeErrorsAreSanitized(t *testing.T) {
	p := okeTokenConfigurationProvider{ConfigurationProviderWithClaimAccess: failedExchangeConfiguration{}}
	for _, call := range []func() (any, error){
		func() (any, error) { value, err := p.KeyID(); return value, err },
		func() (any, error) { value, err := p.PrivateRSAKey(); return value, err },
		func() (any, error) { value, err := p.TenancyOCID(); return value, err },
		func() (any, error) { return p.GetClaim("res_tenant") },
	} {
		value, err := call()
		if err == nil || strings.Contains(fmt.Sprint(value, err), "secretJWT") {
			t.Fatalf("exchange failure exposed credentials: %v %v", value, err)
		}
	}
}

func TestSafeServiceAccountTokenProviderRefreshAndErrors(t *testing.T) {
	value := "first"
	p := safeServiceAccountTokenProvider{serviceAccountTokenFunc(func() (string, error) { return value, nil })}
	for _, want := range []string{"first", "second"} {
		value = want
		if got, err := p.ServiceAccountToken(); err != nil || got != want {
			t.Fatalf("token=%q error=%v", got, err)
		}
	}
	value = " "
	if _, err := p.ServiceAccountToken(); err == nil {
		t.Fatal("accepted empty token")
	}
}
