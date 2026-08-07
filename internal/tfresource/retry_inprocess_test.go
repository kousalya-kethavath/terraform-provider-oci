// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package tfresource

import (
	"net"
	"testing"
	"time"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func TestShortRetryDurationFunctionIsOperationLocal(t *testing.T) {
	originalShort := ShortRetryTime
	originalLong := LongRetryTime
	originalConfigured := ConfiguredRetryDuration
	t.Cleanup(func() {
		ShortRetryTime = originalShort
		LongRetryTime = originalLong
		ConfiguredRetryDuration = originalConfigured
	})

	ShortRetryTime = time.Minute
	LongRetryTime = 10 * time.Minute
	ConfiguredRetryDuration = nil
	response := oci_common.OCIOperationResponse{Error: &net.DNSError{Err: "timeout", IsTimeout: true}}
	override := GetShortRetryDurationFunction(50 * time.Minute)
	if got := override(response, false, "opensearch"); got != 50*time.Minute {
		t.Fatalf("operation-local retry duration = %v", got)
	}
	if got := GetDefaultExpectedRetryDuration(response, false); got != time.Minute {
		t.Fatalf("global retry duration changed to %v", got)
	}
}
