// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_object_storage.ObjectStorageClient", &OracleClient{InitClientFn: initObjectstorageObjectStorageClient})
}

func initObjectstorageObjectStorageClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_object_storage.NewObjectStorageClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}

	SetObjectStorageClientDefaults(&client)
	err = configureClient(&client.BaseClient)

	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

// SetObjectStorageClientDefaults preserves the provider's Object Storage
// endpoint behavior for primary and operation-specific clients without
// modifying process-wide OCI SDK environment configuration.
func SetObjectStorageClientDefaults(client *oci_object_storage.ObjectStorageClient) {
	client.SetCustomClientConfiguration(oci_common.CustomClientConfiguration{
		RealmSpecificServiceEndpointTemplateEnabled: oci_common.Bool(false),
	})
}

func (m *OracleClients) ObjectStorageClient() *oci_object_storage.ObjectStorageClient {
	return m.GetClient("oci_object_storage.ObjectStorageClient").(*oci_object_storage.ObjectStorageClient)
}
