package client

const (
	IstioConfigSpecType           = "istioConfigSpec"
	IstioConfigSpecFieldClusterID = "clusterId"
	IstioConfigSpecFieldEnable    = "enable"
)

type IstioConfigSpec struct {
	ClusterID string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	Enable    bool   `json:"enable,omitempty" yaml:"enable,omitempty"`
}
