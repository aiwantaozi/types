package client

const (
	TargetEventType                       = "targetEvent"
	TargetEventFieldEventType             = "eventType"
	TargetEventFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	TargetEventFieldGroupWaitSeconds      = "groupWaitSeconds"
	TargetEventFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	TargetEventFieldResourceKind          = "resourceKind"
	TargetEventFieldRuleID                = "ruleID"
	TargetEventFieldSeverity              = "severity"
)

type TargetEvent struct {
	EventType             string `json:"eventType,omitempty" yaml:"eventType,omitempty"`
	GroupIntervalSeconds  int64  `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64  `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	RepeatIntervalSeconds int64  `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
	ResourceKind          string `json:"resourceKind,omitempty" yaml:"resourceKind,omitempty"`
	RuleID                string `json:"ruleID,omitempty" yaml:"ruleID,omitempty"`
	Severity              string `json:"severity,omitempty" yaml:"severity,omitempty"`
}
