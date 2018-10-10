package client

const (
	TargetPodType                        = "targetPod"
	TargetPodFieldCondition              = "condition"
	TargetPodFieldPodID                  = "podId"
	TargetPodFieldRestartIntervalSeconds = "restartIntervalSeconds"
	TargetPodFieldRestartTimes           = "restartTimes"
	TargetPodFieldRuleID                 = "ruleID"
	TargetPodFieldSeverity               = "severity"
)

type TargetPod struct {
	Condition              string `json:"condition,omitempty" yaml:"condition,omitempty"`
	PodID                  string `json:"podId,omitempty" yaml:"podId,omitempty"`
	RestartIntervalSeconds int64  `json:"restartIntervalSeconds,omitempty" yaml:"restartIntervalSeconds,omitempty"`
	RestartTimes           int64  `json:"restartTimes,omitempty" yaml:"restartTimes,omitempty"`
	RuleID                 string `json:"ruleID,omitempty" yaml:"ruleID,omitempty"`
	Severity               string `json:"severity,omitempty" yaml:"severity,omitempty"`
}
