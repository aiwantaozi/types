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
	ClusterAlertGroupGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ClusterAlertGroup",
	}
	ClusterAlertGroupResource = metav1.APIResource{
		Name:         "clusteralertgroups",
		SingularName: "clusteralertgroup",
		Namespaced:   true,

		Kind: ClusterAlertGroupGroupVersionKind.Kind,
	}
)

type ClusterAlertGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterAlertGroup
}

type ClusterAlertGroupHandlerFunc func(key string, obj *ClusterAlertGroup) error

type ClusterAlertGroupLister interface {
	List(namespace string, selector labels.Selector) (ret []*ClusterAlertGroup, err error)
	Get(namespace, name string) (*ClusterAlertGroup, error)
}

type ClusterAlertGroupController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() ClusterAlertGroupLister
	AddHandler(name string, handler ClusterAlertGroupHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler ClusterAlertGroupHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ClusterAlertGroupInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*ClusterAlertGroup) (*ClusterAlertGroup, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ClusterAlertGroup, error)
	Get(name string, opts metav1.GetOptions) (*ClusterAlertGroup, error)
	Update(*ClusterAlertGroup) (*ClusterAlertGroup, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ClusterAlertGroupList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ClusterAlertGroupController
	AddHandler(name string, sync ClusterAlertGroupHandlerFunc)
	AddLifecycle(name string, lifecycle ClusterAlertGroupLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync ClusterAlertGroupHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle ClusterAlertGroupLifecycle)
}

type clusterAlertGroupLister struct {
	controller *clusterAlertGroupController
}

func (l *clusterAlertGroupLister) List(namespace string, selector labels.Selector) (ret []*ClusterAlertGroup, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*ClusterAlertGroup))
	})
	return
}

func (l *clusterAlertGroupLister) Get(namespace, name string) (*ClusterAlertGroup, error) {
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
			Group:    ClusterAlertGroupGroupVersionKind.Group,
			Resource: "clusterAlertGroup",
		}, key)
	}
	return obj.(*ClusterAlertGroup), nil
}

type clusterAlertGroupController struct {
	controller.GenericController
}

func (c *clusterAlertGroupController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *clusterAlertGroupController) Lister() ClusterAlertGroupLister {
	return &clusterAlertGroupLister{
		controller: c,
	}
}

func (c *clusterAlertGroupController) AddHandler(name string, handler ClusterAlertGroupHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*ClusterAlertGroup))
	})
}

func (c *clusterAlertGroupController) AddClusterScopedHandler(name, cluster string, handler ClusterAlertGroupHandlerFunc) {
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

		return handler(key, obj.(*ClusterAlertGroup))
	})
}

type clusterAlertGroupFactory struct {
}

func (c clusterAlertGroupFactory) Object() runtime.Object {
	return &ClusterAlertGroup{}
}

func (c clusterAlertGroupFactory) List() runtime.Object {
	return &ClusterAlertGroupList{}
}

func (s *clusterAlertGroupClient) Controller() ClusterAlertGroupController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.clusterAlertGroupControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ClusterAlertGroupGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &clusterAlertGroupController{
		GenericController: genericController,
	}

	s.client.clusterAlertGroupControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type clusterAlertGroupClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ClusterAlertGroupController
}

func (s *clusterAlertGroupClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *clusterAlertGroupClient) Create(o *ClusterAlertGroup) (*ClusterAlertGroup, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ClusterAlertGroup), err
}

func (s *clusterAlertGroupClient) Get(name string, opts metav1.GetOptions) (*ClusterAlertGroup, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ClusterAlertGroup), err
}

func (s *clusterAlertGroupClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ClusterAlertGroup, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*ClusterAlertGroup), err
}

func (s *clusterAlertGroupClient) Update(o *ClusterAlertGroup) (*ClusterAlertGroup, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ClusterAlertGroup), err
}

func (s *clusterAlertGroupClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *clusterAlertGroupClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *clusterAlertGroupClient) List(opts metav1.ListOptions) (*ClusterAlertGroupList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ClusterAlertGroupList), err
}

func (s *clusterAlertGroupClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *clusterAlertGroupClient) Patch(o *ClusterAlertGroup, data []byte, subresources ...string) (*ClusterAlertGroup, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*ClusterAlertGroup), err
}

func (s *clusterAlertGroupClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *clusterAlertGroupClient) AddHandler(name string, sync ClusterAlertGroupHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *clusterAlertGroupClient) AddLifecycle(name string, lifecycle ClusterAlertGroupLifecycle) {
	sync := NewClusterAlertGroupLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *clusterAlertGroupClient) AddClusterScopedHandler(name, clusterName string, sync ClusterAlertGroupHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *clusterAlertGroupClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle ClusterAlertGroupLifecycle) {
	sync := NewClusterAlertGroupLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
