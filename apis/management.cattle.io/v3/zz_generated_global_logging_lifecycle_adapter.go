package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type GlobalLoggingLifecycle interface {
	Create(obj *GlobalLogging) (*GlobalLogging, error)
	Remove(obj *GlobalLogging) (*GlobalLogging, error)
	Updated(obj *GlobalLogging) (*GlobalLogging, error)
}

type globalLoggingLifecycleAdapter struct {
	lifecycle GlobalLoggingLifecycle
}

func (w *globalLoggingLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*GlobalLogging))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *globalLoggingLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*GlobalLogging))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *globalLoggingLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*GlobalLogging))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewGlobalLoggingLifecycleAdapter(name string, clusterScoped bool, client GlobalLoggingInterface, l GlobalLoggingLifecycle) GlobalLoggingHandlerFunc {
	adapter := &globalLoggingLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *GlobalLogging) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
