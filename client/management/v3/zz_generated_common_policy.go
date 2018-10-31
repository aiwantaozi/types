package client

const (
	CommonPolicyType                       = "commonPolicy"
	CommonPolicyFieldDescription           = "description"
	CommonPolicyFieldDisplayName           = "displayName"
	CommonPolicyFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	CommonPolicyFieldGroupWaitSeconds      = "groupWaitSeconds"
	CommonPolicyFieldMetrics               = "metrics"
	CommonPolicyFieldRecipients            = "recipients"
	CommonPolicyFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
)

type CommonPolicy struct {
	Description           string      `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName           string      `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	GroupIntervalSeconds  int64       `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64       `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	Metrics               []Metric    `json:"metrics,omitempty" yaml:"metrics,omitempty"`
	Recipients            []Recipient `json:"recipients,omitempty" yaml:"recipients,omitempty"`
	RepeatIntervalSeconds int64       `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
}
