package schema

import (
	"github.com/rancher/norman/types"
	m "github.com/rancher/norman/types/mapper"
	monitoringv1 "github.com/rancher/types/apis/monitoring.cattle.io/v1"
	"github.com/rancher/types/factory"
)

var (
	Version = types.APIVersion{
		Version: "v1",
		Group:   "monitoring.cattle.io",
		Path:    "/v3/monitoring",
	}
	Schema = factory.Schemas(&Version).
		Init(MonitorTypes)
)

func MonitorTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, monitoringv1.StorageSpec{},
			&m.Drop{Field: "class"},
			&m.Drop{Field: "selector"},
			&m.Drop{Field: "resources"},
		).
		AddMapperForType(&Version, monitoringv1.Prometheus{},
			&m.Drop{Field: "status"},
			&m.AnnotationField{Field: "description"},
		).
		AddMapperForType(&Version, monitoringv1.PrometheusSpec{},
			&m.Drop{Field: "thanos"},
			&m.Drop{Field: "apiserverConfig"},
			&m.Drop{Field: "serviceMonitorNamespaceSelector"},
			&m.Drop{Field: "ruleNamespaceSelector"},
			&m.Drop{Field: "paused"},
			&m.Enum{
				Field: "logLevel",
				Options: []string{
					"all",
					"debug",
					"info",
					"warn",
					"error",
					"none",
				},
			},
		).
		MustImportAndCustomize(&Version, monitoringv1.Prometheus{}, func(schema *types.Schema) {
			schema.MustCustomizeField("name", func(field types.Field) types.Field {
				field.Type = "dnsLabelRestricted"
				field.Nullable = false
				field.Required = true
				return field
			})
		}, projectOverride{}, struct {
			Description string `json:"description"`
		}{}).
		AddMapperForType(&Version, monitoringv1.RelabelConfig{},
			&m.Enum{
				Field: "action",
				Options: []string{
					"replace",
					"keep",
					"drop",
					"hashmod",
					"labelmap",
					"labeldrop",
					"labelkeep",
				},
			},
		).
		AddMapperForType(&Version, monitoringv1.Endpoint{},
			&m.Drop{Field: "port"},
			&m.Drop{Field: "tlsConfig"},
			&m.Drop{Field: "bearerTokenFile"},
			&m.Drop{Field: "honorLabels"},
			&m.Drop{Field: "basicAuth"},
			&m.Drop{Field: "metricRelabelings"},
			&m.Drop{Field: "proxyUrl"},
		).
		AddMapperForType(&Version, monitoringv1.ServiceMonitorSpec{},
			&m.Embed{Field: "namespaceSelector"},
			&m.Drop{Field: "any"},
			&m.Move{From: "matchNames", To: "namespaceSelector"},
		).
		AddMapperForType(&Version, monitoringv1.ServiceMonitor{},
			&m.AnnotationField{Field: "displayName"},
			&m.DisplayName{},
			&m.AnnotationField{Field: "targetService"},
			&m.AnnotationField{Field: "targetWorkload"},
		).
		MustImport(&Version, monitoringv1.ServiceMonitor{}, projectOverride{}, struct {
			DisplayName    string `json:"displayName,omitempty"`
			TargetService  string `json:"targetService,omitempty"`
			TargetWorkload string `json:"targetWorkload,omitempty"`
		}{}).
		MustImport(&Version, monitoringv1.PrometheusRule{}, projectOverride{}).
		AddMapperForType(&Version, monitoringv1.Alertmanager{},
			&m.Drop{Field: "status"},
		).
		MustImport(&Version, monitoringv1.Alertmanager{}, projectOverride{})
}
