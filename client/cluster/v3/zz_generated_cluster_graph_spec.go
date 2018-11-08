package client

const (
	ClusterGraphSpecType           = "clusterGraphSpec"
	ClusterGraphSpecFieldClusterID = "clusterId"
)

type ClusterGraphSpec struct {
	ClusterID string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
}
