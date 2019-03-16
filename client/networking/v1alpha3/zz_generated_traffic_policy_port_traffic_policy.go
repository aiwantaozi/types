package client

const (
	TrafficPolicy_PortTrafficPolicyType                  = "trafficPolicy_PortTrafficPolicy"
	TrafficPolicy_PortTrafficPolicyFieldConnectionPool   = "connection_pool"
	TrafficPolicy_PortTrafficPolicyFieldLoadBalancer     = "load_balancer"
	TrafficPolicy_PortTrafficPolicyFieldOutlierDetection = "outlier_detection"
	TrafficPolicy_PortTrafficPolicyFieldPort             = "port"
	TrafficPolicy_PortTrafficPolicyFieldTls              = "tls"
)

type TrafficPolicy_PortTrafficPolicy struct {
	ConnectionPool   *ConnectionPoolSettings `json:"connection_pool,omitempty" yaml:"connection_pool,omitempty"`
	LoadBalancer     *LoadBalancerSettings   `json:"load_balancer,omitempty" yaml:"load_balancer,omitempty"`
	OutlierDetection *OutlierDetection       `json:"outlier_detection,omitempty" yaml:"outlier_detection,omitempty"`
	Port             *PortSelector           `json:"port,omitempty" yaml:"port,omitempty"`
	Tls              *TLSSettings            `json:"tls,omitempty" yaml:"tls,omitempty"`
}
