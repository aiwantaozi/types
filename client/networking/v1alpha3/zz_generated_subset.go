package client

const (
	SubsetType               = "subset"
	SubsetFieldLabels        = "labels"
	SubsetFieldName          = "name"
	SubsetFieldTrafficPolicy = "traffic_policy"
)

type Subset struct {
	Labels        map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name          string            `json:"name,omitempty" yaml:"name,omitempty"`
	TrafficPolicy *TrafficPolicy    `json:"traffic_policy,omitempty" yaml:"traffic_policy,omitempty"`
}
