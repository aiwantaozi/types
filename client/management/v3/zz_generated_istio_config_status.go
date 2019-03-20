package client

const (
	IstioConfigStatusType            = "istioConfigStatus"
	IstioConfigStatusFieldConditions = "conditions"
)

type IstioConfigStatus struct {
	Conditions []Condition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}
