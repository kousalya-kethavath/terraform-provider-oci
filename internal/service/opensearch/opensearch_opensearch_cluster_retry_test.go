// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package opensearch

import (
	"net"
	"testing"
	"time"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TestOpensearchWorkRequestRetryUsesOperationDuration(t *testing.T) {
	response := oci_common.OCIOperationResponse{
		Error: &net.DNSError{Err: "timeout", IsTimeout: true},
	}

	if retry := opensearchClusterWorkRequestShouldRetryFunc(time.Minute); !retry(response) {
		t.Fatal("default work-request policy did not retry a transient error")
	}

	noRetryDuration := tfresource.GetShortRetryDurationFunction(0)
	if retry := opensearchClusterWorkRequestShouldRetryFunc(time.Minute, noRetryDuration); retry(response) {
		t.Fatal("work-request policy ignored the operation-local retry duration")
	}
}
