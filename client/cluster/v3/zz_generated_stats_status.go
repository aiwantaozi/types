package client

const (
	StatsStatusType                 = "statsStatus"
	StatsStatusFieldMonitorGraphIDs = "monitorGraphIds"
)

type StatsStatus struct {
	MonitorGraphIDs []string `json:"monitorGraphIds,omitempty" yaml:"monitorGraphIds,omitempty"`
}
