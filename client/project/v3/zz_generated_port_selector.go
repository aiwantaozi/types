package client

const (
	PortSelectorType      = "portSelector"
	PortSelectorFieldPort = "port"
)

type PortSelector struct {
	Port interface{} `json:"port,omitempty" yaml:"port,omitempty"`
}
