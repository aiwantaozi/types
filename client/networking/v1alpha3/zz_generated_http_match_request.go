package client

const (
	HTTPMatchRequestType              = "httpMatchRequest"
	HTTPMatchRequestFieldAuthority    = "authority"
	HTTPMatchRequestFieldGateways     = "gateways"
	HTTPMatchRequestFieldHeaders      = "headers"
	HTTPMatchRequestFieldMethod       = "method"
	HTTPMatchRequestFieldPort         = "port"
	HTTPMatchRequestFieldScheme       = "scheme"
	HTTPMatchRequestFieldSourceLabels = "source_labels"
	HTTPMatchRequestFieldUri          = "uri"
)

type HTTPMatchRequest struct {
	Authority    *StringMatch           `json:"authority,omitempty" yaml:"authority,omitempty"`
	Gateways     []string               `json:"gateways,omitempty" yaml:"gateways,omitempty"`
	Headers      map[string]StringMatch `json:"headers,omitempty" yaml:"headers,omitempty"`
	Method       *StringMatch           `json:"method,omitempty" yaml:"method,omitempty"`
	Port         int64                  `json:"port,omitempty" yaml:"port,omitempty"`
	Scheme       *StringMatch           `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	SourceLabels map[string]string      `json:"source_labels,omitempty" yaml:"source_labels,omitempty"`
	Uri          *StringMatch           `json:"uri,omitempty" yaml:"uri,omitempty"`
}
