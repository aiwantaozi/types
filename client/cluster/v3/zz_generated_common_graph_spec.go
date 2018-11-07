package client

const (
	CommonGraphSpecType                      = "commonGraphSpec"
	CommonGraphSpecFieldMonitorGraphSelector = "monitorGraphSelector"
)

type CommonGraphSpec struct {
	MonitorGraphSelector map[string]string `json:"monitorGraphSelector,omitempty" yaml:"monitorGraphSelector,omitempty"`
}
