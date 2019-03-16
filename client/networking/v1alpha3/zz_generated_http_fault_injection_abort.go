package client

const (
	HTTPFaultInjection_AbortType            = "httpFaultInjection_Abort"
	HTTPFaultInjection_AbortFieldErrorType  = "errorType"
	HTTPFaultInjection_AbortFieldPercent    = "percent"
	HTTPFaultInjection_AbortFieldPercentage = "percentage"
)

type HTTPFaultInjection_Abort struct {
	ErrorType  interface{} `json:"errorType,omitempty" yaml:"errorType,omitempty"`
	Percent    int64       `json:"percent,omitempty" yaml:"percent,omitempty"`
	Percentage *Percent    `json:"percentage,omitempty" yaml:"percentage,omitempty"`
}
