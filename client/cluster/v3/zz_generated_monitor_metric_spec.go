package client

const (
	MonitorMetricSpecType                = "monitorMetricSpec"
	MonitorMetricSpecFieldDescription    = "description"
	MonitorMetricSpecFieldDetailsGroupBy = "detailsGroupBy"
	MonitorMetricSpecFieldExpression     = "expression"
	MonitorMetricSpecFieldExtraAddedTags = "extraAddedTags"
	MonitorMetricSpecFieldLegendFormat   = "legendFormat"
	MonitorMetricSpecFieldStep           = "step"
)

type MonitorMetricSpec struct {
	Description    string            `json:"description,omitempty" yaml:"description,omitempty"`
	DetailsGroupBy string            `json:"detailsGroupBy,omitempty" yaml:"detailsGroupBy,omitempty"`
	Expression     string            `json:"expression,omitempty" yaml:"expression,omitempty"`
	ExtraAddedTags map[string]string `json:"extraAddedTags,omitempty" yaml:"extraAddedTags,omitempty"`
	LegendFormat   string            `json:"legendFormat,omitempty" yaml:"legendFormat,omitempty"`
	Step           int64             `json:"step,omitempty" yaml:"step,omitempty"`
}
