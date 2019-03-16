package client

const (
	ConnectionPoolSettings_TCPSettingsType                = "connectionPoolSettings_TCPSettings"
	ConnectionPoolSettings_TCPSettingsFieldConnectTimeout = "connect_timeout"
	ConnectionPoolSettings_TCPSettingsFieldMaxConnections = "max_connections"
	ConnectionPoolSettings_TCPSettingsFieldTcpKeepalive   = "tcp_keepalive"
)

type ConnectionPoolSettings_TCPSettings struct {
	ConnectTimeout *Duration                                        `json:"connect_timeout,omitempty" yaml:"connect_timeout,omitempty"`
	MaxConnections int64                                            `json:"max_connections,omitempty" yaml:"max_connections,omitempty"`
	TcpKeepalive   *ConnectionPoolSettings_TCPSettings_TcpKeepalive `json:"tcp_keepalive,omitempty" yaml:"tcp_keepalive,omitempty"`
}
