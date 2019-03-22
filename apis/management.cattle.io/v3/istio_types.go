package v3

import (
	"github.com/rancher/norman/condition"
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	IstioConditionCertManagerDeployed            condition.Cond = "CertManagerDeployed"
	IstioConditionGalleyDeployedDeployed         condition.Cond = "GalleyDeployed"
	IstioConditionGatewaysDeployedDeployed       condition.Cond = "GatewaysDeployed"
	IstioConditionGrafanaDeployedDeployed        condition.Cond = "GrafanaDeployed"
	IstioConditionKialiDeployedDeployed          condition.Cond = "KialiDeployed"
	IstioConditionMixerDeployedDeployed          condition.Cond = "MixerDeployed"
	IstioConditionPilotDeployedDeployed          condition.Cond = "PilotDeployed"
	IstioConditionPrometheusDeployedDeployed     condition.Cond = "PrometheusDeployed"
	IstioConditionSecurityDeployed               condition.Cond = "SecurityDeployed"
	IstioConditionServiceGraphDeployed           condition.Cond = "ServiceGraphDeployed"
	IstioConditionSidecarInjectorWebhookDeployed condition.Cond = "SidecarInjectorWebhookDeployed"
	IstioConditionTracingDeployed                condition.Cond = "TracingDeployed"

	// istio 1.1 added
	IstioConditionCoreDNSDeployed   condition.Cond = "CoreDNSDeployed"
	IstioConditionNodeAgentDeployed condition.Cond = "NodeAgentDeployed"
	// todo: istio 1.1 removed
	IstioConditionIngressDeployedDeployed  condition.Cond = "IngressDeployed"
	IstioConditionTelemetryGatewayDeployed condition.Cond = "TelemetryGatewayDeployed"
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
	ClusterName string            `json:"clusterName" norman:"type=reference[cluster]"`
	Enable      bool              `json:"enable,omitempty" norman:"required,default=false"`
	AppAnswers  map[string]string `json:"answers,omitempty"`
}

type IstioConfigStatus struct {
	Conditions []Condition `json:"conditions"`
}
