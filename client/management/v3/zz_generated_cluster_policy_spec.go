package client

const (
	ClusterPolicySpecType                       = "clusterPolicySpec"
	ClusterPolicySpecFieldClusterID             = "clusterId"
	ClusterPolicySpecFieldDescription           = "description"
	ClusterPolicySpecFieldDisplayName           = "displayName"
	ClusterPolicySpecFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	ClusterPolicySpecFieldGroupWaitSeconds      = "groupWaitSeconds"
	ClusterPolicySpecFieldMetrics               = "metrics"
	ClusterPolicySpecFieldRecipients            = "recipients"
	ClusterPolicySpecFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	ClusterPolicySpecFieldTargetEvents          = "targetEvents"
	ClusterPolicySpecFieldTargetNodes           = "targetNodes"
	ClusterPolicySpecFieldTargetSystemServices  = "targetSystemServices"
)

type ClusterPolicySpec struct {
	ClusterID             string                `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	Description           string                `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName           string                `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	GroupIntervalSeconds  int64                 `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64                 `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	Metrics               []Metric              `json:"metrics,omitempty" yaml:"metrics,omitempty"`
	Recipients            []Recipient           `json:"recipients,omitempty" yaml:"recipients,omitempty"`
	RepeatIntervalSeconds int64                 `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
	TargetEvents          []TargetEvent         `json:"targetEvents,omitempty" yaml:"targetEvents,omitempty"`
	TargetNodes           []TargetNode          `json:"targetNodes,omitempty" yaml:"targetNodes,omitempty"`
	TargetSystemServices  []TargetSystemService `json:"targetSystemServices,omitempty" yaml:"targetSystemServices,omitempty"`
}
