package client

const (
	ProjectGraphSpecType           = "projectGraphSpec"
	ProjectGraphSpecFieldProjectID = "projectId"
)

type ProjectGraphSpec struct {
	ProjectID string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
}
