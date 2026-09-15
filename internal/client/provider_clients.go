// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"

	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"

	"github.com/oracle/terraform-provider-oci/internal/globalvar"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	utils "github.com/oracle/terraform-provider-oci/internal/utils"
)

var OracleClientRegistrationsVar *OracleClientRegistrations // This is a global registration for all oracle clients. This is invariant information about all clients regardless of region

func RegisterOracleClient(name string, client *OracleClient) {
	if OracleClientRegistrationsVar == nil {
		OracleClientRegistrationsVar = &OracleClientRegistrations{
			RegisteredClients: make(map[string]*OracleClient),
		}
	}
	OracleClientRegistrationsVar.RegisteredClients[name] = client
}

type ConfigureClient func(client *oci_common.BaseClient) error

type InitSdkClientFn func(oci_common.ConfigurationProvider, ConfigureClient, ServiceClientOverrides) (interface{}, error)

type OracleClientRegistrations struct {
	RegisteredClients map[string]*OracleClient
}

type ServiceClientOverrides struct {
	HostUrlOverride string
}

type OracleClient struct {
	InitClientFn InitSdkClientFn
}

type OracleClients struct {
	Configuration     map[string]string
	SdkClientMap      map[string]interface{}
	WorkRequestClient *oci_work_requests.WorkRequestClient
	configureClient   ConfigureClient

	clientMu            sync.Mutex
	configProvider      oci_common.ConfigurationProvider
	clientHostOverrides map[string]string
}

// LazyClientInitializationError identifies an expected failure while creating
// an OCI SDK client on first use. SDKv2 resource callbacks recover only this
// panic type and return it as a Terraform diagnostic; unrelated programming
// panics continue to propagate.
type LazyClientInitializationError struct {
	err error
}

func (e *LazyClientInitializationError) Error() string {
	return e.err.Error()
}

func (e *LazyClientInitializationError) Unwrap() error {
	return e.err
}

// GetClient preserves the existing accessor contract used by generated SDKv2
// code. The Terraform CLI path finds an eagerly initialized client, while the
// in-process path constructs it on first use. Because this legacy signature
// cannot return an error, lazy initialization failures use a typed panic that
// the in-process SDKv2 callback wrapper converts to a Terraform error. Callers
// with an error-returning path should use GetClientWithError directly.
func (m *OracleClients) GetClient(name string) interface{} {
	client, err := m.GetClientWithError(name)
	if err != nil {
		panic(&LazyClientInitializationError{err: err})
	}
	return client
}

// GetClientWithError returns an OCI SDK client, constructing it on first use.
// Client construction is serialized per provider instance so concurrent
// reconciliations cannot create duplicate clients or observe a partial map.
func (m *OracleClients) GetClientWithError(name string) (interface{}, error) {
	if m == nil {
		return nil, fmt.Errorf("cannot initialize OCI SDK client %q: provider clients are nil", name)
	}

	m.clientMu.Lock()
	defer m.clientMu.Unlock()

	if client, ok := m.SdkClientMap[name]; ok {
		return client, nil
	}
	if m.configProvider == nil || m.configureClient == nil {
		return nil, fmt.Errorf("cannot initialize OCI SDK client %q: provider clients are not configured", name)
	}
	if OracleClientRegistrationsVar == nil {
		return nil, fmt.Errorf("cannot initialize OCI SDK client %q: there are no client registrations", name)
	}
	registration, ok := OracleClientRegistrationsVar.RegisteredClients[name]
	if !ok || registration == nil || registration.InitClientFn == nil {
		return nil, fmt.Errorf("cannot initialize OCI SDK client %q: client is not registered", name)
	}
	if !common.CheckForEnabledServices(utils.GetSDKServiceName(name)) {
		return nil, fmt.Errorf("cannot initialize OCI SDK client %q: service is disabled", name)
	}

	overrides := ServiceClientOverrides{}
	if host, ok := m.clientHostOverrides[name]; ok {
		overrides.HostUrlOverride = host
	}
	client, err := registration.InitClientFn(m.configProvider, m.configureClient, overrides)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize OCI SDK client %q: %w", name, err)
	}
	m.SdkClientMap[name] = client
	return client, nil
}

// ConfigureBaseClient applies the configuration captured by this provider
// instance to a client created after the initial provider configuration.
func (m *OracleClients) ConfigureBaseClient(client *oci_common.BaseClient) error {
	if m == nil || m.configureClient == nil {
		return fmt.Errorf("cannot configure OCI client: no configure client is registered")
	}
	return m.configureClient(client)
}

// The following clients require special endpoint information that is only known at Terraform apply time; so they
// Create duplicate clients reusing the same Configuration provider as the initialized client and adding the endpoint
// here.
func (m *OracleClients) FunctionsInvokeClientWithEndpoint(endpoint string) (*oci_functions.FunctionsInvokeClient, error) {
	if client, err := oci_functions.NewFunctionsInvokeClientWithConfigurationProvider(*m.FunctionsInvokeClient().ConfigurationProvider(), endpoint); err == nil {
		if err = m.ConfigureBaseClient(&client.BaseClient); err != nil {
			return nil, err
		}
		return &client, nil
	} else {
		return nil, err
	}
}
func (m *OracleClients) KmsCryptoClientWithEndpoint(endpoint string) (*oci_kms.KmsCryptoClient, error) {
	if client, err := oci_kms.NewKmsCryptoClientWithConfigurationProvider(*m.KmsCryptoClient().ConfigurationProvider(), endpoint); err == nil {
		if err = m.ConfigureBaseClient(&client.BaseClient); err != nil {
			return nil, err
		}
		return &client, nil
	} else {
		return nil, err
	}
}

