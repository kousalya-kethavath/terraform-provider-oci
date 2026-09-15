// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	"fmt"
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	tf_core "github.com/oracle/terraform-provider-oci/internal/service/core"
	tf_load_balancer "github.com/oracle/terraform-provider-oci/internal/service/load_balancer"
	tf_resource "github.com/oracle/terraform-provider-oci/internal/tfresource"
)

var terraformCLISchemaRegistrationMu sync.Mutex

// terraformCLISchemaMaps preserves the legacy eager registration path while
// returning snapshots that are not changed by a later Resource Discovery
// refresh. In-process providers use provider-owned factory inventories instead.
func terraformCLISchemaMaps() (map[string]*schema.Resource, map[string]*schema.Resource) {
	terraformCLISchemaRegistrationMu.Lock()
	defer terraformCLISchemaRegistrationMu.Unlock()

	return buildTerraformCLIResourceSchemas(), buildTerraformCLIDataSourceSchemas()
}

func terraformCLIResourceSchemas() map[string]*schema.Resource {
	terraformCLISchemaRegistrationMu.Lock()
	defer terraformCLISchemaRegistrationMu.Unlock()
	return buildTerraformCLIResourceSchemas()
}

func buildTerraformCLIResourceSchemas() map[string]*schema.Resource {
	// Start with fresh maps so repeated Resource Discovery refreshes replace
	// provider-generated schemas and do not retain disabled or removed entries.
	globalvar.OciResources = make(map[string]*schema.Resource)
	registerResourcesEagerly()
	registerTerraformCLIResourceAliases()

	return copySDKv2SchemaMap(globalvar.OciResources)
}

func terraformCLIDataSourceSchemas() map[string]*schema.Resource {
	terraformCLISchemaRegistrationMu.Lock()
	defer terraformCLISchemaRegistrationMu.Unlock()
	return buildTerraformCLIDataSourceSchemas()
}

func buildTerraformCLIDataSourceSchemas() map[string]*schema.Resource {
	// Start with a fresh map so repeated Resource Discovery refreshes replace
	// provider-generated schemas and do not retain disabled or removed entries.
	globalvar.OciDatasources = make(map[string]*schema.Resource)
	registerDatasourcesEagerly()
	registerTerraformCLIDataSourceAliases()

	return copySDKv2SchemaMap(globalvar.OciDatasources)
}

func copySDKv2SchemaMap(source map[string]*schema.Resource) map[string]*schema.Resource {
	result := make(map[string]*schema.Resource, len(source))
	for name, resourceSchema := range source {
		result[name] = resourceSchema
	}
	return result
}

func registerTerraformCLIResourceAliases() {
	for name, factory := range providerAliasResourceFactories() {
		tf_resource.RegisterResource(name, factory())
	}
}

func registerTerraformCLIDataSourceAliases() {
	for name, factory := range providerAliasDatasourceFactories() {
		tf_resource.RegisterDatasource(name, factory())
	}
}

func inProcessResourceFactories() map[string]func() *schema.Resource {
	factories := generatedSDKv2ResourceFactories()
	for name, factory := range providerAliasResourceFactories() {
		factories[name] = factory
	}
	return factories
}

func inProcessDatasourceFactories() map[string]func() *schema.Resource {
	factories := generatedSDKv2DatasourceFactories()
	for name, factory := range providerAliasDatasourceFactories() {
		factories[name] = factory
	}
	return factories
}

func providerAliasResourceFactories() map[string]func() *schema.Resource {
	factories := make(map[string]func() *schema.Resource)
	if oci_common.CheckForEnabledServices(globalvar.CoreService) {
		factories["oci_core_virtual_network"] = tf_core.CoreVcnResource
	}
	if oci_common.CheckForEnabledServices(globalvar.LoadBalancerService) {
		factories["oci_load_balancer"] = tf_load_balancer.LoadBalancerLoadBalancerResource
		factories["oci_load_balancer_backendset"] = tf_load_balancer.LoadBalancerBackendSetResource
	}
	return factories
}

func providerAliasDatasourceFactories() map[string]func() *schema.Resource {
	factories := make(map[string]func() *schema.Resource)
	if oci_common.CheckForEnabledServices(globalvar.CoreService) {
		factories["oci_core_listing_resource_version"] = tf_core.CoreAppCatalogListingResourceVersionDataSource
		factories["oci_core_listing_resource_versions"] = tf_core.CoreAppCatalogListingResourceVersionsDataSource
		factories["oci_core_shape"] = tf_core.CoreShapesDataSource
		factories["oci_core_virtual_networks"] = tf_core.CoreVcnsDataSource
	}
	if oci_common.CheckForEnabledServices(globalvar.LoadBalancerService) {
		factories["oci_load_balancers"] = tf_load_balancer.LoadBalancerLoadBalancersDataSource
		factories["oci_load_balancer_backendsets"] = tf_load_balancer.LoadBalancerBackendSetsDataSource
	}
	return factories
}

func buildSelectedInProcessSDKv2Schemas(factories map[string]func() *schema.Resource, names ...string) (map[string]*schema.Resource, error) {
	resources := make(map[string]*schema.Resource, len(names))
	for _, name := range names {
		if _, exists := resources[name]; exists {
			continue
		}
		factory, exists := factories[name]
		if !exists {
			return nil, fmt.Errorf("unknown OCI SDKv2 schema %q", name)
		}
		resources[name] = cloneSDKv2Resource(factory())
	}
	return resources, nil
}

func buildInProcessSDKv2Schemas(factories map[string]func() *schema.Resource) map[string]*schema.Resource {
	resources := make(map[string]*schema.Resource, len(factories))
	for name, factory := range factories {
		resources[name] = cloneSDKv2Resource(factory())
	}
	return resources
}
