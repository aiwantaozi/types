package client

const (
	PercentType       = "percent"
	PercentFieldValue = "value"
)

type Percent struct {
	Value float64 `json:"value,omitempty" yaml:"value,omitempty"`
}
