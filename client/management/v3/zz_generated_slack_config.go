package client

const (
	SlackConfigType                  = "slackConfig"
	SlackConfigFieldDefaultRecipient = "defaultRecipient"
	SlackConfigFieldHTTPConfig       = "httpConfig"
	SlackConfigFieldURL              = "url"
)

type SlackConfig struct {
	DefaultRecipient string      `json:"defaultRecipient,omitempty" yaml:"defaultRecipient,omitempty"`
	HTTPConfig       *HTTPConfig `json:"httpConfig,omitempty" yaml:"httpConfig,omitempty"`
	URL              string      `json:"url,omitempty" yaml:"url,omitempty"`
}
