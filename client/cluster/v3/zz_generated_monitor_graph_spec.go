package client

const (
	MonitorGraphSpecType                 = "monitorGraphSpec"
	MonitorGraphSpecFieldDescription     = "description"
	MonitorGraphSpecFieldEnable          = "enable"
	MonitorGraphSpecFieldMetricsSelector = "metricsSelector"
	MonitorGraphSpecFieldThresholds      = "thresholds"
	MonitorGraphSpecFieldTitle           = "title"
	MonitorGraphSpecFieldXAxis           = "xAxis"
	MonitorGraphSpecFieldYAxis           = "yAxis"
)

type MonitorGraphSpec struct {
	Description     string            `json:"description,omitempty" yaml:"description,omitempty"`
	Enable          bool              `json:"enable,omitempty" yaml:"enable,omitempty"`
	MetricsSelector map[string]string `json:"metricsSelector,omitempty" yaml:"metricsSelector,omitempty"`
	Thresholds      *float64          `json:"thresholds,omitempty" yaml:"thresholds,omitempty"`
	Title           string            `json:"title,omitempty" yaml:"title,omitempty"`
	XAxis           *XAxis            `json:"xAxis,omitempty" yaml:"xAxis,omitempty"`
	YAxis           *YAxis            `json:"yAxis,omitempty" yaml:"yAxis,omitempty"`
}
