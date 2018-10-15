package client

const (
	ServiceMonitorSpecType                          = "serviceMonitorSpec"
	ServiceMonitorSpecFieldEndpoints                = "endpoints"
	ServiceMonitorSpecFieldJobLabelOverWrite        = "jobLabelOverWrite"
	ServiceMonitorSpecFieldNamespaceSelector        = "namespaceSelector"
	ServiceMonitorSpecFieldSelector                 = "selector"
	ServiceMonitorSpecFieldTransparentServiceLabels = "transparentServiceLabels"
)

type ServiceMonitorSpec struct {
	Endpoints                []Endpoint     `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
	JobLabelOverWrite        string         `json:"jobLabelOverWrite,omitempty" yaml:"jobLabelOverWrite,omitempty"`
	NamespaceSelector        []string       `json:"namespaceSelector,omitempty" yaml:"namespaceSelector,omitempty"`
	Selector                 *LabelSelector `json:"selector,omitempty" yaml:"selector,omitempty"`
	TransparentServiceLabels []string       `json:"transparentServiceLabels,omitempty" yaml:"transparentServiceLabels,omitempty"`
}
