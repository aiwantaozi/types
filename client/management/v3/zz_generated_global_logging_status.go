package client

const (
	GlobalLoggingStatusType             = "globalLoggingStatus"
	GlobalLoggingStatusFieldAppliedSpec = "appliedSpec"
	GlobalLoggingStatusFieldClusterIDs  = "clusterIds"
	GlobalLoggingStatusFieldConditions  = "conditions"
)

type GlobalLoggingStatus struct {
	AppliedSpec *GlobalLoggingSpec `json:"appliedSpec,omitempty" yaml:"appliedSpec,omitempty"`
	ClusterIDs  []string           `json:"clusterIds,omitempty" yaml:"clusterIds,omitempty"`
	Conditions  []LoggingCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}
