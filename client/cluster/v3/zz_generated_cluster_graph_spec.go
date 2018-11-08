package client

const (
	ClusterGraphSpecType           = "clusterGraphSpec"
	ClusterGraphSpecFieldProjectID = "projectId"
)

type ClusterGraphSpec struct {
	ProjectID string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
}
