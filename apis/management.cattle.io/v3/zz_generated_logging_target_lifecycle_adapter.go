package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type LoggingTargetLifecycle interface {
	Create(obj *LoggingTarget) (*LoggingTarget, error)
	Remove(obj *LoggingTarget) (*LoggingTarget, error)
	Updated(obj *LoggingTarget) (*LoggingTarget, error)
}

type loggingTargetLifecycleAdapter struct {
	lifecycle LoggingTargetLifecycle
}

func (w *loggingTargetLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*LoggingTarget))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *loggingTargetLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*LoggingTarget))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *loggingTargetLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*LoggingTarget))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewLoggingTargetLifecycleAdapter(name string, clusterScoped bool, client LoggingTargetInterface, l LoggingTargetLifecycle) LoggingTargetHandlerFunc {
	adapter := &loggingTargetLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *LoggingTarget) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
