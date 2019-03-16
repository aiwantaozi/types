package client

const (
	ConnectionPoolSettingsType      = "connectionPoolSettings"
	ConnectionPoolSettingsFieldHttp = "http"
	ConnectionPoolSettingsFieldTcp  = "tcp"
)

type ConnectionPoolSettings struct {
	Http *ConnectionPoolSettings_HTTPSettings `json:"http,omitempty" yaml:"http,omitempty"`
	Tcp  *ConnectionPoolSettings_TCPSettings  `json:"tcp,omitempty" yaml:"tcp,omitempty"`
}
