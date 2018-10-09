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
		Path:    "/v1",
	}

	Schemas = factory.Schemas(&Version).
		Init(monitoringTypes)
)

func monitoringTypes(schemas *types.Schemas) *types.Schemas {
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
		AddMapperForType(&Version, monitoringv1.Alertmanager{},
			&m.Drop{Field: "status"},
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
		MustImport(&Version, monitoringv1.PrometheusRule{}, projectOverride{}).
		MustImport(&Version, monitoringv1.Alertmanager{}, projectOverride{}).
		MustImport(&Version, monitoringv1.ServiceMonitor{}, projectOverride{})
}
