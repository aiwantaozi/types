package client

const (
	TLSConfigType            = "tlsConfig"
	TLSConfigFieldCA         = "ca"
	TLSConfigFieldCert       = "cert"
	TLSConfigFieldKey        = "key"
	TLSConfigFieldSSLVerify  = "sslVerify"
	TLSConfigFieldServerName = "serverName"
)

type TLSConfig struct {
	CA         string `json:"ca,omitempty" yaml:"ca,omitempty"`
	Cert       string `json:"cert,omitempty" yaml:"cert,omitempty"`
	Key        string `json:"key,omitempty" yaml:"key,omitempty"`
	SSLVerify  bool   `json:"sslVerify,omitempty" yaml:"sslVerify,omitempty"`
	ServerName string `json:"serverName,omitempty" yaml:"serverName,omitempty"`
}
