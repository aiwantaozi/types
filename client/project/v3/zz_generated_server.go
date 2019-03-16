package client

const (
	ServerType                 = "server"
	ServerFieldBind            = "bind"
	ServerFieldDefaultEndpoint = "default_endpoint"
	ServerFieldHosts           = "hosts"
	ServerFieldPort            = "port"
	ServerFieldTls             = "tls"
)

type Server struct {
	Bind            string             `json:"bind,omitempty" yaml:"bind,omitempty"`
	DefaultEndpoint string             `json:"default_endpoint,omitempty" yaml:"default_endpoint,omitempty"`
	Hosts           []string           `json:"hosts,omitempty" yaml:"hosts,omitempty"`
	Port            *Port              `json:"port,omitempty" yaml:"port,omitempty"`
	Tls             *Server_TLSOptions `json:"tls,omitempty" yaml:"tls,omitempty"`
}
