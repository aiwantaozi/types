package client

const (
	MetricType                       = "metric"
	MetricFieldComparison            = "comparison"
	MetricFieldDuration              = "duration"
	MetricFieldExpression            = "expression"
	MetricFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	MetricFieldGroupWaitSeconds      = "groupWaitSeconds"
	MetricFieldName                  = "name"
	MetricFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	MetricFieldRuleID                = "ruleID"
	MetricFieldSeverity              = "severity"
	MetricFieldThresholdValue        = "thresholdValue"
)

type Metric struct {
	Comparison            string   `json:"comparison,omitempty" yaml:"comparison,omitempty"`
	Duration              string   `json:"duration,omitempty" yaml:"duration,omitempty"`
	Expression            string   `json:"expression,omitempty" yaml:"expression,omitempty"`
	GroupIntervalSeconds  int64    `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64    `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	Name                  string   `json:"name,omitempty" yaml:"name,omitempty"`
	RepeatIntervalSeconds int64    `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
	RuleID                string   `json:"ruleID,omitempty" yaml:"ruleID,omitempty"`
	Severity              string   `json:"severity,omitempty" yaml:"severity,omitempty"`
	ThresholdValue        *float64 `json:"thresholdValue,omitempty" yaml:"thresholdValue,omitempty"`
}
