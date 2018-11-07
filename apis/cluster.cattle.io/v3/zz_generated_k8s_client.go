package v3

import (
	"context"
	"sync"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"github.com/rancher/norman/objectclient/dynamic"
	"github.com/rancher/norman/restwatch"
	"k8s.io/client-go/rest"
)

type Interface interface {
	RESTClient() rest.Interface
	controller.Starter

	MonitorMetricsGetter
	MonitorGraphsGetter
	ClusterGraphsGetter
	ProjectGraphsGetter
}

type Client struct {
	sync.Mutex
	restClient rest.Interface
	starters   []controller.Starter

	monitorMetricControllers map[string]MonitorMetricController
	monitorGraphControllers  map[string]MonitorGraphController
	clusterGraphControllers  map[string]ClusterGraphController
	projectGraphControllers  map[string]ProjectGraphController
}

func NewForConfig(config rest.Config) (Interface, error) {
	if config.NegotiatedSerializer == nil {
		config.NegotiatedSerializer = dynamic.NegotiatedSerializer
	}

	restClient, err := restwatch.UnversionedRESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &Client{
		restClient: restClient,

		monitorMetricControllers: map[string]MonitorMetricController{},
		monitorGraphControllers:  map[string]MonitorGraphController{},
		clusterGraphControllers:  map[string]ClusterGraphController{},
		projectGraphControllers:  map[string]ProjectGraphController{},
	}, nil
}

func (c *Client) RESTClient() rest.Interface {
	return c.restClient
}

func (c *Client) Sync(ctx context.Context) error {
	return controller.Sync(ctx, c.starters...)
}

func (c *Client) Start(ctx context.Context, threadiness int) error {
	return controller.Start(ctx, threadiness, c.starters...)
}

type MonitorMetricsGetter interface {
	MonitorMetrics(namespace string) MonitorMetricInterface
}

func (c *Client) MonitorMetrics(namespace string) MonitorMetricInterface {
	objectClient := objectclient.NewObjectClient(namespace, c.restClient, &MonitorMetricResource, MonitorMetricGroupVersionKind, monitorMetricFactory{})
	return &monitorMetricClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type MonitorGraphsGetter interface {
	MonitorGraphs(namespace string) MonitorGraphInterface
}

func (c *Client) MonitorGraphs(namespace string) MonitorGraphInterface {
	objectClient := objectclient.NewObjectClient(namespace, c.restClient, &MonitorGraphResource, MonitorGraphGroupVersionKind, monitorGraphFactory{})
	return &monitorGraphClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type ClusterGraphsGetter interface {
	ClusterGraphs(namespace string) ClusterGraphInterface
}

func (c *Client) ClusterGraphs(namespace string) ClusterGraphInterface {
	objectClient := objectclient.NewObjectClient(namespace, c.restClient, &ClusterGraphResource, ClusterGraphGroupVersionKind, clusterGraphFactory{})
	return &clusterGraphClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type ProjectGraphsGetter interface {
	ProjectGraphs(namespace string) ProjectGraphInterface
}

func (c *Client) ProjectGraphs(namespace string) ProjectGraphInterface {
	objectClient := objectclient.NewObjectClient(namespace, c.restClient, &ProjectGraphResource, ProjectGraphGroupVersionKind, projectGraphFactory{})
	return &projectGraphClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
