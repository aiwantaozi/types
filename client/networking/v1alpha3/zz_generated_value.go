package client

const (
	ValueType      = "value"
	ValueFieldKind = "kind"
)

type Value struct {
	Kind interface{} `json:"kind,omitempty" yaml:"kind,omitempty"`
}
