package client

const (
	WebhookConfigType            = "webhookConfig"
	WebhookConfigFieldHTTPConfig = "httpConfig"
	WebhookConfigFieldURL        = "url"
)

type WebhookConfig struct {
	HTTPConfig *HTTPConfig `json:"httpConfig,omitempty" yaml:"httpConfig,omitempty"`
	URL        string      `json:"url,omitempty" yaml:"url,omitempty"`
}
