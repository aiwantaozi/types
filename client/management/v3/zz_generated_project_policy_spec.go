package client

const (
	ProjectPolicySpecType                       = "projectPolicySpec"
	ProjectPolicySpecFieldDescription           = "description"
	ProjectPolicySpecFieldDisplayName           = "displayName"
	ProjectPolicySpecFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	ProjectPolicySpecFieldGroupWaitSeconds      = "groupWaitSeconds"
	ProjectPolicySpecFieldProjectID             = "projectId"
	ProjectPolicySpecFieldRecipients            = "recipients"
	ProjectPolicySpecFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	ProjectPolicySpecFieldTargetMetrics         = "targetMetrics"
	ProjectPolicySpecFieldTargetPods            = "targetPods"
	ProjectPolicySpecFieldTargetWorkloads       = "targetWorkloads"
)

type ProjectPolicySpec struct {
	Description           string           `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName           string           `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	GroupIntervalSeconds  int64            `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64            `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	ProjectID             string           `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	Recipients            []Recipient      `json:"recipients,omitempty" yaml:"recipients,omitempty"`
	RepeatIntervalSeconds int64            `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
	TargetMetrics         []TargetMetric   `json:"targetMetrics,omitempty" yaml:"targetMetrics,omitempty"`
	TargetPods            []TargetPod      `json:"targetPods,omitempty" yaml:"targetPods,omitempty"`
	TargetWorkloads       []TargetWorkload `json:"targetWorkloads,omitempty" yaml:"targetWorkloads,omitempty"`
}
