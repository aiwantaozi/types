package client

const (
	PagerdutyConfigType            = "pagerdutyConfig"
	PagerdutyConfigFieldHTTPConfig = "httpConfig"
	PagerdutyConfigFieldServiceKey = "serviceKey"
)

type PagerdutyConfig struct {
	HTTPConfig *HTTPConfig `json:"httpConfig,omitempty" yaml:"httpConfig,omitempty"`
	ServiceKey string      `json:"serviceKey,omitempty" yaml:"serviceKey,omitempty"`
}
