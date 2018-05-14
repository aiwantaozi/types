package client

const (
	MySQLConfigType          = "mySQLConfig"
	MySQLConfigFieldDatabase = "database"
	MySQLConfigFieldEndpoint = "endpoint"
	MySQLConfigFieldPassword = "password"
	MySQLConfigFieldTable    = "table"
	MySQLConfigFieldUsername = "username"
)

type MySQLConfig struct {
	Database string `json:"database,omitempty" yaml:"database,omitempty"`
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	Table    string `json:"table,omitempty" yaml:"table,omitempty"`
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
