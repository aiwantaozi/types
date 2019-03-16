package client

const (
	EnvoyFilterType                = "envoyFilter"
	EnvoyFilterFieldFilters        = "filters"
	EnvoyFilterFieldWorkloadLabels = "workload_labels"
)

type EnvoyFilter struct {
	Filters        []EnvoyFilter_Filter `json:"filters,omitempty" yaml:"filters,omitempty"`
	WorkloadLabels map[string]string    `json:"workload_labels,omitempty" yaml:"workload_labels,omitempty"`
}
