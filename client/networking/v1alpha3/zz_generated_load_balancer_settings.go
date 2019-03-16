package client

const (
	LoadBalancerSettingsType          = "loadBalancerSettings"
	LoadBalancerSettingsFieldLbPolicy = "lbPolicy"
)

type LoadBalancerSettings struct {
	LbPolicy interface{} `json:"lbPolicy,omitempty" yaml:"lbPolicy,omitempty"`
}
