package client

const (
	ClusterAuditLoggingSpecType                   = "clusterAuditLoggingSpec"
	ClusterAuditLoggingSpecFieldClusterId         = "clusterId"
	ClusterAuditLoggingSpecFieldDisplayName       = "displayName"
	ClusterAuditLoggingSpecFieldGlobalLoggingName = "globalLogging"
)

type ClusterAuditLoggingSpec struct {
	ClusterId         string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	DisplayName       string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	GlobalLoggingName string `json:"globalLogging,omitempty" yaml:"globalLogging,omitempty"`
}
