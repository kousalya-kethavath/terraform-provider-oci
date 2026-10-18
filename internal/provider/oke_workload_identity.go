// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package provider

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

// withOKETokenPath selects one explicit token source for this configuration.
// CLI callers rotate the file; embedded callers may supply a Go callback.
func (o InProcessOptions) withOKETokenPath(authType, tokenPath string) (InProcessOptions, error) {
	if err := o.validate(authType); err != nil {
		return o, err
	}
	if tokenPath == "" {
		return o, nil
	}
	if !strings.EqualFold(authType, globalvar.AuthOKEWorkloadIdentity) {
		return o, fmt.Errorf("oke_workload_identity_token_path requires auth OKEWorkloadIdentity")
	}
	if o.OKEWorkloadIdentityTokenProvider != nil {
		return o, fmt.Errorf("configure either an OKE token callback or oke_workload_identity_token_path, not both")
	}
	if !filepath.IsAbs(tokenPath) {
		return o, fmt.Errorf("oke_workload_identity_token_path must be an absolute path")
	}
	o.OKEWorkloadIdentityTokenProvider = okeTokenFile{path: tokenPath}
	return o, nil
}

type okeTokenFile struct{ path string }

// Reopen on every SDK request so atomic file replacement and projected-volume
// symlink rotation are observed. Do not retain a descriptor or the previous JWT.
func (p okeTokenFile) ServiceAccountToken() (string, error) {
	const maxTokenBytes = 128 * 1024
	info, err := os.Stat(p.path)
	if err != nil || !info.Mode().IsRegular() || info.Size() > maxTokenBytes {
		return "", fmt.Errorf("cannot read the configured OKE ServiceAccount token file")
	}
	f, err := os.Open(p.path)
	if err != nil {
		return "", fmt.Errorf("cannot read the configured OKE ServiceAccount token file")
	}
	defer f.Close()
	data, err := io.ReadAll(io.LimitReader(f, maxTokenBytes+1))
	if err != nil || len(data) > maxTokenBytes {
		return "", fmt.Errorf("cannot read the configured OKE ServiceAccount token file")
	}
	token := strings.TrimSpace(string(data))
	if token == "" {
		return "", fmt.Errorf("the configured OKE ServiceAccount token file is empty")
	}
	return token, nil
}
