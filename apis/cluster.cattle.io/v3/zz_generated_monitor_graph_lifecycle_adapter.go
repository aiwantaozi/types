package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type MonitorGraphLifecycle interface {
	Create(obj *MonitorGraph) (*MonitorGraph, error)
	Remove(obj *MonitorGraph) (*MonitorGraph, error)
	Updated(obj *MonitorGraph) (*MonitorGraph, error)
}

type monitorGraphLifecycleAdapter struct {
	lifecycle MonitorGraphLifecycle
}

func (w *monitorGraphLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*MonitorGraph))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *monitorGraphLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*MonitorGraph))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *monitorGraphLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*MonitorGraph))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewMonitorGraphLifecycleAdapter(name string, clusterScoped bool, client MonitorGraphInterface, l MonitorGraphLifecycle) MonitorGraphHandlerFunc {
	adapter := &monitorGraphLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *MonitorGraph) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
