package tfresource

import (
	"fmt"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

type sdkV2SchemaFactory func() *schema.Resource

var sdkV2SchemaFactories = struct {
	sync.RWMutex
	resources   map[string]sdkV2SchemaFactory
	datasources map[string]sdkV2SchemaFactory
}{
	resources:   make(map[string]sdkV2SchemaFactory),
	datasources: make(map[string]sdkV2SchemaFactory),
}

// RegisterResource registers a lazy SDKv2 resource-schema constructor.
// Registration occurs during package initialization; provider construction
// invokes only the constructors needed by that provider instance.
func RegisterResource(name string, factory func() *schema.Resource) {
	sdkV2SchemaFactories.Lock()
	defer sdkV2SchemaFactories.Unlock()
	sdkV2SchemaFactories.resources[name] = factory
}

// RegisterDatasource registers a lazy SDKv2 data-source schema constructor.
func RegisterDatasource(name string, factory func() *schema.Resource) {
	sdkV2SchemaFactories.Lock()
	defer sdkV2SchemaFactories.Unlock()
	sdkV2SchemaFactories.datasources[name] = factory
}

// BuildResources constructs fresh schemas for only the named SDKv2 resources.
func BuildResources(names ...string) (map[string]*schema.Resource, error) {
	factories, err := selectedSDKv2SchemaFactories(sdkV2SchemaFactories.resources, names)
	if err != nil {
		return nil, err
	}
	return buildSDKv2Schemas(factories), nil
}

// BuildAllResources constructs fresh schemas for all registered SDKv2 resources.
func BuildAllResources() map[string]*schema.Resource {
	return buildSDKv2Schemas(allSDKv2SchemaFactories(sdkV2SchemaFactories.resources))
}

// BuildAllDatasources constructs fresh schemas for all registered SDKv2 data sources.
func BuildAllDatasources() map[string]*schema.Resource {
	return buildSDKv2Schemas(allSDKv2SchemaFactories(sdkV2SchemaFactories.datasources))
}

func selectedSDKv2SchemaFactories(source map[string]sdkV2SchemaFactory, names []string) (map[string]sdkV2SchemaFactory, error) {
	sdkV2SchemaFactories.RLock()
	defer sdkV2SchemaFactories.RUnlock()

	selected := make(map[string]sdkV2SchemaFactory, len(names))
	for _, name := range names {
		factory, ok := source[name]
		if !ok {
			return nil, fmt.Errorf("unknown OCI SDKv2 schema %q", name)
		}
		selected[name] = factory
	}
	return selected, nil
}

func allSDKv2SchemaFactories(source map[string]sdkV2SchemaFactory) map[string]sdkV2SchemaFactory {
	sdkV2SchemaFactories.RLock()
	defer sdkV2SchemaFactories.RUnlock()

	result := make(map[string]sdkV2SchemaFactory, len(source))
	for name, factory := range source {
		result[name] = factory
	}
	return result
}

func buildSDKv2Schemas(factories map[string]sdkV2SchemaFactory) map[string]*schema.Resource {
	result := make(map[string]*schema.Resource, len(factories))
	for name, factory := range factories {
		result[name] = factory()
	}
	return result
}

func RegisterFrameworkDatasource(ds func() datasource.DataSource) {
	if globalvar.OciFrameworkDataSources == nil {
		globalvar.OciFrameworkDataSources = make([]func() datasource.DataSource, 0)
	}
	globalvar.OciFrameworkDataSources = append(globalvar.OciFrameworkDataSources, ds)
}

func RegisterFrameworkResource(ds func() resource.Resource) {
	if globalvar.OciFrameworkResources == nil {
		globalvar.OciFrameworkResources = make([]func() resource.Resource, 0)
	}
	globalvar.OciFrameworkResources = append(globalvar.OciFrameworkResources, ds)
}
