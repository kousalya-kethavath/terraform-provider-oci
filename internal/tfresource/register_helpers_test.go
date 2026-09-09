package tfresource

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestSDKv2SchemaFactoriesAreLazyAndSelective(t *testing.T) {
	resourceCalls := 0
	datasourceCalls := 0
	RegisterResource("test_resource", func() *schema.Resource {
		resourceCalls++
		return &schema.Resource{Schema: map[string]*schema.Schema{"resource": {Type: schema.TypeString}}}
	})
	RegisterDatasource("test_datasource", func() *schema.Resource {
		datasourceCalls++
		return &schema.Resource{Schema: map[string]*schema.Schema{"datasource": {Type: schema.TypeString}}}
	})

	if resourceCalls != 0 || datasourceCalls != 0 {
		t.Fatalf("registration constructed schemas: resources=%d datasources=%d", resourceCalls, datasourceCalls)
	}

	resources, err := BuildResources("test_resource", "test_resource")
	if err != nil {
		t.Fatalf("BuildResources: %v", err)
	}
	if len(resources) != 1 || resources["test_resource"] == nil {
		t.Fatalf("unexpected selective resources: %v", resources)
	}
	if resourceCalls != 1 || datasourceCalls != 0 {
		t.Fatalf("selective build calls: resources=%d datasources=%d", resourceCalls, datasourceCalls)
	}

	second, err := BuildResources("test_resource")
	if err != nil {
		t.Fatalf("second BuildResources: %v", err)
	}
	if second["test_resource"] == resources["test_resource"] {
		t.Fatal("resource factory returned a shared schema")
	}
	if resourceCalls != 2 || datasourceCalls != 0 {
		t.Fatalf("second selective build calls: resources=%d datasources=%d", resourceCalls, datasourceCalls)
	}

	datasources := BuildAllDatasources()
	if len(datasources) != 1 || datasources["test_datasource"] == nil {
		t.Fatalf("unexpected data sources: %v", datasources)
	}
	if datasourceCalls != 1 {
		t.Fatalf("all-data-source build calls=%d, want 1", datasourceCalls)
	}
}

func TestBuildResourcesRejectsUnknownName(t *testing.T) {
	if _, err := BuildResources("test_missing_resource"); err == nil {
		t.Fatal("BuildResources accepted an unknown resource")
	}
}
