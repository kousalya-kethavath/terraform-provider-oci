// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package client

import (
	"crypto/rand"
	"crypto/rsa"
	"os"
	"sync"
	"sync/atomic"
	"testing"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"
	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"
)

func TestObjectStorageClientConstructionDoesNotMutateEnvironment(t *testing.T) {
	const envName = "OCI_REALM_SPECIFIC_SERVICE_ENDPOINT_TEMPLATE_ENABLED"
	t.Setenv(envName, "original")
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}
	_, err = initObjectstorageObjectStorageClient(
		endpointTestConfiguration{privateKey: privateKey},
		func(*oci_common.BaseClient) error { return nil },
		ServiceClientOverrides{},
	)
	if err != nil {
		t.Fatalf("construct Object Storage client: %v", err)
	}
	if got := os.Getenv(envName); got != "original" {
		t.Fatalf("environment changed from %q to %q", "original", got)
	}
}

func TestOracleClientsConfigureBaseClientIsolation(t *testing.T) {
	var firstCalls atomic.Int64
	var secondCalls atomic.Int64
	first := &OracleClients{ConfigureClient: func(client *oci_common.BaseClient) error {
		firstCalls.Add(1)
		client.UserAgent = "first"
		return nil
	}}
	second := &OracleClients{ConfigureClient: func(client *oci_common.BaseClient) error {
		secondCalls.Add(1)
		client.UserAgent = "second"
		return nil
	}}

	const calls = 100
	var wg sync.WaitGroup
	for range calls {
		wg.Go(func() {
			client := &oci_common.BaseClient{}
			if err := first.ConfigureBaseClient(client); err != nil {
				t.Errorf("configure first client: %v", err)
			} else if client.UserAgent != "first" {
				t.Errorf("first client user agent = %q", client.UserAgent)
			}
		})
		wg.Go(func() {
			client := &oci_common.BaseClient{}
			if err := second.ConfigureBaseClient(client); err != nil {
				t.Errorf("configure second client: %v", err)
			} else if client.UserAgent != "second" {
				t.Errorf("second client user agent = %q", client.UserAgent)
			}
		})
	}
	wg.Wait()
	if firstCalls.Load() != calls || secondCalls.Load() != calls {
		t.Fatalf("callback calls = (%d, %d), want (%d, %d)", firstCalls.Load(), secondCalls.Load(), calls, calls)
	}
}

func TestOracleClientsConfigureBaseClientRequiresCallback(t *testing.T) {
	if err := (&OracleClients{}).ConfigureBaseClient(&oci_common.BaseClient{}); err == nil {
		t.Fatal("ConfigureBaseClient returned nil without an instance callback")
	}
}

func TestEndpointSpecificClientsUseOwningInstanceCallback(t *testing.T) {
	first := newEndpointTestClients(t, "first-provider")
	second := newEndpointTestClients(t, "second-provider")

	tests := map[string]func(*OracleClients) (*oci_common.BaseClient, error){
		"Functions": func(clients *OracleClients) (*oci_common.BaseClient, error) {
			client, err := clients.FunctionsInvokeClientWithEndpoint("https://functions.example.com")
			return &client.BaseClient, err
		},
		"KMS crypto": func(clients *OracleClients) (*oci_common.BaseClient, error) {
			client, err := clients.KmsCryptoClientWithEndpoint("https://kms-crypto.example.com")
			return &client.BaseClient, err
		},
		"KMS management": func(clients *OracleClients) (*oci_common.BaseClient, error) {
			client, err := clients.KmsManagementClientWithEndpoint("https://kms-management.example.com")
			return &client.BaseClient, err
		},
		"Identity Domains": func(clients *OracleClients) (*oci_common.BaseClient, error) {
			client, err := clients.IdentityDomainsClientWithEndpoint("https://identity-domains.example.com")
			return &client.BaseClient, err
		},
	}

	for name, create := range tests {
		t.Run(name, func(t *testing.T) {
			firstClient, err := create(first)
			if err != nil {
				t.Fatalf("create first endpoint client: %v", err)
			}
			secondClient, err := create(second)
			if err != nil {
				t.Fatalf("create second endpoint client: %v", err)
			}
			if firstClient.UserAgent != "first-provider" || secondClient.UserAgent != "second-provider" {
				t.Fatalf("endpoint client configuration crossed instances: first=%q second=%q", firstClient.UserAgent, secondClient.UserAgent)
			}
		})
	}
}

func newEndpointTestClients(t *testing.T, userAgent string) *OracleClients {
	t.Helper()
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}
	config := endpointTestConfiguration{privateKey: privateKey}
	functionsClient, err := oci_functions.NewFunctionsInvokeClientWithConfigurationProvider(config, "https://functions-primary.example.com")
	if err != nil {
		t.Fatal(err)
	}
	kmsCryptoClient, err := oci_kms.NewKmsCryptoClientWithConfigurationProvider(config, "https://kms-crypto-primary.example.com")
	if err != nil {
		t.Fatal(err)
	}
	kmsManagementClient, err := oci_kms.NewKmsManagementClientWithConfigurationProvider(config, "https://kms-management-primary.example.com")
	if err != nil {
		t.Fatal(err)
	}
	identityDomainsClient, err := oci_identity_domains.NewIdentityDomainsClientWithConfigurationProvider(config, "https://identity-domains-primary.example.com")
	if err != nil {
		t.Fatal(err)
	}
	return &OracleClients{
		SdkClientMap: map[string]interface{}{
			"oci_functions.FunctionsInvokeClient":        &functionsClient,
			"oci_kms.KmsCryptoClient":                    &kmsCryptoClient,
			"oci_kms.KmsManagementClient":                &kmsManagementClient,
			"oci_identity_domains.IdentityDomainsClient": &identityDomainsClient,
		},
		ConfigureClient: func(client *oci_common.BaseClient) error {
			client.UserAgent = userAgent
			return nil
		},
	}
}

type endpointTestConfiguration struct{ privateKey *rsa.PrivateKey }

func (c endpointTestConfiguration) PrivateRSAKey() (*rsa.PrivateKey, error) { return c.privateKey, nil }
func (endpointTestConfiguration) KeyID() (string, error)                    { return "tenancy/user/fingerprint", nil }
func (endpointTestConfiguration) TenancyOCID() (string, error)              { return "tenancy", nil }
func (endpointTestConfiguration) UserOCID() (string, error)                 { return "user", nil }
func (endpointTestConfiguration) KeyFingerprint() (string, error)           { return "fingerprint", nil }
func (endpointTestConfiguration) Region() (string, error)                   { return "us-ashburn-1", nil }
func (endpointTestConfiguration) AuthType() (oci_common.AuthConfig, error) {
	return oci_common.AuthConfig{AuthType: oci_common.UserPrincipal}, nil
}
