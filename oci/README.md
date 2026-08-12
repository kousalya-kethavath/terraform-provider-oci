# Embedding Terraform Provider OCI in process

> **Reference proof of concept:** This package demonstrates a supported public
> construction boundary and the minimum instance-isolation behavior needed by a
> long-running Go consumer. Terraform Provider OCI maintainers may choose a
> different internal implementation while preserving that contract.

`oci` exposes the narrow Go boundary for consumers that run Terraform Provider
OCI in the same process rather than starting Terraform CLI and an external
provider plugin. It returns standard Terraform Plugin SDKv2 or Plugin Framework
provider values and does not expose packages under `internal/`.

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
  schemas for only the named resources and no data-source schemas.
- `ProviderForConfiguration()` returns a fresh SDKv2 provider that validates
  and configures credentials without retaining resource or data-source schemas.
- `New()` returns a fresh Terraform Plugin Framework `provider.Provider`.
- Repeated calls return independently configurable provider instances.

Existing Terraform CLI construction and supported CLI behavior remain available
through the internal provider entry points used by the provider binary.

## In-process safety contract

An embedded consumer can configure multiple provider instances concurrently.
Configuration from one instance must not replace or affect configuration used
by another instance. The in-process constructors therefore guarantee that:

1. SDKv2 and Framework construction returns a fresh provider value; SDKv2
   resource, data-source, and nested schema structures are isolated.
2. OCI client configuration is owned by the specific `OracleClients` instance,
   including delayed, endpoint-specific, secondary, and cross-region clients.
3. Object Storage client construction does not modify process environment.
4. Regional MySQL backup operations create a separate regional client rather
   than changing the primary client's region.
5. OpenSearch retry overrides are local to an operation rather than changing
   package-wide defaults.
6. Provider aliases are registered once if several provider instances are
   constructed concurrently.

## Process-global options

Some existing provider options are backed by mutable process-global state. The
in-process constructors accept their defaults but reject non-default values for:

- `ignore_defined_tags`
- `realm_specific_service_endpoint_template_enabled`
- `dual_stack_endpoint_enabled`
- `disable_auto_retries`
- `retry_duration_seconds`
- `retries_config_file`

The normal Terraform CLI path continues to support these options because each
CLI provider has its own process state.

## Selective and lazy construction

This optional embedded-consumer optimization avoids retaining the schemas and
OCI SDK clients for services a long-running consumer will not use:

- SDKv2 resource and data-source registrations are factories. `Provider()`
  still materializes the complete normal Terraform provider, while selective
  and configuration-only providers materialize only what they require.
- `ProviderForResources()` rejects an unknown resource name rather than
  silently constructing an incomplete provider.
- OCI SDK service clients are constructed on first use and cached on the
  owning provider instance. The generic work-request client remains eager for
  compatibility with resources that access its exported field directly.
- Selective schema construction does not change OCI SDK service enablement;
  a resource can still construct its secondary or cross-region clients.

The performance behavior is intentionally separate from the public API and
isolation contract. It can be reviewed and evolved independently, provided the
complete Terraform CLI path and public constructor behavior remain compatible.

## Validation

Run the focused contract tests and the selective/lazy-construction tests with:

```sh
go test ./oci
go test ./internal/client ./internal/service/mysql ./internal/tfresource \
  ./internal/provider \
  -run 'TestObjectStorageClientConstructionDoesNotMutateEnvironment|TestOracleClientsConfigureBaseClientIsolation|TestOracleClientsConfigureBaseClientRequiresCallback|TestOracleClientsConstructSDKClientsLazily|TestEndpointSpecificClientsUseOwningInstanceCallback|TestCreateDbBackupClientInRegionDoesNotMutatePrimary|TestShortRetryDurationFunctionIsOperationLocal|TestSDKv2SchemaFactoriesAreLazyAndSelective|TestBuildResourcesRejectsUnknownName|TestProviderConstructorsReturnFreshInstances|TestSelectiveInProcessProvider|TestConfigurationOnlyInProcessProvider|TestCloneSDKv2ResourceIsolatesMutableStructures|TestValidateInProcessProviderConfig|TestFrameworkInProcessProviderRejectsIgnoreDefinedTags|TestInternalAndEmbeddedSDKv2SchemasRemainCompatible'
```

Before release, run the normal configured provider unit, integration, and
acceptance suites, including existing Terraform CLI workflows. An embedded
consumer should separately validate concurrent configurations, CRUD, import,
refresh, and representative endpoint-specific and cross-region resources.

## Maintenance checklist

When provider configuration, client construction, retry behavior, or service
resources change, review new code for package-global mutable state and verify
that any later-created client remains bound to its owning provider instance.
For generated SDKv2 registrations, also verify that each registration supplies
a factory and that full-provider construction still exposes every resource and
data source. Run the focused tests and the configured upstream provider suite.
