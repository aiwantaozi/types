package client

const (
	StringMatchType           = "stringMatch"
	StringMatchFieldMatchType = "matchType"
)

type StringMatch struct {
	MatchType interface{} `json:"matchType,omitempty" yaml:"matchType,omitempty"`
}
