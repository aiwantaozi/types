package client

const (
	GatewayType          = "gateway"
	GatewayFieldSelector = "selector"
	GatewayFieldServers  = "servers"
)

type Gateway struct {
	Selector map[string]string `json:"selector,omitempty" yaml:"selector,omitempty"`
	Servers  []Server          `json:"servers,omitempty" yaml:"servers,omitempty"`
}
