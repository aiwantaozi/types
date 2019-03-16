package client

const (
	OutlierDetectionType                    = "outlierDetection"
	OutlierDetectionFieldBaseEjectionTime   = "base_ejection_time"
	OutlierDetectionFieldConsecutiveErrors  = "consecutive_errors"
	OutlierDetectionFieldInterval           = "interval"
	OutlierDetectionFieldMaxEjectionPercent = "max_ejection_percent"
	OutlierDetectionFieldMinHealthPercent   = "min_health_percent"
)

type OutlierDetection struct {
	BaseEjectionTime   *Duration `json:"base_ejection_time,omitempty" yaml:"base_ejection_time,omitempty"`
	ConsecutiveErrors  int64     `json:"consecutive_errors,omitempty" yaml:"consecutive_errors,omitempty"`
	Interval           *Duration `json:"interval,omitempty" yaml:"interval,omitempty"`
	MaxEjectionPercent int64     `json:"max_ejection_percent,omitempty" yaml:"max_ejection_percent,omitempty"`
	MinHealthPercent   int64     `json:"min_health_percent,omitempty" yaml:"min_health_percent,omitempty"`
}
