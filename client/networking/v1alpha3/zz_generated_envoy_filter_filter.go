package client

const (
	EnvoyFilter_FilterType                = "envoyFilter_Filter"
	EnvoyFilter_FilterFieldFilterConfig   = "filter_config"
	EnvoyFilter_FilterFieldFilterName     = "filter_name"
	EnvoyFilter_FilterFieldFilterType     = "filter_type"
	EnvoyFilter_FilterFieldInsertPosition = "insert_position"
	EnvoyFilter_FilterFieldListenerMatch  = "listener_match"
)

type EnvoyFilter_Filter struct {
	FilterConfig   *Struct                     `json:"filter_config,omitempty" yaml:"filter_config,omitempty"`
	FilterName     string                      `json:"filter_name,omitempty" yaml:"filter_name,omitempty"`
	FilterType     int64                       `json:"filter_type,omitempty" yaml:"filter_type,omitempty"`
	InsertPosition *EnvoyFilter_InsertPosition `json:"insert_position,omitempty" yaml:"insert_position,omitempty"`
	ListenerMatch  *EnvoyFilter_ListenerMatch  `json:"listener_match,omitempty" yaml:"listener_match,omitempty"`
}
