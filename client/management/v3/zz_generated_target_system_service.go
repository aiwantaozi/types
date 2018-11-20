package client

const (
	TargetSystemServiceType                       = "targetSystemService"
	TargetSystemServiceFieldCondition             = "condition"
	TargetSystemServiceFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	TargetSystemServiceFieldGroupWaitSeconds      = "groupWaitSeconds"
	TargetSystemServiceFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	TargetSystemServiceFieldRuleID                = "ruleID"
	TargetSystemServiceFieldSeverity              = "severity"
)

type TargetSystemService struct {
	Condition             string `json:"condition,omitempty" yaml:"condition,omitempty"`
	GroupIntervalSeconds  int64  `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64  `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	RepeatIntervalSeconds int64  `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
	RuleID                string `json:"ruleID,omitempty" yaml:"ruleID,omitempty"`
	Severity              string `json:"severity,omitempty" yaml:"severity,omitempty"`
}
