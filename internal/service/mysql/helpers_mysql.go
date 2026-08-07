package mysql

import (
	"fmt"

	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"
)

func (s *MysqlMysqlBackupResourceCrud) createDbBackupClientInRegion(region string) (*oci_mysql.DbBackupsClient, error) {
	if s.Client == nil {
		return nil, fmt.Errorf("cannot create client for the region: primary MySQL backup client is nil")
	}

	dbBackupClient, err := oci_mysql.NewDbBackupsClientWithConfigurationProvider(*s.Client.ConfigurationProvider())
	if err != nil {
		return nil, fmt.Errorf("cannot create client for the region: %v", err)
	}
	if err := s.ConfigureClient(&dbBackupClient.BaseClient); err != nil {
		return nil, fmt.Errorf("cannot configure client for the region: %v", err)
	}
	dbBackupClient.SetRegion(region)
	return &dbBackupClient, nil
}
