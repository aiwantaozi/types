package client

const (
	EnvoyFilter_ListenerMatchType                  = "envoyFilter_ListenerMatch"
	EnvoyFilter_ListenerMatchFieldAddress          = "address"
	EnvoyFilter_ListenerMatchFieldListenerProtocol = "listener_protocol"
	EnvoyFilter_ListenerMatchFieldListenerType     = "listener_type"
	EnvoyFilter_ListenerMatchFieldPortNamePrefix   = "port_name_prefix"
	EnvoyFilter_ListenerMatchFieldPortNumber       = "port_number"
)

type EnvoyFilter_ListenerMatch struct {
	Address          []string `json:"address,omitempty" yaml:"address,omitempty"`
	ListenerProtocol int64    `json:"listener_protocol,omitempty" yaml:"listener_protocol,omitempty"`
	ListenerType     int64    `json:"listener_type,omitempty" yaml:"listener_type,omitempty"`
	PortNamePrefix   string   `json:"port_name_prefix,omitempty" yaml:"port_name_prefix,omitempty"`
	PortNumber       int64    `json:"port_number,omitempty" yaml:"port_number,omitempty"`
}
