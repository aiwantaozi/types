package client

const (
	TargetNodeType                       = "targetNode"
	TargetNodeFieldCPUThreshold          = "cpuThreshold"
	TargetNodeFieldCondition             = "condition"
	TargetNodeFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	TargetNodeFieldGroupWaitSeconds      = "groupWaitSeconds"
	TargetNodeFieldMemThreshold          = "memThreshold"
	TargetNodeFieldNodeID                = "nodeId"
	TargetNodeFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	TargetNodeFieldRuleID                = "ruleID"
	TargetNodeFieldSelector              = "selector"
	TargetNodeFieldSeverity              = "severity"
)

type TargetNode struct {
	CPUThreshold          int64             `json:"cpuThreshold,omitempty" yaml:"cpuThreshold,omitempty"`
	Condition             string            `json:"condition,omitempty" yaml:"condition,omitempty"`
	GroupIntervalSeconds  int64             `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64             `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	MemThreshold          int64             `json:"memThreshold,omitempty" yaml:"memThreshold,omitempty"`
	NodeID                string            `json:"nodeId,omitempty" yaml:"nodeId,omitempty"`
	RepeatIntervalSeconds int64             `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
	RuleID                string            `json:"ruleID,omitempty" yaml:"ruleID,omitempty"`
	Selector              map[string]string `json:"selector,omitempty" yaml:"selector,omitempty"`
	Severity              string            `json:"severity,omitempty" yaml:"severity,omitempty"`
}
