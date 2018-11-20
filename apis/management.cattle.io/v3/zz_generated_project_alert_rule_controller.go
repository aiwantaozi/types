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
	ProjectAlertRuleGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ProjectAlertRule",
	}
	ProjectAlertRuleResource = metav1.APIResource{
		Name:         "projectalertrules",
		SingularName: "projectalertrule",
		Namespaced:   true,

		Kind: ProjectAlertRuleGroupVersionKind.Kind,
	}
)

type ProjectAlertRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectAlertRule
}

type ProjectAlertRuleHandlerFunc func(key string, obj *ProjectAlertRule) error

type ProjectAlertRuleLister interface {
	List(namespace string, selector labels.Selector) (ret []*ProjectAlertRule, err error)
	Get(namespace, name string) (*ProjectAlertRule, error)
}

type ProjectAlertRuleController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() ProjectAlertRuleLister
	AddHandler(name string, handler ProjectAlertRuleHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler ProjectAlertRuleHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ProjectAlertRuleInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*ProjectAlertRule) (*ProjectAlertRule, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ProjectAlertRule, error)
	Get(name string, opts metav1.GetOptions) (*ProjectAlertRule, error)
	Update(*ProjectAlertRule) (*ProjectAlertRule, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ProjectAlertRuleList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ProjectAlertRuleController
	AddHandler(name string, sync ProjectAlertRuleHandlerFunc)
	AddLifecycle(name string, lifecycle ProjectAlertRuleLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync ProjectAlertRuleHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle ProjectAlertRuleLifecycle)
}

type projectAlertRuleLister struct {
	controller *projectAlertRuleController
}

func (l *projectAlertRuleLister) List(namespace string, selector labels.Selector) (ret []*ProjectAlertRule, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*ProjectAlertRule))
	})
	return
}

func (l *projectAlertRuleLister) Get(namespace, name string) (*ProjectAlertRule, error) {
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
			Group:    ProjectAlertRuleGroupVersionKind.Group,
			Resource: "projectAlertRule",
		}, key)
	}
	return obj.(*ProjectAlertRule), nil
}

type projectAlertRuleController struct {
	controller.GenericController
}

func (c *projectAlertRuleController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *projectAlertRuleController) Lister() ProjectAlertRuleLister {
	return &projectAlertRuleLister{
		controller: c,
	}
}

func (c *projectAlertRuleController) AddHandler(name string, handler ProjectAlertRuleHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*ProjectAlertRule))
	})
}

func (c *projectAlertRuleController) AddClusterScopedHandler(name, cluster string, handler ProjectAlertRuleHandlerFunc) {
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

		return handler(key, obj.(*ProjectAlertRule))
	})
}

type projectAlertRuleFactory struct {
}

func (c projectAlertRuleFactory) Object() runtime.Object {
	return &ProjectAlertRule{}
}

func (c projectAlertRuleFactory) List() runtime.Object {
	return &ProjectAlertRuleList{}
}

func (s *projectAlertRuleClient) Controller() ProjectAlertRuleController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.projectAlertRuleControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ProjectAlertRuleGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &projectAlertRuleController{
		GenericController: genericController,
	}

	s.client.projectAlertRuleControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type projectAlertRuleClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ProjectAlertRuleController
}

func (s *projectAlertRuleClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *projectAlertRuleClient) Create(o *ProjectAlertRule) (*ProjectAlertRule, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ProjectAlertRule), err
}

func (s *projectAlertRuleClient) Get(name string, opts metav1.GetOptions) (*ProjectAlertRule, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ProjectAlertRule), err
}

func (s *projectAlertRuleClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ProjectAlertRule, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*ProjectAlertRule), err
}

func (s *projectAlertRuleClient) Update(o *ProjectAlertRule) (*ProjectAlertRule, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ProjectAlertRule), err
}

func (s *projectAlertRuleClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *projectAlertRuleClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *projectAlertRuleClient) List(opts metav1.ListOptions) (*ProjectAlertRuleList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ProjectAlertRuleList), err
}

func (s *projectAlertRuleClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *projectAlertRuleClient) Patch(o *ProjectAlertRule, data []byte, subresources ...string) (*ProjectAlertRule, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*ProjectAlertRule), err
}

func (s *projectAlertRuleClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *projectAlertRuleClient) AddHandler(name string, sync ProjectAlertRuleHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *projectAlertRuleClient) AddLifecycle(name string, lifecycle ProjectAlertRuleLifecycle) {
	sync := NewProjectAlertRuleLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *projectAlertRuleClient) AddClusterScopedHandler(name, clusterName string, sync ProjectAlertRuleHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *projectAlertRuleClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle ProjectAlertRuleLifecycle) {
	sync := NewProjectAlertRuleLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
