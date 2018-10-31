package client

const (
	TargetPodType                        = "targetPod"
	TargetPodFieldCondition              = "condition"
	TargetPodFieldGroupIntervalSeconds   = "groupIntervalSeconds"
	TargetPodFieldGroupWaitSeconds       = "groupWaitSeconds"
	TargetPodFieldPodID                  = "podId"
	TargetPodFieldRepeatIntervalSeconds  = "repeatIntervalSeconds"
	TargetPodFieldRestartIntervalSeconds = "restartIntervalSeconds"
	TargetPodFieldRestartTimes           = "restartTimes"
	TargetPodFieldRuleID                 = "ruleID"
	TargetPodFieldSeverity               = "severity"
)

type TargetPod struct {
	Condition              string `json:"condition,omitempty" yaml:"condition,omitempty"`
	GroupIntervalSeconds   int64  `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds       int64  `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	PodID                  string `json:"podId,omitempty" yaml:"podId,omitempty"`
	RepeatIntervalSeconds  int64  `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
	RestartIntervalSeconds int64  `json:"restartIntervalSeconds,omitempty" yaml:"restartIntervalSeconds,omitempty"`
	RestartTimes           int64  `json:"restartTimes,omitempty" yaml:"restartTimes,omitempty"`
	RuleID                 string `json:"ruleID,omitempty" yaml:"ruleID,omitempty"`
	Severity               string `json:"severity,omitempty" yaml:"severity,omitempty"`
}
