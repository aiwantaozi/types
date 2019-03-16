package client

const (
	L4MatchAttributesType                    = "l4MatchAttributes"
	L4MatchAttributesFieldDestinationSubnets = "destination_subnets"
	L4MatchAttributesFieldGateways           = "gateways"
	L4MatchAttributesFieldPort               = "port"
	L4MatchAttributesFieldSourceLabels       = "source_labels"
	L4MatchAttributesFieldSourceSubnet       = "source_subnet"
)

type L4MatchAttributes struct {
	DestinationSubnets []string          `json:"destination_subnets,omitempty" yaml:"destination_subnets,omitempty"`
	Gateways           []string          `json:"gateways,omitempty" yaml:"gateways,omitempty"`
	Port               int64             `json:"port,omitempty" yaml:"port,omitempty"`
	SourceLabels       map[string]string `json:"source_labels,omitempty" yaml:"source_labels,omitempty"`
	SourceSubnet       string            `json:"source_subnet,omitempty" yaml:"source_subnet,omitempty"`
}
