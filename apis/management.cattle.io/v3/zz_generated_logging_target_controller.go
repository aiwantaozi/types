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
	LoggingTargetGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "LoggingTarget",
	}
	LoggingTargetResource = metav1.APIResource{
		Name:         "loggingtargets",
		SingularName: "loggingtarget",
		Namespaced:   false,
		Kind:         LoggingTargetGroupVersionKind.Kind,
	}
)

type LoggingTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoggingTarget
}

type LoggingTargetHandlerFunc func(key string, obj *LoggingTarget) error

type LoggingTargetLister interface {
	List(namespace string, selector labels.Selector) (ret []*LoggingTarget, err error)
	Get(namespace, name string) (*LoggingTarget, error)
}

type LoggingTargetController interface {
	Informer() cache.SharedIndexInformer
	Lister() LoggingTargetLister
	AddHandler(name string, handler LoggingTargetHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler LoggingTargetHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type LoggingTargetInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*LoggingTarget) (*LoggingTarget, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*LoggingTarget, error)
	Get(name string, opts metav1.GetOptions) (*LoggingTarget, error)
	Update(*LoggingTarget) (*LoggingTarget, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*LoggingTargetList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() LoggingTargetController
	AddHandler(name string, sync LoggingTargetHandlerFunc)
	AddLifecycle(name string, lifecycle LoggingTargetLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync LoggingTargetHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle LoggingTargetLifecycle)
}

type loggingTargetLister struct {
	controller *loggingTargetController
}

func (l *loggingTargetLister) List(namespace string, selector labels.Selector) (ret []*LoggingTarget, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*LoggingTarget))
	})
	return
}

func (l *loggingTargetLister) Get(namespace, name string) (*LoggingTarget, error) {
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
			Group:    LoggingTargetGroupVersionKind.Group,
			Resource: "loggingTarget",
		}, name)
	}
	return obj.(*LoggingTarget), nil
}

type loggingTargetController struct {
	controller.GenericController
}

func (c *loggingTargetController) Lister() LoggingTargetLister {
	return &loggingTargetLister{
		controller: c,
	}
}

func (c *loggingTargetController) AddHandler(name string, handler LoggingTargetHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*LoggingTarget))
	})
}

func (c *loggingTargetController) AddClusterScopedHandler(name, cluster string, handler LoggingTargetHandlerFunc) {
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

		return handler(key, obj.(*LoggingTarget))
	})
}

type loggingTargetFactory struct {
}

func (c loggingTargetFactory) Object() runtime.Object {
	return &LoggingTarget{}
}

func (c loggingTargetFactory) List() runtime.Object {
	return &LoggingTargetList{}
}

func (s *loggingTargetClient) Controller() LoggingTargetController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.loggingTargetControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(LoggingTargetGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &loggingTargetController{
		GenericController: genericController,
	}

	s.client.loggingTargetControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type loggingTargetClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   LoggingTargetController
}

func (s *loggingTargetClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *loggingTargetClient) Create(o *LoggingTarget) (*LoggingTarget, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*LoggingTarget), err
}

func (s *loggingTargetClient) Get(name string, opts metav1.GetOptions) (*LoggingTarget, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*LoggingTarget), err
}

func (s *loggingTargetClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*LoggingTarget, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*LoggingTarget), err
}

func (s *loggingTargetClient) Update(o *LoggingTarget) (*LoggingTarget, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*LoggingTarget), err
}

func (s *loggingTargetClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *loggingTargetClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *loggingTargetClient) List(opts metav1.ListOptions) (*LoggingTargetList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*LoggingTargetList), err
}

func (s *loggingTargetClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *loggingTargetClient) Patch(o *LoggingTarget, data []byte, subresources ...string) (*LoggingTarget, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*LoggingTarget), err
}

func (s *loggingTargetClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *loggingTargetClient) AddHandler(name string, sync LoggingTargetHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *loggingTargetClient) AddLifecycle(name string, lifecycle LoggingTargetLifecycle) {
	sync := NewLoggingTargetLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *loggingTargetClient) AddClusterScopedHandler(name, clusterName string, sync LoggingTargetHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *loggingTargetClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle LoggingTargetLifecycle) {
	sync := NewLoggingTargetLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
