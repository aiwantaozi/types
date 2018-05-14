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
	GlobalLoggingGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "GlobalLogging",
	}
	GlobalLoggingResource = metav1.APIResource{
		Name:         "globalloggings",
		SingularName: "globallogging",
		Namespaced:   false,
		Kind:         GlobalLoggingGroupVersionKind.Kind,
	}
)

type GlobalLoggingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalLogging
}

type GlobalLoggingHandlerFunc func(key string, obj *GlobalLogging) error

type GlobalLoggingLister interface {
	List(namespace string, selector labels.Selector) (ret []*GlobalLogging, err error)
	Get(namespace, name string) (*GlobalLogging, error)
}

type GlobalLoggingController interface {
	Informer() cache.SharedIndexInformer
	Lister() GlobalLoggingLister
	AddHandler(name string, handler GlobalLoggingHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler GlobalLoggingHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type GlobalLoggingInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*GlobalLogging) (*GlobalLogging, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*GlobalLogging, error)
	Get(name string, opts metav1.GetOptions) (*GlobalLogging, error)
	Update(*GlobalLogging) (*GlobalLogging, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*GlobalLoggingList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() GlobalLoggingController
	AddHandler(name string, sync GlobalLoggingHandlerFunc)
	AddLifecycle(name string, lifecycle GlobalLoggingLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync GlobalLoggingHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle GlobalLoggingLifecycle)
}

type globalLoggingLister struct {
	controller *globalLoggingController
}

func (l *globalLoggingLister) List(namespace string, selector labels.Selector) (ret []*GlobalLogging, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*GlobalLogging))
	})
	return
}

func (l *globalLoggingLister) Get(namespace, name string) (*GlobalLogging, error) {
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
			Group:    GlobalLoggingGroupVersionKind.Group,
			Resource: "globalLogging",
		}, name)
	}
	return obj.(*GlobalLogging), nil
}

type globalLoggingController struct {
	controller.GenericController
}

func (c *globalLoggingController) Lister() GlobalLoggingLister {
	return &globalLoggingLister{
		controller: c,
	}
}

func (c *globalLoggingController) AddHandler(name string, handler GlobalLoggingHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*GlobalLogging))
	})
}

func (c *globalLoggingController) AddClusterScopedHandler(name, cluster string, handler GlobalLoggingHandlerFunc) {
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

		return handler(key, obj.(*GlobalLogging))
	})
}

type globalLoggingFactory struct {
}

func (c globalLoggingFactory) Object() runtime.Object {
	return &GlobalLogging{}
}

func (c globalLoggingFactory) List() runtime.Object {
	return &GlobalLoggingList{}
}

func (s *globalLoggingClient) Controller() GlobalLoggingController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.globalLoggingControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(GlobalLoggingGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &globalLoggingController{
		GenericController: genericController,
	}

	s.client.globalLoggingControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type globalLoggingClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   GlobalLoggingController
}

func (s *globalLoggingClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *globalLoggingClient) Create(o *GlobalLogging) (*GlobalLogging, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*GlobalLogging), err
}

func (s *globalLoggingClient) Get(name string, opts metav1.GetOptions) (*GlobalLogging, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*GlobalLogging), err
}

func (s *globalLoggingClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*GlobalLogging, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*GlobalLogging), err
}

func (s *globalLoggingClient) Update(o *GlobalLogging) (*GlobalLogging, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*GlobalLogging), err
}

func (s *globalLoggingClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *globalLoggingClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *globalLoggingClient) List(opts metav1.ListOptions) (*GlobalLoggingList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*GlobalLoggingList), err
}

func (s *globalLoggingClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *globalLoggingClient) Patch(o *GlobalLogging, data []byte, subresources ...string) (*GlobalLogging, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*GlobalLogging), err
}

func (s *globalLoggingClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *globalLoggingClient) AddHandler(name string, sync GlobalLoggingHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *globalLoggingClient) AddLifecycle(name string, lifecycle GlobalLoggingLifecycle) {
	sync := NewGlobalLoggingLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *globalLoggingClient) AddClusterScopedHandler(name, clusterName string, sync GlobalLoggingHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *globalLoggingClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle GlobalLoggingLifecycle) {
	sync := NewGlobalLoggingLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
