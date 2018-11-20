package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ProjectAlertPolicyLifecycle interface {
	Create(obj *ProjectAlertPolicy) (runtime.Object, error)
	Remove(obj *ProjectAlertPolicy) (runtime.Object, error)
	Updated(obj *ProjectAlertPolicy) (runtime.Object, error)
}

type projectAlertPolicyLifecycleAdapter struct {
	lifecycle ProjectAlertPolicyLifecycle
}

func (w *projectAlertPolicyLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ProjectAlertPolicy))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *projectAlertPolicyLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ProjectAlertPolicy))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *projectAlertPolicyLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ProjectAlertPolicy))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewProjectAlertPolicyLifecycleAdapter(name string, clusterScoped bool, client ProjectAlertPolicyInterface, l ProjectAlertPolicyLifecycle) ProjectAlertPolicyHandlerFunc {
	adapter := &projectAlertPolicyLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ProjectAlertPolicy) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
