package client

const (
	HTTPFaultInjection_DelayType               = "httpFaultInjection_Delay"
	HTTPFaultInjection_DelayFieldHttpDelayType = "httpDelayType"
	HTTPFaultInjection_DelayFieldPercent       = "percent"
	HTTPFaultInjection_DelayFieldPercentage    = "percentage"
)

type HTTPFaultInjection_Delay struct {
	HttpDelayType interface{} `json:"httpDelayType,omitempty" yaml:"httpDelayType,omitempty"`
	Percent       int64       `json:"percent,omitempty" yaml:"percent,omitempty"`
	Percentage    *Percent    `json:"percentage,omitempty" yaml:"percentage,omitempty"`
}
