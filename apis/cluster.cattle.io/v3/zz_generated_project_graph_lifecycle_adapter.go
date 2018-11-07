package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ProjectGraphLifecycle interface {
	Create(obj *ProjectGraph) (*ProjectGraph, error)
	Remove(obj *ProjectGraph) (*ProjectGraph, error)
	Updated(obj *ProjectGraph) (*ProjectGraph, error)
}

type projectGraphLifecycleAdapter struct {
	lifecycle ProjectGraphLifecycle
}

func (w *projectGraphLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ProjectGraph))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *projectGraphLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ProjectGraph))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *projectGraphLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ProjectGraph))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewProjectGraphLifecycleAdapter(name string, clusterScoped bool, client ProjectGraphInterface, l ProjectGraphLifecycle) ProjectGraphHandlerFunc {
	adapter := &projectGraphLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ProjectGraph) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
