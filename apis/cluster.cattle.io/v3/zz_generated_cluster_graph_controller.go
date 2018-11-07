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
	ClusterGraphGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ClusterGraph",
	}
	ClusterGraphResource = metav1.APIResource{
		Name:         "clustergraphs",
		SingularName: "clustergraph",
		Namespaced:   true,

		Kind: ClusterGraphGroupVersionKind.Kind,
	}
)

type ClusterGraphList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterGraph
}

type ClusterGraphHandlerFunc func(key string, obj *ClusterGraph) error

type ClusterGraphLister interface {
	List(namespace string, selector labels.Selector) (ret []*ClusterGraph, err error)
	Get(namespace, name string) (*ClusterGraph, error)
}

type ClusterGraphController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() ClusterGraphLister
	AddHandler(name string, handler ClusterGraphHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler ClusterGraphHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ClusterGraphInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*ClusterGraph) (*ClusterGraph, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ClusterGraph, error)
	Get(name string, opts metav1.GetOptions) (*ClusterGraph, error)
	Update(*ClusterGraph) (*ClusterGraph, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ClusterGraphList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ClusterGraphController
	AddHandler(name string, sync ClusterGraphHandlerFunc)
	AddLifecycle(name string, lifecycle ClusterGraphLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync ClusterGraphHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle ClusterGraphLifecycle)
}

type clusterGraphLister struct {
	controller *clusterGraphController
}

func (l *clusterGraphLister) List(namespace string, selector labels.Selector) (ret []*ClusterGraph, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*ClusterGraph))
	})
	return
}

func (l *clusterGraphLister) Get(namespace, name string) (*ClusterGraph, error) {
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
			Group:    ClusterGraphGroupVersionKind.Group,
			Resource: "clusterGraph",
		}, key)
	}
	return obj.(*ClusterGraph), nil
}

type clusterGraphController struct {
	controller.GenericController
}

func (c *clusterGraphController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *clusterGraphController) Lister() ClusterGraphLister {
	return &clusterGraphLister{
		controller: c,
	}
}

func (c *clusterGraphController) AddHandler(name string, handler ClusterGraphHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*ClusterGraph))
	})
}

func (c *clusterGraphController) AddClusterScopedHandler(name, cluster string, handler ClusterGraphHandlerFunc) {
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

		return handler(key, obj.(*ClusterGraph))
	})
}

type clusterGraphFactory struct {
}

func (c clusterGraphFactory) Object() runtime.Object {
	return &ClusterGraph{}
}

func (c clusterGraphFactory) List() runtime.Object {
	return &ClusterGraphList{}
}

func (s *clusterGraphClient) Controller() ClusterGraphController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.clusterGraphControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ClusterGraphGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &clusterGraphController{
		GenericController: genericController,
	}

	s.client.clusterGraphControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type clusterGraphClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ClusterGraphController
}

func (s *clusterGraphClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *clusterGraphClient) Create(o *ClusterGraph) (*ClusterGraph, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ClusterGraph), err
}

func (s *clusterGraphClient) Get(name string, opts metav1.GetOptions) (*ClusterGraph, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ClusterGraph), err
}

func (s *clusterGraphClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ClusterGraph, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*ClusterGraph), err
}

func (s *clusterGraphClient) Update(o *ClusterGraph) (*ClusterGraph, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ClusterGraph), err
}

func (s *clusterGraphClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *clusterGraphClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *clusterGraphClient) List(opts metav1.ListOptions) (*ClusterGraphList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ClusterGraphList), err
}

func (s *clusterGraphClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *clusterGraphClient) Patch(o *ClusterGraph, data []byte, subresources ...string) (*ClusterGraph, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*ClusterGraph), err
}

func (s *clusterGraphClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *clusterGraphClient) AddHandler(name string, sync ClusterGraphHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *clusterGraphClient) AddLifecycle(name string, lifecycle ClusterGraphLifecycle) {
	sync := NewClusterGraphLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *clusterGraphClient) AddClusterScopedHandler(name, clusterName string, sync ClusterGraphHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *clusterGraphClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle ClusterGraphLifecycle) {
	sync := NewClusterGraphLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
