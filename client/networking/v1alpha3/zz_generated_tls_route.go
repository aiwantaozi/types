package client

const (
	TLSRouteType       = "tlsRoute"
	TLSRouteFieldMatch = "match"
	TLSRouteFieldRoute = "route"
)

type TLSRoute struct {
	Match []TLSMatchAttributes `json:"match,omitempty" yaml:"match,omitempty"`
	Route []RouteDestination   `json:"route,omitempty" yaml:"route,omitempty"`
}
