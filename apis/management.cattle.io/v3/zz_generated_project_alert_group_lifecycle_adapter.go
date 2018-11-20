package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ProjectAlertGroupLifecycle interface {
	Create(obj *ProjectAlertGroup) (*ProjectAlertGroup, error)
	Remove(obj *ProjectAlertGroup) (*ProjectAlertGroup, error)
	Updated(obj *ProjectAlertGroup) (*ProjectAlertGroup, error)
}

type projectAlertGroupLifecycleAdapter struct {
	lifecycle ProjectAlertGroupLifecycle
}

func (w *projectAlertGroupLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ProjectAlertGroup))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *projectAlertGroupLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ProjectAlertGroup))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *projectAlertGroupLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ProjectAlertGroup))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewProjectAlertGroupLifecycleAdapter(name string, clusterScoped bool, client ProjectAlertGroupInterface, l ProjectAlertGroupLifecycle) ProjectAlertGroupHandlerFunc {
	adapter := &projectAlertGroupLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ProjectAlertGroup) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
