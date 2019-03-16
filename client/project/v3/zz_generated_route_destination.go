package client

const (
	RouteDestinationType             = "routeDestination"
	RouteDestinationFieldDestination = "destination"
	RouteDestinationFieldWeight      = "weight"
)

type RouteDestination struct {
	Destination *Destination `json:"destination,omitempty" yaml:"destination,omitempty"`
	Weight      int64        `json:"weight,omitempty" yaml:"weight,omitempty"`
}
