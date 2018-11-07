package client

import (
	"github.com/rancher/norman/clientbase"
)

type Client struct {
	clientbase.APIBaseClient

	Namespace        NamespaceOperations
	PersistentVolume PersistentVolumeOperations
	StorageClass     StorageClassOperations
	MonitorMetric    MonitorMetricOperations
	MonitorGraph     MonitorGraphOperations
	ClusterGraph     ClusterGraphOperations
	ProjectGraph     ProjectGraphOperations
}

func NewClient(opts *clientbase.ClientOpts) (*Client, error) {
	baseClient, err := clientbase.NewAPIClient(opts)
	if err != nil {
		return nil, err
	}

	client := &Client{
		APIBaseClient: baseClient,
	}

	client.Namespace = newNamespaceClient(client)
	client.PersistentVolume = newPersistentVolumeClient(client)
	client.StorageClass = newStorageClassClient(client)
	client.MonitorMetric = newMonitorMetricClient(client)
	client.MonitorGraph = newMonitorGraphClient(client)
	client.ClusterGraph = newClusterGraphClient(client)
	client.ProjectGraph = newProjectGraphClient(client)

	return client, nil
}
