package client

const (
	LoggingTargetStatusType             = "loggingTargetStatus"
	LoggingTargetStatusFieldAppliedSpec = "appliedSpec"
	LoggingTargetStatusFieldClusterIDs  = "clusterIds"
	LoggingTargetStatusFieldConditions  = "conditions"
)

type LoggingTargetStatus struct {
	AppliedSpec *LoggingTargetSpec `json:"appliedSpec,omitempty" yaml:"appliedSpec,omitempty"`
	ClusterIDs  []string           `json:"clusterIds,omitempty" yaml:"clusterIds,omitempty"`
	Conditions  []LoggingCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}
