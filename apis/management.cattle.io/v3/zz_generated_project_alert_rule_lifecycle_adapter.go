package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ProjectAlertRuleLifecycle interface {
	Create(obj *ProjectAlertRule) (*ProjectAlertRule, error)
	Remove(obj *ProjectAlertRule) (*ProjectAlertRule, error)
	Updated(obj *ProjectAlertRule) (*ProjectAlertRule, error)
}

type projectAlertRuleLifecycleAdapter struct {
	lifecycle ProjectAlertRuleLifecycle
}

func (w *projectAlertRuleLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ProjectAlertRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *projectAlertRuleLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ProjectAlertRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *projectAlertRuleLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ProjectAlertRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewProjectAlertRuleLifecycleAdapter(name string, clusterScoped bool, client ProjectAlertRuleInterface, l ProjectAlertRuleLifecycle) ProjectAlertRuleHandlerFunc {
	adapter := &projectAlertRuleLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ProjectAlertRule) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