func (m *OracleClients) KmsManagementClientWithEndpoint(endpoint string) (*oci_kms.KmsManagementClient, error) {
	if client, err := oci_kms.NewKmsManagementClientWithConfigurationProvider(*m.KmsManagementClient().ConfigurationProvider(), endpoint); err == nil {
		if err = m.ConfigureBaseClient(&client.BaseClient); err != nil {
			return nil, err
		}
		return &client, nil
	} else {
		return nil, err
	}
}

func (m *OracleClients) IdentityDomainsClientWithEndpoint(endpoint string) (*oci_identity_domains.IdentityDomainsClient, error) {
	if client, err := oci_identity_domains.NewIdentityDomainsClientWithConfigurationProvider(*m.IdentityDomainsClient().ConfigurationProvider(), endpoint); err == nil {
		if err = m.ConfigureBaseClient(&client.BaseClient); err != nil {
			return nil, err
		}
		return &client, nil
	} else {
		return nil, err
	}
}

func getClientHostOverrides() map[string]string {
	// Get the host URL override for clients
	clientHostOverrides := make(map[string]string)
	clientHostOverridesString := utils.GetEnvSettingWithBlankDefault(globalvar.ClientHostOverridesEnv)
	if clientHostOverridesString == "" {
		return clientHostOverrides
	}

	clientHostFlags := strings.Split(clientHostOverridesString, globalvar.ColonDelimiter)
	for _, item := range clientHostFlags {
		clientNameHost := strings.Split(item, globalvar.EqualToOperatorDelimiter)
		if clientNameHost == nil || len(clientNameHost) != 2 {
			continue
		}
		clientHostOverrides[clientNameHost[0]] = clientNameHost[1]
	}
	return clientHostOverrides
}

// CreateSDKClients constructs all enabled OCI SDK clients eagerly. This
// preserves the established Terraform CLI and Resource Discovery behavior,
// including surfacing client-construction errors during provider configuration.
func CreateSDKClients(clients *OracleClients, configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) error {
	if err := prepareSDKClients(clients, configProvider, configureClient); err != nil {
		return err
	}

	clientHostOverrides := getClientHostOverrides()
	for serviceName, registration := range OracleClientRegistrationsVar.RegisteredClients {
		if !common.CheckForEnabledServices(utils.GetSDKServiceName(serviceName)) {
			continue
		}
		overrides := ServiceClientOverrides{HostUrlOverride: clientHostOverrides[serviceName]}
		client, err := registration.InitClientFn(configProvider, configureClient, overrides)
		if err != nil {
			return err
		}
		clients.SdkClientMap[serviceName] = client
	}
	return createWorkRequestClient(clients, configProvider, configureClient)
}

// CreateSDKClientsLazy records the owning provider configuration and creates
// OCI SDK clients on first use. This is used by embedded providers only; the
// standard Terraform path retains eager initialization semantics.
func CreateSDKClientsLazy(clients *OracleClients, configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) error {
	if err := prepareSDKClients(clients, configProvider, configureClient); err != nil {
		return err
	}

	clients.clientMu.Lock()
	clients.configProvider = configProvider
	clients.clientHostOverrides = getClientHostOverrides()
	clients.clientMu.Unlock()

	// The generic work request client is retained eagerly because legacy
	// resources access the exported field directly instead of using GetClient.
	return createWorkRequestClient(clients, configProvider, configureClient)
}

func prepareSDKClients(clients *OracleClients, configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) error {
	if clients == nil {
		return fmt.Errorf("cannot configure nil OracleClients")
	}
	clients.configureClient = configureClient

	if OracleClientRegistrationsVar == nil || len(OracleClientRegistrationsVar.RegisteredClients) == 0 {
		return fmt.Errorf("there are no clients to Create")
	}
	for serviceName, registration := range OracleClientRegistrationsVar.RegisteredClients {
		if registration == nil || registration.InitClientFn == nil {
			return fmt.Errorf("unable to initialize '%s' client", serviceName)
		}
	}
	clients.clientMu.Lock()
	clients.configProvider = configProvider
	if clients.SdkClientMap == nil {
		clients.SdkClientMap = make(map[string]interface{})
	}
	clients.clientMu.Unlock()
	return nil
}

func createWorkRequestClient(clients *OracleClients, configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) error {
	if common.CheckForEnabledServices(globalvar.WorkRequest) {
		workRequestClient, err := oci_work_requests.NewWorkRequestClientWithConfigurationProvider(configProvider)
		if err != nil {
			return err
		}
		err = configureClient(&workRequestClient.BaseClient)
		if err != nil {
			return err
		}
		clients.WorkRequestClient = &workRequestClient
	}
	return nil
}
func setCustomConfiguration(oClient interface {
	SetCustomClientConfiguration(config common.CustomClientConfiguration)
}) error {
	if tfresource.RealmSpecificServiceEndpointTemplateEnabled != "" {
		value, err := strconv.ParseBool(tfresource.RealmSpecificServiceEndpointTemplateEnabled)
		if err != nil {
			return err
		}
		oClient.SetCustomClientConfiguration(oci_common.CustomClientConfiguration{
			RealmSpecificServiceEndpointTemplateEnabled: oci_common.Bool(value),
		})
	}
	return nil
}

func SetDualStackEndpointEnabled(oClient *oci_common.BaseClient) error {
	if tfresource.DualStackEndpointTemplateEnabled != "" {
		dualStack, err := strconv.ParseBool(tfresource.DualStackEndpointTemplateEnabled)
		if err != nil {
			return err
		}
		oClient.EnableDualStackEndpoints(dualStack)
	}
	return nil
}
