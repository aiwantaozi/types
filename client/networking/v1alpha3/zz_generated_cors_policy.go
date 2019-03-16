package client

const (
	CorsPolicyType                  = "corsPolicy"
	CorsPolicyFieldAllowCredentials = "allow_credentials"
	CorsPolicyFieldAllowHeaders     = "allow_headers"
	CorsPolicyFieldAllowMethods     = "allow_methods"
	CorsPolicyFieldAllowOrigin      = "allow_origin"
	CorsPolicyFieldExposeHeaders    = "expose_headers"
	CorsPolicyFieldMaxAge           = "max_age"
)

type CorsPolicy struct {
	AllowCredentials *BoolValue `json:"allow_credentials,omitempty" yaml:"allow_credentials,omitempty"`
	AllowHeaders     []string   `json:"allow_headers,omitempty" yaml:"allow_headers,omitempty"`
	AllowMethods     []string   `json:"allow_methods,omitempty" yaml:"allow_methods,omitempty"`
	AllowOrigin      []string   `json:"allow_origin,omitempty" yaml:"allow_origin,omitempty"`
	ExposeHeaders    []string   `json:"expose_headers,omitempty" yaml:"expose_headers,omitempty"`
	MaxAge           *Duration  `json:"max_age,omitempty" yaml:"max_age,omitempty"`
}
