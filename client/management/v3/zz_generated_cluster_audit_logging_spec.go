package client

const (
	ClusterAuditLoggingSpecType                 = "clusterAuditLoggingSpec"
	ClusterAuditLoggingSpecFieldClusterId       = "clusterId"
	ClusterAuditLoggingSpecFieldDisplayName     = "displayName"
	ClusterAuditLoggingSpecFieldLoggingTargetId = "loggingTargetId"
)

type ClusterAuditLoggingSpec struct {
	ClusterId       string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	DisplayName     string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	LoggingTargetId string `json:"loggingTargetId,omitempty" yaml:"loggingTargetId,omitempty"`
}
