package client

const (
	MetricRuleType                = "metricRule"
	MetricRuleFieldComparison     = "comparison"
	MetricRuleFieldDescription    = "description"
	MetricRuleFieldDuration       = "duration"
	MetricRuleFieldExpression     = "expression"
	MetricRuleFieldLegendFormat   = "legendFormat"
	MetricRuleFieldThresholdValue = "thresholdValue"
)

type MetricRule struct {
	Comparison     string  `json:"comparison,omitempty" yaml:"comparison,omitempty"`
	Description    string  `json:"description,omitempty" yaml:"description,omitempty"`
	Duration       string  `json:"duration,omitempty" yaml:"duration,omitempty"`
	Expression     string  `json:"expression,omitempty" yaml:"expression,omitempty"`
	LegendFormat   string  `json:"legendFormat,omitempty" yaml:"legendFormat,omitempty"`
	ThresholdValue float64 `json:"thresholdValue,omitempty" yaml:"thresholdValue,omitempty"`
}
