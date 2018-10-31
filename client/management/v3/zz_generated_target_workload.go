package client

const (
	TargetWorkloadType                       = "targetWorkload"
	TargetWorkloadFieldAvailablePercentage   = "availablePercentage"
	TargetWorkloadFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	TargetWorkloadFieldGroupWaitSeconds      = "groupWaitSeconds"
	TargetWorkloadFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	TargetWorkloadFieldRuleID                = "ruleID"
	TargetWorkloadFieldSelector              = "selector"
	TargetWorkloadFieldSeverity              = "severity"
	TargetWorkloadFieldWorkloadID            = "workloadId"
)

type TargetWorkload struct {
	AvailablePercentage   int64             `json:"availablePercentage,omitempty" yaml:"availablePercentage,omitempty"`
	GroupIntervalSeconds  int64             `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64             `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	RepeatIntervalSeconds int64             `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
	RuleID                string            `json:"ruleID,omitempty" yaml:"ruleID,omitempty"`
	Selector              map[string]string `json:"selector,omitempty" yaml:"selector,omitempty"`
	Severity              string            `json:"severity,omitempty" yaml:"severity,omitempty"`
	WorkloadID            string            `json:"workloadId,omitempty" yaml:"workloadId,omitempty"`
}
