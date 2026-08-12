// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package provider

import (
	"fmt"
	"slices"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// cloneSDKv2ResourceMap returns an isolated copy of an SDKv2 resource map.
// CRUD callbacks and other function values are immutable and may be shared,
// while resource and schema structures that consumers can modify are copied.
func cloneSDKv2ResourceMap(source map[string]*schema.Resource) map[string]*schema.Resource {
	if source == nil {
		return nil
	}

	result := make(map[string]*schema.Resource, len(source))
	for name, resource := range source {
		result[name] = cloneSDKv2Resource(resource)
	}
	return result
}

// cloneSelectedSDKv2Resources returns isolated copies of only the named
// resources. It avoids cloning the complete provider schema for consumers that
// run a known subset of Terraform resources in process.
func cloneSelectedSDKv2Resources(source map[string]*schema.Resource, names []string) (map[string]*schema.Resource, error) {
	result := make(map[string]*schema.Resource, len(names))
	for _, name := range names {
		if _, exists := result[name]; exists {
			continue
		}
		resource, ok := source[name]
		if !ok {
			return nil, fmt.Errorf("unknown Terraform Provider OCI resource %q", name)
		}
		result[name] = cloneSDKv2Resource(resource)
	}
	return result, nil
}

func cloneSDKv2Resource(source *schema.Resource) *schema.Resource {
	if source == nil {
		return nil
	}

	result := *source
	result.Schema = cloneSDKv2SchemaMap(source.Schema)
	if source.SchemaFunc != nil {
		schemaFunc := source.SchemaFunc
		result.SchemaFunc = func() map[string]*schema.Schema {
			return cloneSDKv2SchemaMap(schemaFunc())
		}
	}
	result.StateUpgraders = slices.Clone(source.StateUpgraders)
	result.ValidateRawResourceConfigFuncs = slices.Clone(source.ValidateRawResourceConfigFuncs)

	if source.Identity != nil {
		identity := *source.Identity
		identity.IdentityUpgraders = slices.Clone(source.Identity.IdentityUpgraders)
		if source.Identity.SchemaFunc != nil {
			identitySchemaFunc := source.Identity.SchemaFunc
			identity.SchemaFunc = func() map[string]*schema.Schema {
				return cloneSDKv2SchemaMap(identitySchemaFunc())
			}
		}
		result.Identity = &identity
	}
	if source.Importer != nil {
		importer := *source.Importer
		result.Importer = &importer
	}
	result.Timeouts = cloneSDKv2ResourceTimeout(source.Timeouts)

	return &result
}

func cloneSDKv2SchemaMap(source map[string]*schema.Schema) map[string]*schema.Schema {
	if source == nil {
		return nil
	}

	result := make(map[string]*schema.Schema, len(source))
	for name, field := range source {
		result[name] = cloneSDKv2Schema(field)
	}
	return result
}

func cloneSDKv2Schema(source *schema.Schema) *schema.Schema {
	if source == nil {
		return nil
	}

	result := *source
	result.ComputedWhen = slices.Clone(source.ComputedWhen)
	result.ConflictsWith = slices.Clone(source.ConflictsWith)
	result.ExactlyOneOf = slices.Clone(source.ExactlyOneOf)
	result.AtLeastOneOf = slices.Clone(source.AtLeastOneOf)
	result.RequiredWith = slices.Clone(source.RequiredWith)

	switch element := source.Elem.(type) {
	case *schema.Schema:
		result.Elem = cloneSDKv2Schema(element)
	case *schema.Resource:
		result.Elem = cloneSDKv2Resource(element)
	}

	return &result
}

func cloneSDKv2ResourceTimeout(source *schema.ResourceTimeout) *schema.ResourceTimeout {
	if source == nil {
		return nil
	}

	result := *source
	result.Create = cloneDuration(source.Create)
	result.Read = cloneDuration(source.Read)
	result.Update = cloneDuration(source.Update)
	result.Delete = cloneDuration(source.Delete)
	result.Default = cloneDuration(source.Default)
	return &result
}

func cloneDuration(source *time.Duration) *time.Duration {
	if source == nil {
		return nil
	}

	result := *source
	return &result
}
