package client

const (
	StructType        = "struct"
	StructFieldFields = "fields"
)

type Struct struct {
	Fields map[string]Value `json:"fields,omitempty" yaml:"fields,omitempty"`
}
