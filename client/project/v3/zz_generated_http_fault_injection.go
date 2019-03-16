package client

const (
	HTTPFaultInjectionType       = "httpFaultInjection"
	HTTPFaultInjectionFieldAbort = "abort"
	HTTPFaultInjectionFieldDelay = "delay"
)

type HTTPFaultInjection struct {
	Abort *HTTPFaultInjection_Abort `json:"abort,omitempty" yaml:"abort,omitempty"`
	Delay *HTTPFaultInjection_Delay `json:"delay,omitempty" yaml:"delay,omitempty"`
}
