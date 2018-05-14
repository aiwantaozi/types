package client

const (
	GlobalLoggingStatusType            = "globalLoggingStatus"
	GlobalLoggingStatusFieldConditions = "conditions"
)

type GlobalLoggingStatus struct {
	Conditions []LoggingCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}
