package client

const (
	DurationType         = "duration"
	DurationFieldNanos   = "nanos"
	DurationFieldSeconds = "seconds"
)

type Duration struct {
	Nanos   int64 `json:"nanos,omitempty" yaml:"nanos,omitempty"`
	Seconds int64 `json:"seconds,omitempty" yaml:"seconds,omitempty"`
}
