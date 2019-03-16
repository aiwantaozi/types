package client

const (
	VirtualServiceType          = "virtualService"
	VirtualServiceFieldExportTo = "export_to"
	VirtualServiceFieldGateways = "gateways"
	VirtualServiceFieldHosts    = "hosts"
	VirtualServiceFieldHttp     = "http"
	VirtualServiceFieldTcp      = "tcp"
	VirtualServiceFieldTls      = "tls"
)

type VirtualService struct {
	ExportTo []string    `json:"export_to,omitempty" yaml:"export_to,omitempty"`
	Gateways []string    `json:"gateways,omitempty" yaml:"gateways,omitempty"`
	Hosts    []string    `json:"hosts,omitempty" yaml:"hosts,omitempty"`
	Http     []HTTPRoute `json:"http,omitempty" yaml:"http,omitempty"`
	Tcp      []TCPRoute  `json:"tcp,omitempty" yaml:"tcp,omitempty"`
	Tls      []TLSRoute  `json:"tls,omitempty" yaml:"tls,omitempty"`
}
