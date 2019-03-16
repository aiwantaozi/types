package client

const (
	BoolValueType       = "boolValue"
	BoolValueFieldValue = "value"
)

type BoolValue struct {
	Value bool `json:"value,omitempty" yaml:"value,omitempty"`
}
