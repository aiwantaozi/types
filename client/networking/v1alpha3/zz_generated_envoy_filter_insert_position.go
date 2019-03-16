package client

const (
	EnvoyFilter_InsertPositionType            = "envoyFilter_InsertPosition"
	EnvoyFilter_InsertPositionFieldIndex      = "index"
	EnvoyFilter_InsertPositionFieldRelativeTo = "relative_to"
)

type EnvoyFilter_InsertPosition struct {
	Index      int64  `json:"index,omitempty" yaml:"index,omitempty"`
	RelativeTo string `json:"relative_to,omitempty" yaml:"relative_to,omitempty"`
}
