package client

const (
	WorkloadMetricType          = "workloadMetric"
	WorkloadMetricFieldPath     = "path"
	WorkloadMetricFieldPort     = "port"
	WorkloadMetricFieldProtocol = "protocol"
)

type WorkloadMetric struct {
	Path     string `json:"path,omitempty" yaml:"path,omitempty"`
	Port     int64  `json:"port,omitempty" yaml:"port,omitempty"`
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}
