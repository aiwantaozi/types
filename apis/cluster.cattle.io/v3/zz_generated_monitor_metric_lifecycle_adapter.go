package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type MonitorMetricLifecycle interface {
	Create(obj *MonitorMetric) (*MonitorMetric, error)
	Remove(obj *MonitorMetric) (*MonitorMetric, error)
	Updated(obj *MonitorMetric) (*MonitorMetric, error)
}

type monitorMetricLifecycleAdapter struct {
	lifecycle MonitorMetricLifecycle
}

func (w *monitorMetricLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*MonitorMetric))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *monitorMetricLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*MonitorMetric))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *monitorMetricLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*MonitorMetric))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewMonitorMetricLifecycleAdapter(name string, clusterScoped bool, client MonitorMetricInterface, l MonitorMetricLifecycle) MonitorMetricHandlerFunc {
	adapter := &monitorMetricLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *MonitorMetric) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
