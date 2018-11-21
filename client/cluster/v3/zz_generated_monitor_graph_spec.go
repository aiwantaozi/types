package client

const (
	MonitorGraphSpecType                        = "monitorGraphSpec"
	MonitorGraphSpecFieldDescription            = "description"
	MonitorGraphSpecFieldDetailsMetricsSelector = "detailsMetricsSelector"
	MonitorGraphSpecFieldEnable                 = "enable"
	MonitorGraphSpecFieldIsSingleStats          = "isSingleStats"
	MonitorGraphSpecFieldMetricsSelector        = "metricsSelector"
	MonitorGraphSpecFieldPriority               = "priority"
	MonitorGraphSpecFieldThresholds             = "thresholds"
	MonitorGraphSpecFieldTitle                  = "title"
	MonitorGraphSpecFieldXAxis                  = "xAxis"
	MonitorGraphSpecFieldYAxis                  = "yAxis"
)

type MonitorGraphSpec struct {
	Description            string            `json:"description,omitempty" yaml:"description,omitempty"`
	DetailsMetricsSelector map[string]string `json:"detailsMetricsSelector,omitempty" yaml:"detailsMetricsSelector,omitempty"`
	Enable                 bool              `json:"enable,omitempty" yaml:"enable,omitempty"`
	IsSingleStats          bool              `json:"isSingleStats,omitempty" yaml:"isSingleStats,omitempty"`
	MetricsSelector        map[string]string `json:"metricsSelector,omitempty" yaml:"metricsSelector,omitempty"`
	Priority               int64             `json:"priority,omitempty" yaml:"priority,omitempty"`
	Thresholds             *float64          `json:"thresholds,omitempty" yaml:"thresholds,omitempty"`
	Title                  string            `json:"title,omitempty" yaml:"title,omitempty"`
	XAxis                  *XAxis            `json:"xAxis,omitempty" yaml:"xAxis,omitempty"`
	YAxis                  *YAxis            `json:"yAxis,omitempty" yaml:"yAxis,omitempty"`
}
