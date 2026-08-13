# Embedding Terraform Provider OCI in process

> **Reference proof of concept:** This branch and README demonstrate one tested
> implementation of the requested public provider boundary, correctness
> guarantees, and related performance optimizations. They are provided for
> reference only. Terraform Provider OCI maintainers may choose a different
> internal design while preserving the public contract, independent
> provider-instance behavior, and selective construction of only the schemas
> and clients required by an embedded consumer.

This package exposes the narrow Go construction boundary needed by consumers
that run Terraform Provider OCI in the same process instead of starting the
Terraform CLI and an external provider plugin.

The package does not expose the implementation under `internal/`. Consumers
receive standard Terraform Plugin SDKv2 or Plugin Framework provider values and
continue to use the provider's existing schemas and CRUD callbacks.

## Public constructors

```go
import tfoci "github.com/oracle/terraform-provider-oci/oci"

sdkProvider := tfoci.Provider()
serviceProvider, err := tfoci.ProviderForResources(serviceResourceNames...)
configurationProvider := tfoci.ProviderForConfiguration()
frameworkProvider := tfoci.New()
```

- `Provider()` returns a fresh Terraform Plugin SDKv2 `*schema.Provider`.
- `ProviderForResources()` returns a fresh SDKv2 provider containing isolated
  schemas for only the requested resources and no data sources. SDKv2 resource
  and data-source constructors are registered lazily, so unrequested schemas
  are not materialized and retained by selective or configuration-only
  providers.
- `ProviderForConfiguration()` returns a fresh SDKv2 provider containing only
  the provider schema and configuration callback.
- `New()` returns a fresh Terraform Plugin Framework `provider.Provider`.
- Repeated calls return distinct provider instances.

Crossplane Provider OCI registers `ProviderForResources()` with each
service-scoped SDKv2 runtime and uses `ProviderForConfiguration()` where only
provider configuration metadata is required. The Framework constructor provides
the same public boundary for future resources that are implemented and
validated with Plugin Framework.

## In-process safety contract

An embedded process can configure more than one provider instance at the same
time. Configuration belonging to one instance must therefore not replace or
affect configuration used by another instance.

The in-process constructors enforce the following behavior:

1. Provider construction returns a fresh SDKv2 or Framework provider value;
   SDKv2 resource, data-source, and nested schema structures are isolated.
2. The OCI client-configuration callback is stored on the owning
   `OracleClients` value rather than in a package-global variable.
3. Endpoint-specific and cross-region clients use their owning provider
   instance's callback.
4. Creating an Object Storage client does not modify a process environment
   variable.
5. Regional MySQL backup operations create a separate regional client instead
   of changing the primary client's region.
6. OpenSearch retry overrides are scoped to the operation instead of changing
   package-wide retry defaults.
7. Provider aliases are registered idempotently and safely when several
   provider instances are constructed concurrently.
8. OCI SDK service clients are constructed on first use and cached by their
   owning provider instance. The generic work-request client remains eager for
   compatibility with resources that access its exported field directly.
9. Selective schema construction does not change OCI SDK service enablement;
   resources may continue to construct secondary and cross-region clients.

The existing Terraform CLI constructors and supported CLI behavior remain
available through the internal provider entry points used by the provider
binary.

## Process-global options

Some existing provider options are backed by mutable process-global state. The
in-process constructors accept their defaults but reject non-default values:

- `ignore_defined_tags`
- `realm_specific_service_endpoint_template_enabled`
- `dual_stack_endpoint_enabled`
- `disable_auto_retries`
- `retry_duration_seconds`
- `retries_config_file`

Rejecting these values prevents two embedded provider instances from silently
changing each other's behavior. The Terraform CLI path continues to support
these options because each CLI provider process has its own process state.

## Implementation map

The implementation is intentionally divided into these reviewable areas:

1. **Provider instances** — `internal/provider/provider.go` and
   `internal/provider/provider_framework.go` construct fresh providers, capture
   instance-specific configuration, and validate in-process options.
2. **Client and operation isolation** — `internal/client`, the affected
   secondary/cross-region service helpers, and `internal/tfresource/retry.go`
   prevent later client creation and retry overrides from consulting or
   changing shared process state.
3. **Public bridge** — `oci/provider.go` delegates to the internal in-process
   constructors without exposing internal implementation types.

The service files changed by this work are the known delayed-client and
operation-local override paths identified by the implementation review. They
are not intended to be a permanent exhaustive list: new paths must follow the
same instance-ownership rule.

## Commit guide

The proof of concept is split into small commits so that the public API,
correctness work, and performance work can be understood and reviewed in
logical stages.

### Required public API and correctness scope

| Commit | What it demonstrates | Why it is required |
|---|---|---|
| `Prepare providers for independent in-process use` | Returns fresh SDKv2 and Framework provider instances, isolates mutable SDKv2 schemas, scopes provider configuration to its owning instance, and validates process-global options for the embedded path. | Multiple provider instances share one process in an embedded consumer. Configuring one instance must not replace or mutate another instance's behavior. |
| `Isolate in-process OCI client and retry behavior` | Moves client configuration to the owning `OracleClients`, applies it to delayed, secondary, and cross-region clients, and prevents operation-local endpoint, region, environment, and retry changes from mutating shared process state. | Public constructors are not safe by themselves if clients created later still consult or change package-global state. |
| `Expose public OCI provider constructors` | Adds the narrow public `oci` package for SDKv2 and Plugin Framework construction without exposing internal implementation packages or types. | External Go consumers cannot import constructors under `internal/provider`. This is the requested supported integration boundary. |

