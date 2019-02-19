package client

const (
	HTTPConfigType             = "httpConfig"
	HTTPConfigFieldBasicAuth   = "basicAuth"
	HTTPConfigFieldBearerToken = "bearerToken"
	HTTPConfigFieldProxyURL    = "proxyURL"
	HTTPConfigFieldTLSConfig   = "tlsConfig"
)

type HTTPConfig struct {
	BasicAuth   *BasicAuth `json:"basicAuth,omitempty" yaml:"basicAuth,omitempty"`
	BearerToken string     `json:"bearerToken,omitempty" yaml:"bearerToken,omitempty"`
	ProxyURL    string     `json:"proxyURL,omitempty" yaml:"proxyURL,omitempty"`
	TLSConfig   *TLSConfig `json:"tlsConfig,omitempty" yaml:"tlsConfig,omitempty"`
}
