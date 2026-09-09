// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package provider

import (
	"slices"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestCloneSDKv2NilInputs(t *testing.T) {
	t.Parallel()

	if cloneSDKv2ResourceMap(nil) != nil {
		t.Fatal("nil resource map was not preserved")
	}
	if cloneSDKv2Resource(nil) != nil {
		t.Fatal("nil resource was not preserved")
	}
	if cloneSDKv2SchemaMap(nil) != nil {
		t.Fatal("nil schema map was not preserved")
	}
	if cloneSDKv2Schema(nil) != nil {
		t.Fatal("nil schema was not preserved")
	}
	if cloneSDKv2ResourceTimeout(nil) != nil {
		t.Fatal("nil resource timeout was not preserved")
	}
	if cloneDuration(nil) != nil {
		t.Fatal("nil duration was not preserved")
	}
}

func TestCloneSDKv2ResourceMapIsolatesEntries(t *testing.T) {
	t.Parallel()

	source := map[string]*schema.Resource{
		"resource": {
			Schema: map[string]*schema.Schema{
				"field": {Type: schema.TypeString},
			},
		},
		"nil": nil,
	}

	cloned := cloneSDKv2ResourceMap(source)
	if len(cloned) != len(source) {
		t.Fatalf("cloned resource map length = %d, want %d", len(cloned), len(source))
	}
	if cloned["resource"] == source["resource"] {
		t.Fatal("resource map retained a source resource pointer")
	}
	if cloned["nil"] != nil {
		t.Fatal("resource map did not preserve a nil entry")
	}

	cloned["resource"].Schema["field"].Optional = true
	cloned["additional"] = &schema.Resource{}
	if source["resource"].Schema["field"].Optional {
		t.Fatal("mutating a cloned resource affected the source resource")
	}
	if _, exists := source["additional"]; exists {
		t.Fatal("adding an entry to the cloned map affected the source map")
	}
}

func TestCloneSDKv2ResourceIsolatesMutableStructures(t *testing.T) {
	t.Parallel()

	createTimeout := time.Minute
	readTimeout := 2 * time.Minute
	updateTimeout := 3 * time.Minute
	deleteTimeout := 4 * time.Minute
	defaultTimeout := 5 * time.Minute

	sharedResourceSchema := map[string]*schema.Schema{
		"dynamic": {Type: schema.TypeString},
	}
	sharedIdentitySchema := map[string]*schema.Schema{
		"identity": {Type: schema.TypeString},
	}
	source := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"field": {
				Type:          schema.TypeString,
				ComputedWhen:  []string{"computed"},
				ConflictsWith: []string{"other"},
				ExactlyOneOf:  []string{"field", "other"},
				AtLeastOneOf:  []string{"field", "other"},
				RequiredWith:  []string{"other"},
			},
		},
		SchemaFunc: func() map[string]*schema.Schema {
			return sharedResourceSchema
		},
		StateUpgraders:                 []schema.StateUpgrader{{Version: 1}},
		ValidateRawResourceConfigFuncs: make([]schema.ValidateRawResourceConfigFunc, 1),
		Identity: &schema.ResourceIdentity{
			SchemaFunc: func() map[string]*schema.Schema {
				return sharedIdentitySchema
			},
			IdentityUpgraders: []schema.IdentityUpgrader{{Version: 1}},
		},
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create:  &createTimeout,
			Read:    &readTimeout,
			Update:  &updateTimeout,
			Delete:  &deleteTimeout,
			Default: &defaultTimeout,
		},
	}

	cloned := cloneSDKv2Resource(source)
	if cloned == source || cloned.Schema["field"] == source.Schema["field"] {
		t.Fatal("resource schema was not cloned")
	}
	if cloned.Identity == source.Identity {
		t.Fatal("resource identity was not cloned")
	}
	if cloned.Importer == source.Importer {
		t.Fatal("resource importer was not cloned")
	}
	if cloned.Timeouts == source.Timeouts {
		t.Fatal("resource timeouts were not cloned")
	}

	assertStringSlicesIsolated(t, "ComputedWhen", source.Schema["field"].ComputedWhen, cloned.Schema["field"].ComputedWhen)
	assertStringSlicesIsolated(t, "ConflictsWith", source.Schema["field"].ConflictsWith, cloned.Schema["field"].ConflictsWith)
	assertStringSlicesIsolated(t, "ExactlyOneOf", source.Schema["field"].ExactlyOneOf, cloned.Schema["field"].ExactlyOneOf)
	assertStringSlicesIsolated(t, "AtLeastOneOf", source.Schema["field"].AtLeastOneOf, cloned.Schema["field"].AtLeastOneOf)
	assertStringSlicesIsolated(t, "RequiredWith", source.Schema["field"].RequiredWith, cloned.Schema["field"].RequiredWith)

	if &cloned.StateUpgraders[0] == &source.StateUpgraders[0] {
		t.Fatal("state upgrader slice was not cloned")
	}
	if &cloned.ValidateRawResourceConfigFuncs[0] == &source.ValidateRawResourceConfigFuncs[0] {
		t.Fatal("raw resource validator slice was not cloned")
	}
	if &cloned.Identity.IdentityUpgraders[0] == &source.Identity.IdentityUpgraders[0] {
		t.Fatal("identity upgrader slice was not cloned")
	}

	assertSchemaFuncIsolated(t, "resource", sharedResourceSchema, cloned.SchemaFunc)
	assertSchemaFuncIsolated(t, "identity", sharedIdentitySchema, cloned.Identity.SchemaFunc)
	assertTimeoutsIsolated(t, source.Timeouts, cloned.Timeouts)

	cloned.Schema["field"].ConflictsWith[0] = "changed"
	cloned.StateUpgraders[0].Version = 2
	cloned.Identity.IdentityUpgraders[0].Version = 2
	if source.Schema["field"].ConflictsWith[0] != "other" {
		t.Fatal("mutating a cloned schema constraint affected the source")
	}
	if source.StateUpgraders[0].Version != 1 {
		t.Fatal("mutating a cloned state upgrader affected the source")
	}
	if source.Identity.IdentityUpgraders[0].Version != 1 {
		t.Fatal("mutating a cloned identity upgrader affected the source")
	}
}

