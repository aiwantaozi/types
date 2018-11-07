package v3

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	MonitorMetricGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "MonitorMetric",
	}
	MonitorMetricResource = metav1.APIResource{
		Name:         "monitormetrics",
		SingularName: "monitormetric",
		Namespaced:   true,

		Kind: MonitorMetricGroupVersionKind.Kind,
	}
)

type MonitorMetricList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorMetric
}

type MonitorMetricHandlerFunc func(key string, obj *MonitorMetric) error

type MonitorMetricLister interface {
	List(namespace string, selector labels.Selector) (ret []*MonitorMetric, err error)
	Get(namespace, name string) (*MonitorMetric, error)
}

type MonitorMetricController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() MonitorMetricLister
	AddHandler(name string, handler MonitorMetricHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler MonitorMetricHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type MonitorMetricInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*MonitorMetric) (*MonitorMetric, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*MonitorMetric, error)
	Get(name string, opts metav1.GetOptions) (*MonitorMetric, error)
	Update(*MonitorMetric) (*MonitorMetric, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*MonitorMetricList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() MonitorMetricController
	AddHandler(name string, sync MonitorMetricHandlerFunc)
	AddLifecycle(name string, lifecycle MonitorMetricLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync MonitorMetricHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle MonitorMetricLifecycle)
}

type monitorMetricLister struct {
	controller *monitorMetricController
}

func (l *monitorMetricLister) List(namespace string, selector labels.Selector) (ret []*MonitorMetric, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*MonitorMetric))
	})
	return
}

func (l *monitorMetricLister) Get(namespace, name string) (*MonitorMetric, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    MonitorMetricGroupVersionKind.Group,
			Resource: "monitorMetric",
		}, key)
	}
	return obj.(*MonitorMetric), nil
}

type monitorMetricController struct {
	controller.GenericController
}

func (c *monitorMetricController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *monitorMetricController) Lister() MonitorMetricLister {
	return &monitorMetricLister{
		controller: c,
	}
}

func (c *monitorMetricController) AddHandler(name string, handler MonitorMetricHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*MonitorMetric))
	})
}

func (c *monitorMetricController) AddClusterScopedHandler(name, cluster string, handler MonitorMetricHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}

		if !controller.ObjectInCluster(cluster, obj) {
			return nil
		}

		return handler(key, obj.(*MonitorMetric))
	})
}

type monitorMetricFactory struct {
}

func (c monitorMetricFactory) Object() runtime.Object {
	return &MonitorMetric{}
}

func (c monitorMetricFactory) List() runtime.Object {
	return &MonitorMetricList{}
}

func (s *monitorMetricClient) Controller() MonitorMetricController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.monitorMetricControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(MonitorMetricGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &monitorMetricController{
		GenericController: genericController,
	}

	s.client.monitorMetricControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type monitorMetricClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   MonitorMetricController
}

func (s *monitorMetricClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *monitorMetricClient) Create(o *MonitorMetric) (*MonitorMetric, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*MonitorMetric), err
}

func (s *monitorMetricClient) Get(name string, opts metav1.GetOptions) (*MonitorMetric, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*MonitorMetric), err
}

func (s *monitorMetricClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*MonitorMetric, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*MonitorMetric), err
}

func (s *monitorMetricClient) Update(o *MonitorMetric) (*MonitorMetric, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*MonitorMetric), err
}

func (s *monitorMetricClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *monitorMetricClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *monitorMetricClient) List(opts metav1.ListOptions) (*MonitorMetricList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*MonitorMetricList), err
}

func (s *monitorMetricClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *monitorMetricClient) Patch(o *MonitorMetric, data []byte, subresources ...string) (*MonitorMetric, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*MonitorMetric), err
}

func (s *monitorMetricClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *monitorMetricClient) AddHandler(name string, sync MonitorMetricHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *monitorMetricClient) AddLifecycle(name string, lifecycle MonitorMetricLifecycle) {
	sync := NewMonitorMetricLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *monitorMetricClient) AddClusterScopedHandler(name, clusterName string, sync MonitorMetricHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *monitorMetricClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle MonitorMetricLifecycle) {
	sync := NewMonitorMetricLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
