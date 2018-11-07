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
	ProjectGraphGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ProjectGraph",
	}
	ProjectGraphResource = metav1.APIResource{
		Name:         "projectgraphs",
		SingularName: "projectgraph",
		Namespaced:   true,

		Kind: ProjectGraphGroupVersionKind.Kind,
	}
)

type ProjectGraphList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectGraph
}

type ProjectGraphHandlerFunc func(key string, obj *ProjectGraph) error

type ProjectGraphLister interface {
	List(namespace string, selector labels.Selector) (ret []*ProjectGraph, err error)
	Get(namespace, name string) (*ProjectGraph, error)
}

type ProjectGraphController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() ProjectGraphLister
	AddHandler(name string, handler ProjectGraphHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler ProjectGraphHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ProjectGraphInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*ProjectGraph) (*ProjectGraph, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ProjectGraph, error)
	Get(name string, opts metav1.GetOptions) (*ProjectGraph, error)
	Update(*ProjectGraph) (*ProjectGraph, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ProjectGraphList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ProjectGraphController
	AddHandler(name string, sync ProjectGraphHandlerFunc)
	AddLifecycle(name string, lifecycle ProjectGraphLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync ProjectGraphHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle ProjectGraphLifecycle)
}

type projectGraphLister struct {
	controller *projectGraphController
}

func (l *projectGraphLister) List(namespace string, selector labels.Selector) (ret []*ProjectGraph, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*ProjectGraph))
	})
	return
}

func (l *projectGraphLister) Get(namespace, name string) (*ProjectGraph, error) {
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
			Group:    ProjectGraphGroupVersionKind.Group,
			Resource: "projectGraph",
		}, key)
	}
	return obj.(*ProjectGraph), nil
}

type projectGraphController struct {
	controller.GenericController
}

func (c *projectGraphController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *projectGraphController) Lister() ProjectGraphLister {
	return &projectGraphLister{
		controller: c,
	}
}

func (c *projectGraphController) AddHandler(name string, handler ProjectGraphHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*ProjectGraph))
	})
}

func (c *projectGraphController) AddClusterScopedHandler(name, cluster string, handler ProjectGraphHandlerFunc) {
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

		return handler(key, obj.(*ProjectGraph))
	})
}

type projectGraphFactory struct {
}

func (c projectGraphFactory) Object() runtime.Object {
	return &ProjectGraph{}
}

func (c projectGraphFactory) List() runtime.Object {
	return &ProjectGraphList{}
}

func (s *projectGraphClient) Controller() ProjectGraphController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.projectGraphControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ProjectGraphGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &projectGraphController{
		GenericController: genericController,
	}

	s.client.projectGraphControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type projectGraphClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ProjectGraphController
}

func (s *projectGraphClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *projectGraphClient) Create(o *ProjectGraph) (*ProjectGraph, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ProjectGraph), err
}

func (s *projectGraphClient) Get(name string, opts metav1.GetOptions) (*ProjectGraph, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ProjectGraph), err
}

func (s *projectGraphClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ProjectGraph, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*ProjectGraph), err
}

func (s *projectGraphClient) Update(o *ProjectGraph) (*ProjectGraph, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ProjectGraph), err
}

func (s *projectGraphClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *projectGraphClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *projectGraphClient) List(opts metav1.ListOptions) (*ProjectGraphList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ProjectGraphList), err
}

func (s *projectGraphClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *projectGraphClient) Patch(o *ProjectGraph, data []byte, subresources ...string) (*ProjectGraph, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*ProjectGraph), err
}

func (s *projectGraphClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *projectGraphClient) AddHandler(name string, sync ProjectGraphHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *projectGraphClient) AddLifecycle(name string, lifecycle ProjectGraphLifecycle) {
	sync := NewProjectGraphLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *projectGraphClient) AddClusterScopedHandler(name, clusterName string, sync ProjectGraphHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *projectGraphClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle ProjectGraphLifecycle) {
	sync := NewProjectGraphLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