func TestCloneSDKv2SchemaRecursivelyClonesElements(t *testing.T) {
	t.Parallel()

	source := map[string]*schema.Schema{
		"schema_element": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type:          schema.TypeString,
				ConflictsWith: []string{"other"},
			},
		},
		"resource_element": {
			Type: schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"nested": {Type: schema.TypeString},
				},
			},
		},
		"value_element": {
			Type: schema.TypeList,
			Elem: schema.TypeString,
		},
	}

	cloned := cloneSDKv2SchemaMap(source)
	clonedSchemaElement := cloned["schema_element"].Elem.(*schema.Schema)
	sourceSchemaElement := source["schema_element"].Elem.(*schema.Schema)
	if clonedSchemaElement == sourceSchemaElement {
		t.Fatal("nested schema element was not cloned")
	}

	clonedResourceElement := cloned["resource_element"].Elem.(*schema.Resource)
	sourceResourceElement := source["resource_element"].Elem.(*schema.Resource)
	if clonedResourceElement == sourceResourceElement || clonedResourceElement.Schema["nested"] == sourceResourceElement.Schema["nested"] {
		t.Fatal("nested resource element was not cloned recursively")
	}

	clonedSchemaElement.ConflictsWith[0] = "changed"
	clonedResourceElement.Schema["nested"].Optional = true
	if sourceSchemaElement.ConflictsWith[0] != "other" {
		t.Fatal("mutating a cloned nested schema affected the source")
	}
	if sourceResourceElement.Schema["nested"].Optional {
		t.Fatal("mutating a cloned nested resource affected the source")
	}
	if cloned["value_element"].Elem != source["value_element"].Elem {
		t.Fatal("immutable schema value element was not preserved")
	}
}

func assertStringSlicesIsolated(t *testing.T, name string, source, cloned []string) {
	t.Helper()
	if !slices.Equal(source, cloned) {
		t.Fatalf("%s = %v, want %v", name, cloned, source)
	}
	if len(source) != 0 && &source[0] == &cloned[0] {
		t.Fatalf("%s slice was not cloned", name)
	}
}

func assertSchemaFuncIsolated(t *testing.T, name string, source map[string]*schema.Schema, cloned func() map[string]*schema.Schema) {
	t.Helper()
	first := cloned()
	second := cloned()
	if len(first) != len(source) || len(second) != len(source) {
		t.Fatalf("%s SchemaFunc map lengths = (%d, %d), want %d", name, len(first), len(second), len(source))
	}
	for field, sourceSchema := range source {
		if first[field] == sourceSchema || second[field] == sourceSchema || first[field] == second[field] {
			t.Fatalf("%s SchemaFunc returned a shared schema for %q", name, field)
		}
		if first[field].Type != sourceSchema.Type || second[field].Type != sourceSchema.Type {
			t.Fatalf("%s SchemaFunc changed the schema type for %q", name, field)
		}
	}
}

func assertTimeoutsIsolated(t *testing.T, source, cloned *schema.ResourceTimeout) {
	t.Helper()
	tests := map[string]struct {
		source *time.Duration
		cloned *time.Duration
	}{
		"create":  {source: source.Create, cloned: cloned.Create},
		"read":    {source: source.Read, cloned: cloned.Read},
		"update":  {source: source.Update, cloned: cloned.Update},
		"delete":  {source: source.Delete, cloned: cloned.Delete},
		"default": {source: source.Default, cloned: cloned.Default},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.source == tt.cloned {
				t.Fatal("timeout duration pointer was not cloned")
			}
			if *tt.cloned != *tt.source {
				t.Fatalf("cloned timeout = %s, want %s", *tt.cloned, *tt.source)
			}
			original := *tt.source
			*tt.cloned += time.Second
			if *tt.source != original {
				t.Fatal("mutating a cloned timeout affected the source")
			}
		})
	}
}
