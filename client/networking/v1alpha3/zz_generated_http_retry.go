package client

const (
	HTTPRetryType               = "httpRetry"
	HTTPRetryFieldAttempts      = "attempts"
	HTTPRetryFieldPerTryTimeout = "per_try_timeout"
	HTTPRetryFieldRetryOn       = "retry_on"
)

type HTTPRetry struct {
	Attempts      int64     `json:"attempts,omitempty" yaml:"attempts,omitempty"`
	PerTryTimeout *Duration `json:"per_try_timeout,omitempty" yaml:"per_try_timeout,omitempty"`
	RetryOn       string    `json:"retry_on,omitempty" yaml:"retry_on,omitempty"`
}
