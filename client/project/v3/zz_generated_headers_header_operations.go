package client

const (
	Headers_HeaderOperationsType        = "headers_HeaderOperations"
	Headers_HeaderOperationsFieldAdd    = "add"
	Headers_HeaderOperationsFieldRemove = "remove"
	Headers_HeaderOperationsFieldSet    = "set"
)

type Headers_HeaderOperations struct {
	Add    map[string]string `json:"add,omitempty" yaml:"add,omitempty"`
	Remove []string          `json:"remove,omitempty" yaml:"remove,omitempty"`
	Set    map[string]string `json:"set,omitempty" yaml:"set,omitempty"`
}
