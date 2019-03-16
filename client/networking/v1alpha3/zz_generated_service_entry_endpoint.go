package client

const (
	ServiceEntry_EndpointType          = "serviceEntry_Endpoint"
	ServiceEntry_EndpointFieldAddress  = "address"
	ServiceEntry_EndpointFieldLabels   = "labels"
	ServiceEntry_EndpointFieldLocality = "locality"
	ServiceEntry_EndpointFieldNetwork  = "network"
	ServiceEntry_EndpointFieldPorts    = "ports"
	ServiceEntry_EndpointFieldWeight   = "weight"
)

type ServiceEntry_Endpoint struct {
	Address  string            `json:"address,omitempty" yaml:"address,omitempty"`
	Labels   map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Locality string            `json:"locality,omitempty" yaml:"locality,omitempty"`
	Network  string            `json:"network,omitempty" yaml:"network,omitempty"`
	Ports    map[string]int64  `json:"ports,omitempty" yaml:"ports,omitempty"`
	Weight   int64             `json:"weight,omitempty" yaml:"weight,omitempty"`
}
