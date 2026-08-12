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
frameworkProvider := tfoci.New()
```

- `Provider()` returns a fresh Terraform Plugin SDKv2 `*schema.Provider`.
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

## Validation

Run the focused contract tests with:

```sh
go test ./oci
go test ./internal/client ./internal/service/mysql ./internal/tfresource \
  ./internal/provider \
  -run 'TestObjectStorageClientConstructionDoesNotMutateEnvironment|TestOracleClientsConfigureBaseClientIsolation|TestOracleClientsConfigureBaseClientRequiresCallback|TestEndpointSpecificClientsUseOwningInstanceCallback|TestCreateDbBackupClientInRegionDoesNotMutatePrimary|TestShortRetryDurationFunctionIsOperationLocal|TestProviderConstructorsReturnFreshInstances|TestCloneSDKv2ResourceIsolatesMutableStructures|TestValidateInProcessProviderConfig|TestFrameworkInProcessProviderRejectsIgnoreDefinedTags|TestInternalAndEmbeddedSDKv2SchemasRemainCompatible'
```

Before release, run the normal configured provider unit, integration, and
acceptance suites, including existing Terraform CLI workflows. An embedded
consumer should separately validate concurrent configurations, CRUD, import,
refresh, and representative endpoint-specific and cross-region resources.

## Maintenance checklist

When provider configuration, client construction, retry behavior, or service
resources change, review new code for package-global mutable state and verify
that any later-created client remains bound to its owning provider instance.
Run the focused isolation tests and the configured upstream provider suite.
