package client

const (
	HTTPRouteDestinationType                       = "httpRouteDestination"
	HTTPRouteDestinationFieldAppendRequestHeaders  = "append_request_headers"
	HTTPRouteDestinationFieldAppendResponseHeaders = "append_response_headers"
	HTTPRouteDestinationFieldDestination           = "destination"
	HTTPRouteDestinationFieldHeaders               = "headers"
	HTTPRouteDestinationFieldRemoveRequestHeaders  = "remove_request_headers"
	HTTPRouteDestinationFieldRemoveResponseHeaders = "remove_response_headers"
	HTTPRouteDestinationFieldWeight                = "weight"
)

type HTTPRouteDestination struct {
	AppendRequestHeaders  map[string]string `json:"append_request_headers,omitempty" yaml:"append_request_headers,omitempty"`
	AppendResponseHeaders map[string]string `json:"append_response_headers,omitempty" yaml:"append_response_headers,omitempty"`
	Destination           *Destination      `json:"destination,omitempty" yaml:"destination,omitempty"`
	Headers               *Headers          `json:"headers,omitempty" yaml:"headers,omitempty"`
	RemoveRequestHeaders  []string          `json:"remove_request_headers,omitempty" yaml:"remove_request_headers,omitempty"`
	RemoveResponseHeaders []string          `json:"remove_response_headers,omitempty" yaml:"remove_response_headers,omitempty"`
	Weight                int64             `json:"weight,omitempty" yaml:"weight,omitempty"`
}
