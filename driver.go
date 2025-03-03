package bigquery

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

const (
	//ConnectionStringEnvKey Connection String Environment Key
	ConnectionStringEnvKey = "BIGQUERY_CONNECTION_STRING"
)

//Driver Database Driver
type Driver struct {
}

//Open Open a database driver connection
func (d *Driver) Open(connectionString string) (c driver.Conn, err error) {
	connector, err := d.OpenConnector(connectionString)
	if err != nil {
		return
	}
	return connector.Connect(context.Background())
}

//OpenConnector Open a database driver connection
func (d *Driver) OpenConnector(connectionString string) (c driver.Connector, err error) {
	if _, err := ConfigFromConnString(connectionString); err != nil { // validates connection string
		return nil, err
	}
	return NewConnector(connectionString), nil
}

func init() {
	sql.Register("bigquery", &Driver{})
}
