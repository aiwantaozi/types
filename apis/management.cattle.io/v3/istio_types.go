package v3

import (
	"github.com/rancher/norman/condition"
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	IstioConditionCertManagerDeployed            condition.Cond = "CertManagerDeployed"
	IstioConditionSecurityDeployed               condition.Cond = "SecurityDeployed"
	IstioConditionSidecarInjectorWebhookDeployed condition.Cond = "SidecarInjectorWebhookDeployed"
	IstioConditionTelemetryGatewayDeployed       condition.Cond = "TelemetryGatewayDeployed"
	IstioConditionTracingDeployed                condition.Cond = "TracingDeployed"
	IstioConditionPrometheusDeployedDeployed     condition.Cond = "PrometheusDeployed"
	IstioConditionPilotDeployedDeployed          condition.Cond = "PilotDeployed"
	IstioConditionMixerDeployedDeployed          condition.Cond = "MixerDeployed"
	IstioConditionKialiDeployedDeployed          condition.Cond = "KialiDeployed"
	IstioConditionIngressDeployedDeployed        condition.Cond = "IngressDeployed"
	IstioConditionGrafanaDeployedDeployed        condition.Cond = "GrafanaDeployed"
	IstioConditionGatewaysDeployedDeployed       condition.Cond = "GatewaysDeployed"
	IstioConditionGalleyDeployedDeployed         condition.Cond = "GalleyDeployed"
)

type IstioConfig struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec IstioConfigSpec `json:"spec"`

	Status IstioConfigStatus `json:"status"`
}

type IstioConfigSpec struct {
	ClusterName string `json:"clusterName" norman:"type=reference[cluster]"`
	Enable      bool   `json:"enable,omitempty" norman:"required,default=false"`
}

type IstioConfigStatus struct {
	Conditions []Condition `json:"conditions"`
}
