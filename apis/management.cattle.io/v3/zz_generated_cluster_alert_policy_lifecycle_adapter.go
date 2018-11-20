package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ClusterAlertPolicyLifecycle interface {
	Create(obj *ClusterAlertPolicy) (runtime.Object, error)
	Remove(obj *ClusterAlertPolicy) (runtime.Object, error)
	Updated(obj *ClusterAlertPolicy) (runtime.Object, error)
}

type clusterAlertPolicyLifecycleAdapter struct {
	lifecycle ClusterAlertPolicyLifecycle
}

func (w *clusterAlertPolicyLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ClusterAlertPolicy))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterAlertPolicyLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ClusterAlertPolicy))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterAlertPolicyLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ClusterAlertPolicy))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewClusterAlertPolicyLifecycleAdapter(name string, clusterScoped bool, client ClusterAlertPolicyInterface, l ClusterAlertPolicyLifecycle) ClusterAlertPolicyHandlerFunc {
	adapter := &clusterAlertPolicyLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ClusterAlertPolicy) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
