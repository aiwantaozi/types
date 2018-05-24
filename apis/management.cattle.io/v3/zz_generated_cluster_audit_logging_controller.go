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
	ClusterAuditLoggingGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ClusterAuditLogging",
	}
	ClusterAuditLoggingResource = metav1.APIResource{
		Name:         "clusterauditloggings",
		SingularName: "clusterauditlogging",
		Namespaced:   true,

		Kind: ClusterAuditLoggingGroupVersionKind.Kind,
	}
)

type ClusterAuditLoggingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterAuditLogging
}

type ClusterAuditLoggingHandlerFunc func(key string, obj *ClusterAuditLogging) error

type ClusterAuditLoggingLister interface {
	List(namespace string, selector labels.Selector) (ret []*ClusterAuditLogging, err error)
	Get(namespace, name string) (*ClusterAuditLogging, error)
}

type ClusterAuditLoggingController interface {
	Informer() cache.SharedIndexInformer
	Lister() ClusterAuditLoggingLister
	AddHandler(name string, handler ClusterAuditLoggingHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler ClusterAuditLoggingHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ClusterAuditLoggingInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*ClusterAuditLogging) (*ClusterAuditLogging, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ClusterAuditLogging, error)
	Get(name string, opts metav1.GetOptions) (*ClusterAuditLogging, error)
	Update(*ClusterAuditLogging) (*ClusterAuditLogging, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ClusterAuditLoggingList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ClusterAuditLoggingController
	AddHandler(name string, sync ClusterAuditLoggingHandlerFunc)
	AddLifecycle(name string, lifecycle ClusterAuditLoggingLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync ClusterAuditLoggingHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle ClusterAuditLoggingLifecycle)
}

type clusterAuditLoggingLister struct {
	controller *clusterAuditLoggingController
}

func (l *clusterAuditLoggingLister) List(namespace string, selector labels.Selector) (ret []*ClusterAuditLogging, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*ClusterAuditLogging))
	})
	return
}

func (l *clusterAuditLoggingLister) Get(namespace, name string) (*ClusterAuditLogging, error) {
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
			Group:    ClusterAuditLoggingGroupVersionKind.Group,
			Resource: "clusterAuditLogging",
		}, name)
	}
	return obj.(*ClusterAuditLogging), nil
}

type clusterAuditLoggingController struct {
	controller.GenericController
}

func (c *clusterAuditLoggingController) Lister() ClusterAuditLoggingLister {
	return &clusterAuditLoggingLister{
		controller: c,
	}
}

func (c *clusterAuditLoggingController) AddHandler(name string, handler ClusterAuditLoggingHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*ClusterAuditLogging))
	})
}

func (c *clusterAuditLoggingController) AddClusterScopedHandler(name, cluster string, handler ClusterAuditLoggingHandlerFunc) {
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

		return handler(key, obj.(*ClusterAuditLogging))
	})
}

type clusterAuditLoggingFactory struct {
}

func (c clusterAuditLoggingFactory) Object() runtime.Object {
	return &ClusterAuditLogging{}
}

func (c clusterAuditLoggingFactory) List() runtime.Object {
	return &ClusterAuditLoggingList{}
}

func (s *clusterAuditLoggingClient) Controller() ClusterAuditLoggingController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.clusterAuditLoggingControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ClusterAuditLoggingGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &clusterAuditLoggingController{
		GenericController: genericController,
	}

	s.client.clusterAuditLoggingControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type clusterAuditLoggingClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ClusterAuditLoggingController
}

func (s *clusterAuditLoggingClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *clusterAuditLoggingClient) Create(o *ClusterAuditLogging) (*ClusterAuditLogging, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ClusterAuditLogging), err
}

func (s *clusterAuditLoggingClient) Get(name string, opts metav1.GetOptions) (*ClusterAuditLogging, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ClusterAuditLogging), err
}

func (s *clusterAuditLoggingClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ClusterAuditLogging, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*ClusterAuditLogging), err
}

func (s *clusterAuditLoggingClient) Update(o *ClusterAuditLogging) (*ClusterAuditLogging, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ClusterAuditLogging), err
}

func (s *clusterAuditLoggingClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *clusterAuditLoggingClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *clusterAuditLoggingClient) List(opts metav1.ListOptions) (*ClusterAuditLoggingList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ClusterAuditLoggingList), err
}

func (s *clusterAuditLoggingClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *clusterAuditLoggingClient) Patch(o *ClusterAuditLogging, data []byte, subresources ...string) (*ClusterAuditLogging, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*ClusterAuditLogging), err
}

func (s *clusterAuditLoggingClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *clusterAuditLoggingClient) AddHandler(name string, sync ClusterAuditLoggingHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *clusterAuditLoggingClient) AddLifecycle(name string, lifecycle ClusterAuditLoggingLifecycle) {
	sync := NewClusterAuditLoggingLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *clusterAuditLoggingClient) AddClusterScopedHandler(name, clusterName string, sync ClusterAuditLoggingHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *clusterAuditLoggingClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle ClusterAuditLoggingLifecycle) {
	sync := NewClusterAuditLoggingLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
