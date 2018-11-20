package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ClusterAlertRuleLifecycle interface {
	Create(obj *ClusterAlertRule) (*ClusterAlertRule, error)
	Remove(obj *ClusterAlertRule) (*ClusterAlertRule, error)
	Updated(obj *ClusterAlertRule) (*ClusterAlertRule, error)
}

type clusterAlertRuleLifecycleAdapter struct {
	lifecycle ClusterAlertRuleLifecycle
}

func (w *clusterAlertRuleLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ClusterAlertRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterAlertRuleLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ClusterAlertRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterAlertRuleLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ClusterAlertRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewClusterAlertRuleLifecycleAdapter(name string, clusterScoped bool, client ClusterAlertRuleInterface, l ClusterAlertRuleLifecycle) ClusterAlertRuleHandlerFunc {
	adapter := &clusterAlertRuleLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ClusterAlertRule) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
