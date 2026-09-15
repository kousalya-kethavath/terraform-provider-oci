// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package provider

import (
	"context"
	"slices"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfclient "github.com/oracle/terraform-provider-oci/internal/client"
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

func cloneSDKv2Resource(source *schema.Resource) *schema.Resource {
	if source == nil {
		return nil
	}

	result := *source
	wrapSDKv2ResourceCallbacks(&result)
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
		if importer.State != nil {
			callback := importer.State
			importer.State = func(data *schema.ResourceData, meta interface{}) (state []*schema.ResourceData, err error) {
				defer recoverLazyClientInitializationError(&err)
				return callback(data, meta)
			}
		}
		if importer.StateContext != nil {
			callback := importer.StateContext
			importer.StateContext = func(ctx context.Context, data *schema.ResourceData, meta interface{}) (state []*schema.ResourceData, err error) {
				defer recoverLazyClientInitializationError(&err)
				return callback(ctx, data, meta)
			}
		}
		result.Importer = &importer
	}
	result.Timeouts = cloneSDKv2ResourceTimeout(source.Timeouts)

	return &result
}

func wrapSDKv2ResourceCallbacks(resource *schema.Resource) {
	if resource.Create != nil {
		callback := resource.Create
		resource.Create = func(data *schema.ResourceData, meta interface{}) (err error) {
			defer recoverLazyClientInitializationError(&err)
			return callback(data, meta)
		}
	}
	if resource.Read != nil {
		callback := resource.Read
		resource.Read = func(data *schema.ResourceData, meta interface{}) (err error) {
			defer recoverLazyClientInitializationError(&err)
			return callback(data, meta)
		}
	}
	if resource.Update != nil {
		callback := resource.Update
		resource.Update = func(data *schema.ResourceData, meta interface{}) (err error) {
			defer recoverLazyClientInitializationError(&err)
			return callback(data, meta)
		}
	}
	if resource.Delete != nil {
		callback := resource.Delete
		resource.Delete = func(data *schema.ResourceData, meta interface{}) (err error) {
			defer recoverLazyClientInitializationError(&err)
			return callback(data, meta)
		}
	}
	if resource.Exists != nil {
		callback := resource.Exists
		resource.Exists = func(data *schema.ResourceData, meta interface{}) (exists bool, err error) {
			defer recoverLazyClientInitializationError(&err)
			return callback(data, meta)
		}
	}
	if resource.CustomizeDiff != nil {
		callback := resource.CustomizeDiff
		resource.CustomizeDiff = func(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) (err error) {
			defer recoverLazyClientInitializationError(&err)
			return callback(ctx, diff, meta)
		}
	}

	resource.CreateContext = wrapSDKv2ContextCallback(resource.CreateContext)
	resource.ReadContext = wrapSDKv2ContextCallback(resource.ReadContext)
	resource.UpdateContext = wrapSDKv2ContextCallback(resource.UpdateContext)
	resource.DeleteContext = wrapSDKv2ContextCallback(resource.DeleteContext)
	resource.CreateWithoutTimeout = wrapSDKv2ContextCallback(resource.CreateWithoutTimeout)
	resource.ReadWithoutTimeout = wrapSDKv2ContextCallback(resource.ReadWithoutTimeout)
	resource.UpdateWithoutTimeout = wrapSDKv2ContextCallback(resource.UpdateWithoutTimeout)
	resource.DeleteWithoutTimeout = wrapSDKv2ContextCallback(resource.DeleteWithoutTimeout)
}

func wrapSDKv2ContextCallback[T ~func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics](callback T) T {
	if callback == nil {
		return nil
	}
	return T(func(ctx context.Context, data *schema.ResourceData, meta interface{}) (diagnostics diag.Diagnostics) {
		defer func() {
			if recovered := recover(); recovered != nil {
				if err, ok := recovered.(*tfclient.LazyClientInitializationError); ok {
					diagnostics = append(diagnostics, diag.Diagnostic{Severity: diag.Error, Summary: err.Error()})
					return
				}
				panic(recovered)
			}
		}()
		return callback(ctx, data, meta)
	})
}

func recoverLazyClientInitializationError(result *error) {
	if recovered := recover(); recovered != nil {
		if err, ok := recovered.(*tfclient.LazyClientInitializationError); ok {
			*result = err
			return
		}
		panic(recovered)
	}
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
