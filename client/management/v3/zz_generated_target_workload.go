package client

const (
	TargetWorkloadType                     = "targetWorkload"
	TargetWorkloadFieldAvailablePercentage = "availablePercentage"
	TargetWorkloadFieldRuleID              = "ruleID"
	TargetWorkloadFieldSelector            = "selector"
	TargetWorkloadFieldSeverity            = "severity"
	TargetWorkloadFieldWorkloadID          = "workloadId"
)

type TargetWorkload struct {
	AvailablePercentage int64             `json:"availablePercentage,omitempty" yaml:"availablePercentage,omitempty"`
	RuleID              string            `json:"ruleID,omitempty" yaml:"ruleID,omitempty"`
	Selector            map[string]string `json:"selector,omitempty" yaml:"selector,omitempty"`
	Severity            string            `json:"severity,omitempty" yaml:"severity,omitempty"`
	WorkloadID          string            `json:"workloadId,omitempty" yaml:"workloadId,omitempty"`
}
