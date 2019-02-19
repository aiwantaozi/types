package client

const (
	WechatConfigType                  = "wechatConfig"
	WechatConfigFieldAgent            = "agent"
	WechatConfigFieldCorp             = "corp"
	WechatConfigFieldDefaultRecipient = "defaultRecipient"
	WechatConfigFieldHTTPConfig       = "httpConfig"
	WechatConfigFieldRecipientType    = "recipientType"
	WechatConfigFieldSecret           = "secret"
)

type WechatConfig struct {
	Agent            string      `json:"agent,omitempty" yaml:"agent,omitempty"`
	Corp             string      `json:"corp,omitempty" yaml:"corp,omitempty"`
	DefaultRecipient string      `json:"defaultRecipient,omitempty" yaml:"defaultRecipient,omitempty"`
	HTTPConfig       *HTTPConfig `json:"httpConfig,omitempty" yaml:"httpConfig,omitempty"`
	RecipientType    string      `json:"recipientType,omitempty" yaml:"recipientType,omitempty"`
	Secret           string      `json:"secret,omitempty" yaml:"secret,omitempty"`
}
