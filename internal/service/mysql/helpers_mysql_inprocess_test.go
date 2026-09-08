// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package mysql

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"strings"
	"testing"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"
	"github.com/oracle/terraform-provider-oci/internal/client"
)

func TestCreateDbBackupClientInRegionDoesNotMutatePrimary(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}
	primary, err := oci_mysql.NewDbBackupsClientWithConfigurationProvider(mysqlTestConfiguration{privateKey: privateKey})
	if err != nil {
		t.Fatalf("construct primary client: %v", err)
	}
	primary.UserAgent = "primary"
	primaryHost := primary.Host
	crud := &MysqlMysqlBackupResourceCrud{
		Client: &primary,
		ConfigureClient: client.ConfigureClient(func(base *oci_common.BaseClient) error {
			base.UserAgent = "regional"
			return nil
		}),
	}
	regional, err := crud.createDbBackupClientInRegion("us-phoenix-1")
	if err != nil {
		t.Fatalf("construct regional client: %v", err)
	}
	if regional == crud.Client {
		t.Fatal("regional operation reused the cached primary client")
	}
	if crud.Client.UserAgent != "primary" || crud.Client.Host != primaryHost {
		t.Fatalf("primary client was mutated: userAgent=%q host=%q", crud.Client.UserAgent, crud.Client.Host)
	}
	if regional.UserAgent != "regional" || regional.Host == primaryHost {
		t.Fatalf("regional client was not independently configured: userAgent=%q host=%q", regional.UserAgent, regional.Host)
	}
}

func TestCreateDbBackupClientInRegionReportsClientAndRegion(t *testing.T) {
	const region = "us-phoenix-1"
	_, err := (&MysqlMysqlBackupResourceCrud{}).createDbBackupClientInRegion(region)
	if err == nil {
		t.Fatal("createDbBackupClientInRegion returned nil with no primary client")
	}
	if !strings.Contains(err.Error(), "MySQL backup client") || !strings.Contains(err.Error(), region) {
		t.Fatalf("error %q does not identify the client and region", err)
	}
}

func TestCreateDbBackupClientInRegionWrapsConfigurationError(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}
	primary, err := oci_mysql.NewDbBackupsClientWithConfigurationProvider(mysqlTestConfiguration{privateKey: privateKey})
	if err != nil {
		t.Fatalf("construct primary client: %v", err)
	}
	want := errors.New("configure regional client")
	crud := &MysqlMysqlBackupResourceCrud{
		Client: &primary,
		ConfigureClient: client.ConfigureClient(func(*oci_common.BaseClient) error {
			return want
		}),
	}

	_, err = crud.createDbBackupClientInRegion("us-phoenix-1")
	if !errors.Is(err, want) {
		t.Fatalf("configuration error = %v, want wrapped %v", err, want)
	}
}

type mysqlTestConfiguration struct{ privateKey *rsa.PrivateKey }

func (c mysqlTestConfiguration) PrivateRSAKey() (*rsa.PrivateKey, error) { return c.privateKey, nil }
func (mysqlTestConfiguration) KeyID() (string, error)                    { return "tenancy/user/fingerprint", nil }
func (mysqlTestConfiguration) TenancyOCID() (string, error)              { return "tenancy", nil }
func (mysqlTestConfiguration) UserOCID() (string, error)                 { return "user", nil }
func (mysqlTestConfiguration) KeyFingerprint() (string, error)           { return "fingerprint", nil }
func (mysqlTestConfiguration) Region() (string, error)                   { return "us-ashburn-1", nil }
func (mysqlTestConfiguration) AuthType() (oci_common.AuthConfig, error) {
	return oci_common.AuthConfig{AuthType: oci_common.UserPrincipal}, nil
}
