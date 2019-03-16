package client

const (
	HeadersType          = "headers"
	HeadersFieldRequest  = "request"
	HeadersFieldResponse = "response"
)

type Headers struct {
	Request  *Headers_HeaderOperations `json:"request,omitempty" yaml:"request,omitempty"`
	Response *Headers_HeaderOperations `json:"response,omitempty" yaml:"response,omitempty"`
}
