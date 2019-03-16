package client

const (
	ConnectionPoolSettings_HTTPSettingsType                          = "connectionPoolSettings_HTTPSettings"
	ConnectionPoolSettings_HTTPSettingsFieldHttp1MaxPendingRequests  = "http1_max_pending_requests"
	ConnectionPoolSettings_HTTPSettingsFieldHttp2MaxRequests         = "http2_max_requests"
	ConnectionPoolSettings_HTTPSettingsFieldMaxRequestsPerConnection = "max_requests_per_connection"
	ConnectionPoolSettings_HTTPSettingsFieldMaxRetries               = "max_retries"
)

type ConnectionPoolSettings_HTTPSettings struct {
	Http1MaxPendingRequests  int64 `json:"http1_max_pending_requests,omitempty" yaml:"http1_max_pending_requests,omitempty"`
	Http2MaxRequests         int64 `json:"http2_max_requests,omitempty" yaml:"http2_max_requests,omitempty"`
	MaxRequestsPerConnection int64 `json:"max_requests_per_connection,omitempty" yaml:"max_requests_per_connection,omitempty"`
	MaxRetries               int64 `json:"max_retries,omitempty" yaml:"max_retries,omitempty"`
}
