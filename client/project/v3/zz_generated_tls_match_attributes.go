package client

const (
	TLSMatchAttributesType                    = "tlsMatchAttributes"
	TLSMatchAttributesFieldDestinationSubnets = "destination_subnets"
	TLSMatchAttributesFieldGateways           = "gateways"
	TLSMatchAttributesFieldPort               = "port"
	TLSMatchAttributesFieldSniHosts           = "sni_hosts"
	TLSMatchAttributesFieldSourceLabels       = "source_labels"
	TLSMatchAttributesFieldSourceSubnet       = "source_subnet"
)

type TLSMatchAttributes struct {
	DestinationSubnets []string          `json:"destination_subnets,omitempty" yaml:"destination_subnets,omitempty"`
	Gateways           []string          `json:"gateways,omitempty" yaml:"gateways,omitempty"`
	Port               int64             `json:"port,omitempty" yaml:"port,omitempty"`
	SniHosts           []string          `json:"sni_hosts,omitempty" yaml:"sni_hosts,omitempty"`
	SourceLabels       map[string]string `json:"source_labels,omitempty" yaml:"source_labels,omitempty"`
	SourceSubnet       string            `json:"source_subnet,omitempty" yaml:"source_subnet,omitempty"`
}
