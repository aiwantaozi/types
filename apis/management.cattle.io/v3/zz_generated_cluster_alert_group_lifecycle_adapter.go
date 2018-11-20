package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ClusterAlertGroupLifecycle interface {
	Create(obj *ClusterAlertGroup) (*ClusterAlertGroup, error)
	Remove(obj *ClusterAlertGroup) (*ClusterAlertGroup, error)
	Updated(obj *ClusterAlertGroup) (*ClusterAlertGroup, error)
}

type clusterAlertGroupLifecycleAdapter struct {
	lifecycle ClusterAlertGroupLifecycle
}

func (w *clusterAlertGroupLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ClusterAlertGroup))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterAlertGroupLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ClusterAlertGroup))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterAlertGroupLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ClusterAlertGroup))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewClusterAlertGroupLifecycleAdapter(name string, clusterScoped bool, client ClusterAlertGroupInterface, l ClusterAlertGroupLifecycle) ClusterAlertGroupHandlerFunc {
	adapter := &clusterAlertGroupLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ClusterAlertGroup) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