These commits describe the minimum behavior needed for a safe public provider
package. The exact internal refactoring is not prescribed by this reference.

### Performance scope for production use

| Commit | What it demonstrates | Observed purpose |
|---|---|---|
| `Add selective in-process provider constructors` | Adds configuration-only and resource-selective SDKv2 constructors while retaining isolated resource schemas. | Allows a service-scoped consumer to request only the Terraform resources it reconciles. |
| `Construct OCI SDK clients lazily` | Constructs service clients on first use and caches them on the owning provider instance. | Avoids retaining OCI SDK clients for services that a provider process never uses. |
| `Optimize in-process providers with lazy schema construction` | Registers SDKv2 resource and data-source constructor functions and materializes only requested schemas; the normal Terraform path still constructs the complete provider. | Removes eager retention of unrelated Terraform schemas. The isolated selective-provider benchmark reduced retained heap by approximately 85%, and the Crossplane canary reduced service-pod working set by approximately 42–47% compared with the previous no-fork build that still eagerly constructed all SDKv2 schemas. |

The public constructors can be exposed independently of these three
optimizations. However, without selective and lazy construction, service-scoped
embedded consumers retain schemas and clients for services they do not use,
leaving a substantial portion of the measured memory benefit unrealized.
Including these behaviors in the upstream implementation would help make the
public package practical and memory-efficient for long-running embedded
consumers. The proof of concept demonstrates one implementation; Terraform
Provider OCI maintainers can determine the most appropriate internal design
for achieving equivalent behavior.

## Validation

Run the focused in-process tests with:

```sh
go test ./oci
go test ./internal/client ./internal/service/mysql ./internal/tfresource \
	./internal/provider ./internal/service-framework/vault ./internal/resourcediscovery \
	-run 'TestObjectStorageClientConstructionDoesNotMutateEnvironment|TestOracleClientsConfigureBaseClientIsolation|TestOracleClientsConfigureBaseClientRequiresCallback|TestOracleClientsConstructSDKClientsLazily|TestOracleClientsPreserveEagerTerraformInitialization|TestEndpointSpecificClientsUseOwningInstanceCallback|TestCreateDbBackupClientInRegionDoesNotMutatePrimary|TestShortRetryDurationFunctionIsOperationLocal|TestSDKv2SchemaFactoriesAreLazyAndSelective|TestBuildResourcesRejectsUnknownName|TestProviderConstructorsReturnFreshInstances|TestSelectiveInProcessProvider|TestConfigurationOnlyInProcessProvider|TestCloneSDKv2ResourceIsolatesMutableStructures|TestClonedSDKv2ResourceReturnsLazyClientInitializationErrors|TestClonedSDKv2ResourceDoesNotMaskUnrelatedPanics|TestValidateInProcessProviderConfig|TestFrameworkInProcessProviderRejectsIgnoreDefinedTags|TestInternalAndEmbeddedSDKv2SchemasRemainCompatible|TestReadReturnsLazyClientInitializationError|TestConfigureRejectsUnexpectedProviderData|TestIdentityOperationsReturnLazyClientInitializationErrors|TestRefreshProviderSchemaMapsPreservesDiscoveryEntries'
```

The complete upstream provider test suite includes tests that require an OCI
Resource Principal environment. Run that suite in its normal configured test
environment; a local machine without those credentials cannot validate those
tests.

Consumers should additionally validate:

- Two differently configured provider instances used concurrently.
- Create, read, update, delete, import, and refresh behavior.
- Cross-region and endpoint-specific resources.
- Existing Terraform CLI workflows, to confirm the public package does not
  change normal provider behavior.

## Maintenance and release checklist

Terraform Provider OCI maintainers should run this review when a future release
adds or regenerates resources, changes provider configuration, introduces new
client-construction paths, or changes retry and endpoint behavior:

1. Confirm `oci.Provider()` and `oci.New()` still return fresh instances and
   that SDKv2 resource, data-source, and nested schemas are not shared.
2. Review new provider configuration fields for mutable package-global state.
3. Search for new package-global OCI client configuration callbacks or clients
   created after initial provider configuration.
4. Review new cross-region, endpoint-specific, copy, replication, backup, and
   restore paths for instance-owned configuration.
5. Review new environment-variable writes and retry/default mutations reachable
   from provider configuration or resource CRUD.
6. Run the focused isolation tests and the configured upstream provider suite.
7. Ask embedded consumers to run schema compatibility, concurrent
   ProviderConfig, dependency-upgrade, restart, and CRUD tests before adopting
   the new Terraform Provider OCI release.

This checklist is necessary because importing the provider as a Go library
places provider instances in one long-lived process, while the Terraform CLI
normally isolates each provider process.
