// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package provider

import (
	"crypto/rsa"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

// InProcessOptions supplies authentication dependencies to one embedded provider.
// Callers own the token provider's lifetime and renewal. Tokens are never copied
// into the Terraform configuration or shared through process-wide settings.
type InProcessOptions struct {
	// OKEWorkloadIdentityTokenProvider supplies renewable tokens for this
	// provider instance only. A nil interface retains the default pod identity.
	OKEWorkloadIdentityTokenProvider auth.ServiceAccountTokenProvider
}

func firstInProcessOptions(options []InProcessOptions) InProcessOptions {
	if len(options) == 0 {
		return InProcessOptions{}
	}
	return options[0]
}

func (o InProcessOptions) validate(authType string) error {
	if o.OKEWorkloadIdentityTokenProvider != nil {
		value := reflect.ValueOf(o.OKEWorkloadIdentityTokenProvider)
		switch value.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
			if value.IsNil() {
				return fmt.Errorf("the custom OKE token provider must not be a typed nil")
			}
		}
	}
	if o.OKEWorkloadIdentityTokenProvider != nil && !strings.EqualFold(authType, globalvar.AuthOKEWorkloadIdentity) {
		return fmt.Errorf("a custom OKE token provider requires auth OKEWorkloadIdentity")
	}
	return nil
}

// NewSDKv2ProviderForInProcessConfigurationWithOptions constructs an isolated
// configuration-only provider with optional authentication dependencies.
func NewSDKv2ProviderForInProcessConfigurationWithOptions(options InProcessOptions) *schema.Provider {
	return newSDKv2ProviderWithMaps(true, nil, nil, options)
}

// NewFrameworkProviderForInProcessWithOptions constructs an isolated Framework
// provider with optional authentication dependencies.
func NewFrameworkProviderForInProcessWithOptions(options InProcessOptions) provider.Provider {
	return &ociPluginProvider{inProcess: true, inProcessOptions: options}
}

// Token supplier errors may include credentials. Do not expose them through SDK
// logs, Terraform diagnostics or Kubernetes events in the embedding host.
type safeServiceAccountTokenProvider struct {
	auth.ServiceAccountTokenProvider
}

func (p safeServiceAccountTokenProvider) ServiceAccountToken() (string, error) {
	token, err := p.ServiceAccountTokenProvider.ServiceAccountToken()
	if err != nil || strings.TrimSpace(token) == "" {
		return "", fmt.Errorf("cannot obtain the configured ServiceAccount token")
	}
	return token, nil
}

type okeTokenConfigurationProvider struct {
	auth.ConfigurationProviderWithClaimAccess
	region string
}

func (p okeTokenConfigurationProvider) Region() (string, error) {
	if p.region != "" {
		return p.region, nil
	}
	return p.ConfigurationProviderWithClaimAccess.Region()
}

func (p okeTokenConfigurationProvider) Refreshable() bool { return true }

// The SDK's exchange errors can contain remote response bodies. Keep those out
// of host diagnostics as well as errors from the Kubernetes token supplier.
func (p okeTokenConfigurationProvider) KeyID() (string, error) {
	value, err := p.ConfigurationProviderWithClaimAccess.KeyID()
	if err != nil {
		return "", fmt.Errorf("cannot obtain OCI credentials for the configured ServiceAccount")
	}
	return value, nil
}

func (p okeTokenConfigurationProvider) PrivateRSAKey() (*rsa.PrivateKey, error) {
	value, err := p.ConfigurationProviderWithClaimAccess.PrivateRSAKey()
	if err != nil {
		return nil, fmt.Errorf("cannot obtain the OCI signing key for the configured ServiceAccount")
	}
	return value, nil
}

func (p okeTokenConfigurationProvider) TenancyOCID() (string, error) {
	value, err := p.ConfigurationProviderWithClaimAccess.TenancyOCID()
	if err != nil {
		return "", fmt.Errorf("cannot obtain the OCI tenancy for the configured ServiceAccount")
	}
	return value, nil
}

func (p okeTokenConfigurationProvider) GetClaim(claim string) (any, error) {
	value, err := p.ConfigurationProviderWithClaimAccess.GetClaim(claim)
	if err != nil {
		return nil, fmt.Errorf("cannot obtain an OCI credential claim for the configured ServiceAccount")
	}
	return value, nil
}

func (o InProcessOptions) okeConfigurationProvider(region string) (auth.ConfigurationProviderWithClaimAccess, error) {
	if o.OKEWorkloadIdentityTokenProvider == nil {
		return auth.OkeWorkloadIdentityConfigurationProvider()
	}
	p, err := auth.OkeWorkloadIdentityConfigurationProviderWithServiceAccountTokenProvider(
		safeServiceAccountTokenProvider{o.OKEWorkloadIdentityTokenProvider})
	if err != nil {
		return nil, fmt.Errorf("cannot initialize OKE workload identity with the configured ServiceAccount token")
	}
	return okeTokenConfigurationProvider{ConfigurationProviderWithClaimAccess: p, region: region}, nil
}
