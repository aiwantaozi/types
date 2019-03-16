package client

const (
	ConnectionPoolSettings_TCPSettings_TcpKeepaliveType          = "connectionPoolSettings_TCPSettings_TcpKeepalive"
	ConnectionPoolSettings_TCPSettings_TcpKeepaliveFieldInterval = "interval"
	ConnectionPoolSettings_TCPSettings_TcpKeepaliveFieldProbes   = "probes"
	ConnectionPoolSettings_TCPSettings_TcpKeepaliveFieldTime     = "time"
)

type ConnectionPoolSettings_TCPSettings_TcpKeepalive struct {
	Interval *Duration `json:"interval,omitempty" yaml:"interval,omitempty"`
	Probes   int64     `json:"probes,omitempty" yaml:"probes,omitempty"`
	Time     *Duration `json:"time,omitempty" yaml:"time,omitempty"`
}
