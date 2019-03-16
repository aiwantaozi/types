package client

const (
	DestinationRuleType               = "destinationRule"
	DestinationRuleFieldExportTo      = "export_to"
	DestinationRuleFieldHost          = "host"
	DestinationRuleFieldSubsets       = "subsets"
	DestinationRuleFieldTrafficPolicy = "traffic_policy"
)

type DestinationRule struct {
	ExportTo      []string       `json:"export_to,omitempty" yaml:"export_to,omitempty"`
	Host          string         `json:"host,omitempty" yaml:"host,omitempty"`
	Subsets       []Subset       `json:"subsets,omitempty" yaml:"subsets,omitempty"`
	TrafficPolicy *TrafficPolicy `json:"traffic_policy,omitempty" yaml:"traffic_policy,omitempty"`
}
