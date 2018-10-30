package client

import (
	"github.com/rancher/norman/clientbase"
)

type Client struct {
	clientbase.APIBaseClient

	PersistentVolumeClaim PersistentVolumeClaimOperations
	Prometheus            PrometheusOperations
	ServiceMonitor        ServiceMonitorOperations
	PrometheusRule        PrometheusRuleOperations
	Alertmanager          AlertmanagerOperations
}

func NewClient(opts *clientbase.ClientOpts) (*Client, error) {
	baseClient, err := clientbase.NewAPIClient(opts)
	if err != nil {
		return nil, err
	}

	client := &Client{
		APIBaseClient: baseClient,
	}

	client.PersistentVolumeClaim = newPersistentVolumeClaimClient(client)
	client.Prometheus = newPrometheusClient(client)
	client.ServiceMonitor = newServiceMonitorClient(client)
	client.PrometheusRule = newPrometheusRuleClient(client)
	client.Alertmanager = newAlertmanagerClient(client)

	return client, nil
}
