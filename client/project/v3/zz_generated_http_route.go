package client

const (
	HTTPRouteType                       = "httpRoute"
	HTTPRouteFieldAppendHeaders         = "append_headers"
	HTTPRouteFieldAppendRequestHeaders  = "append_request_headers"
	HTTPRouteFieldAppendResponseHeaders = "append_response_headers"
	HTTPRouteFieldCorsPolicy            = "cors_policy"
	HTTPRouteFieldFault                 = "fault"
	HTTPRouteFieldHeaders               = "headers"
	HTTPRouteFieldMatch                 = "match"
	HTTPRouteFieldMirror                = "mirror"
	HTTPRouteFieldRedirect              = "redirect"
	HTTPRouteFieldRemoveRequestHeaders  = "remove_request_headers"
	HTTPRouteFieldRemoveResponseHeaders = "remove_response_headers"
	HTTPRouteFieldRetries               = "retries"
	HTTPRouteFieldRewrite               = "rewrite"
	HTTPRouteFieldRoute                 = "route"
	HTTPRouteFieldTimeout               = "timeout"
	HTTPRouteFieldWebsocketUpgrade      = "websocket_upgrade"
)

type HTTPRoute struct {
	AppendHeaders         map[string]string      `json:"append_headers,omitempty" yaml:"append_headers,omitempty"`
	AppendRequestHeaders  map[string]string      `json:"append_request_headers,omitempty" yaml:"append_request_headers,omitempty"`
	AppendResponseHeaders map[string]string      `json:"append_response_headers,omitempty" yaml:"append_response_headers,omitempty"`
	CorsPolicy            *CorsPolicy            `json:"cors_policy,omitempty" yaml:"cors_policy,omitempty"`
	Fault                 *HTTPFaultInjection    `json:"fault,omitempty" yaml:"fault,omitempty"`
	Headers               *Headers               `json:"headers,omitempty" yaml:"headers,omitempty"`
	Match                 []HTTPMatchRequest     `json:"match,omitempty" yaml:"match,omitempty"`
	Mirror                *Destination           `json:"mirror,omitempty" yaml:"mirror,omitempty"`
	Redirect              *HTTPRedirect          `json:"redirect,omitempty" yaml:"redirect,omitempty"`
	RemoveRequestHeaders  []string               `json:"remove_request_headers,omitempty" yaml:"remove_request_headers,omitempty"`
	RemoveResponseHeaders []string               `json:"remove_response_headers,omitempty" yaml:"remove_response_headers,omitempty"`
	Retries               *HTTPRetry             `json:"retries,omitempty" yaml:"retries,omitempty"`
	Rewrite               *HTTPRewrite           `json:"rewrite,omitempty" yaml:"rewrite,omitempty"`
	Route                 []HTTPRouteDestination `json:"route,omitempty" yaml:"route,omitempty"`
	Timeout               *Duration              `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	WebsocketUpgrade      bool                   `json:"websocket_upgrade,omitempty" yaml:"websocket_upgrade,omitempty"`
}
