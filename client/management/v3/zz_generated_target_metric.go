package client

const (
	TargetMetricType                       = "targetMetric"
	TargetMetricFieldComparison            = "comparison"
	TargetMetricFieldDescription           = "description"
	TargetMetricFieldDuration              = "duration"
	TargetMetricFieldExpression            = "expression"
	TargetMetricFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	TargetMetricFieldGroupWaitSeconds      = "groupWaitSeconds"
	TargetMetricFieldLegendFormat          = "legendFormat"
	TargetMetricFieldName                  = "name"
	TargetMetricFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	TargetMetricFieldRuleID                = "ruleID"
	TargetMetricFieldSeverity              = "severity"
	TargetMetricFieldStep                  = "step"
	TargetMetricFieldThresholdValue        = "thresholdValue"
)

type TargetMetric struct {
	Comparison            string   `json:"comparison,omitempty" yaml:"comparison,omitempty"`
	Description           string   `json:"description,omitempty" yaml:"description,omitempty"`
	Duration              string   `json:"duration,omitempty" yaml:"duration,omitempty"`
	Expression            string   `json:"expression,omitempty" yaml:"expression,omitempty"`
	GroupIntervalSeconds  int64    `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64    `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	LegendFormat          string   `json:"legendFormat,omitempty" yaml:"legendFormat,omitempty"`
	Name                  string   `json:"name,omitempty" yaml:"name,omitempty"`
	RepeatIntervalSeconds int64    `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
	RuleID                string   `json:"ruleID,omitempty" yaml:"ruleID,omitempty"`
	Severity              string   `json:"severity,omitempty" yaml:"severity,omitempty"`
	Step                  int64    `json:"step,omitempty" yaml:"step,omitempty"`
	ThresholdValue        *float64 `json:"thresholdValue,omitempty" yaml:"thresholdValue,omitempty"`
}
