package v3

import (
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterAlertPolicy struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterPolicySpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status AlertStatus `json:"status"`
}

type ProjectAlertPolicy struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProjectPolicySpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status AlertStatus `json:"status"`
}

type ClusterPolicySpec struct {
	ClusterName          string                `json:"clusterName" norman:"type=reference[cluster]"`
	TargetNodes          []TargetNode          `json:"targetNodes,omitempty"`
	TargetEvents         []TargetEvent         `json:"targetEvents,omitempty"`
	TargetSystemServices []TargetSystemService `json:"targetSystemServices,omitempty"`
	Metrics              []Metric              `json:"metrics,omitempty"`
	Recipients           []Recipient           `json:"recipients,omitempty" norman:"required"`
	CommonPolicy
}

type ProjectPolicySpec struct {
	ProjectName     string           `json:"projectName" norman:"type=reference[project]"`
	TargetPods      []TargetPod      `json:"targetPods,omitempty"`
	TargetWorkloads []TargetWorkload `json:"targetWorkloads,omitempty"`
	Metrics         []Metric         `json:"metrics,omitempty"`
	Recipients      []Recipient      `json:"recipients,omitempty" norman:"required"`
	CommonPolicy
}

type CommonPolicy struct {
	DisplayName string `json:"displayName,omitempty" norman:"required"`
	Description string `json:"description,omitempty"`
	TimingField
}

type CommonRuleField struct {
	RuleID   string `json:"ruleID,omitempty"`
	Severity string `json:"severity,omitempty" norman:"required,options=info|critical|warning,default=critical"`
	TimingField
}

type Metric struct {
	CommonRuleField
	Name           string  `json:"name,omitempty"`
	Expression     string  `json:"expression,omitempty"`
	Comparison     string  `json:"comparison,omitempty" norman:"required,type=enum,options=equal|not-equal|greater-than|less-than|greater-or-equal|less-or-equal"`
	Duration       string  `json:"duration",omitempty`
	ThresholdValue float64 `json:"thresholdValue,omitempty" norman:"required,type=float"`
}

type TimingField struct {
	GroupWaitSeconds      int `json:"groupWaitSeconds,omitempty" norman:"required,default=30,min=0"`
	GroupIntervalSeconds  int `json:"groupIntervalSeconds,omitempty" norman:"required,default=180,min=0"`
	RepeatIntervalSeconds int `json:"repeatIntervalSeconds,omitempty"  norman:"required,default=3600,min=0"`
}
