package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ClusterGraphLifecycle interface {
	Create(obj *ClusterGraph) (*ClusterGraph, error)
	Remove(obj *ClusterGraph) (*ClusterGraph, error)
	Updated(obj *ClusterGraph) (*ClusterGraph, error)
}

type clusterGraphLifecycleAdapter struct {
	lifecycle ClusterGraphLifecycle
}

func (w *clusterGraphLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ClusterGraph))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterGraphLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ClusterGraph))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterGraphLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ClusterGraph))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewClusterGraphLifecycleAdapter(name string, clusterScoped bool, client ClusterGraphInterface, l ClusterGraphLifecycle) ClusterGraphHandlerFunc {
	adapter := &clusterGraphLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ClusterGraph) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
