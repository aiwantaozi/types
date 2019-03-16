package client

const (
	TrafficPolicyType                   = "trafficPolicy"
	TrafficPolicyFieldConnectionPool    = "connection_pool"
	TrafficPolicyFieldLoadBalancer      = "load_balancer"
	TrafficPolicyFieldOutlierDetection  = "outlier_detection"
	TrafficPolicyFieldPortLevelSettings = "port_level_settings"
	TrafficPolicyFieldTls               = "tls"
)

type TrafficPolicy struct {
	ConnectionPool    *ConnectionPoolSettings           `json:"connection_pool,omitempty" yaml:"connection_pool,omitempty"`
	LoadBalancer      *LoadBalancerSettings             `json:"load_balancer,omitempty" yaml:"load_balancer,omitempty"`
	OutlierDetection  *OutlierDetection                 `json:"outlier_detection,omitempty" yaml:"outlier_detection,omitempty"`
	PortLevelSettings []TrafficPolicy_PortTrafficPolicy `json:"port_level_settings,omitempty" yaml:"port_level_settings,omitempty"`
	Tls               *TLSSettings                      `json:"tls,omitempty" yaml:"tls,omitempty"`
}
