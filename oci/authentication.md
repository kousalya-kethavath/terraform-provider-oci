# Supplying OKE workload tokens to an embedded provider

`ProviderForConfigurationWithOptions` and `NewWithOptions` accept instance-local
authentication dependencies. Existing constructors and zero-value options keep
their default authentication behavior.

```go
options := oci.InProcessOptions{
    OKEWorkloadIdentityTokenProvider: renewableTokenProvider,
}
sdkv2 := oci.ProviderForConfigurationWithOptions(options)
framework := oci.NewWithOptions(options)
```

The supplied value implements `auth.ServiceAccountTokenProvider` from OCI Go SDK.
Configure `auth = "OKEWorkloadIdentity"`; other authentication modes reject an
injected OKE supplier. These are Go embedding options, not Terraform HCL fields.
Kubernetes lookup, TokenRequest authorization, account scope checks and token
renewal belong to the host. Do not put Kubernetes JWTs into Terraform config.

The supplier must remain valid after Configure returns, support concurrent calls,
and return renewed tokens throughout the lifetime of clients configured with it.
Do not capture a short-lived Configure/reconciliation context. SDK OCI credential
renewal calls the supplier again. A static token is not sufficient for long-lived
clients. Supplier failures and empty tokens are deliberately converted to safe
diagnostics without exposing the underlying error or returned token.

Each provider instance retains its own supplier. No globals or environment
mutations are used to select tenant tokens. For custom suppliers, an explicit
provider region overrides the OKE SDK provider's API region through a wrapper;
claim access and refresh behavior are retained. The usual OKE runtime environment
is still required to construct the SDK exchange provider, including its region,
version, Kubernetes host and CA. With no supplier, default region behavior is
unchanged.

A custom supplier is the exclusive authentication source: this path bypasses
the composing provider's fallback to ambient credentials from configuration files
or environment variables. SDK exchange errors are sanitized before returning
them to the host because remote error bodies may contain sensitive material.

Keep the supplier identity in the embedding host's client cache key. Account
recreation or selection changes must invalidate cached clients, while ordinary
token renewal should not force all OCI clients to be rebuilt. Keep credentials
available until resource deletion and asynchronous operations complete.

This extension does not request Kubernetes tokens or grant tenant access by
itself. Live Enhanced OKE exchange, tenant IAM denial tests, and renewal across
both Kubernetes and OCI credential lifetimes must be validated by the host.

## Terraform CLI token files

The normal SDKv2 and Plugin Framework providers also accept an explicit token
file. This does not require in-process execution:

```hcl
provider "oci" {
  auth                             = "OKEWorkloadIdentity"
  region                           = "us-ashburn-1"
  oke_workload_identity_token_path = "/run/oci/tenant-a.jwt"
}
```

The path must be absolute. A host such as Crossplane obtains the selected
ServiceAccount JWT and atomically replaces this file before expiry. The provider
reopens it whenever the OCI SDK needs a JWT; it does not request Kubernetes tokens
or write to the file. Empty, missing and oversized files fail closed. Configure
neither a file nor a callback to retain the existing pod-token behavior; setting
both is an error. This OKE-specific option does not change the existing non-OKE
`workload_identity_token_path` federation option. The standard OKE runtime
requirements still apply.

Only the path is provider configuration. Do not embed the JWT in HCL. Use private
file permissions and avoid shared tenant volumes. The SDK caches and renews OCI
credentials independently from the host's JWT rotation. A rotated file alone is
not proof of RPST renewal; test both credential lifetimes.
