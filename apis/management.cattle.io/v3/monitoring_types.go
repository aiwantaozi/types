package v3

import (
	"github.com/rancher/norman/condition"
	"github.com/rancher/norman/types"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterGraph struct {
	types.Namespaced
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterGraphSpec `json:"spec"`
}

type ProjectGraph struct {
	types.Namespaced
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectGraphSpec `json:"spec"`
}

type ClusterGraphSpec struct {
	ClusterName string `json:"clusterName" norman:"type=reference[cluster]"`
}

type ProjectGraphSpec struct {
	ProjectName string `json:"projectName" yaml:"projectName" norman:"required,type=reference[project]"`
}

type MonitoringStatus struct {
	GrafanaEndpoint string                `json:"grafanaEndpoint,omitempty"`
	Conditions      []MonitoringCondition `json:"conditions,omitempty"`
}

type MonitoringCondition struct {
	// Type of cluster condition.
	Type ClusterConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition
	Message string `json:"message,omitempty"`
}

const (
	MonitoringConditionGrafanaDeployed           condition.Cond = "GrafanaDeployed"
	MonitoringConditionPrometheusDeployed        condition.Cond = "PrometheusDeployed"
	MonitoringConditionAlertmaanagerDeployed     condition.Cond = "AlertmanagerDeployed"
	MonitoringConditionNodeExporterDeployed      condition.Cond = "NodeExporterDeployed"
	MonitoringConditionKubeStateExporterDeployed condition.Cond = "KubeStateExporterDeployed"
)
