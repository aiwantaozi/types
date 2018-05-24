package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ClusterAuditLoggingLifecycle interface {
	Create(obj *ClusterAuditLogging) (*ClusterAuditLogging, error)
	Remove(obj *ClusterAuditLogging) (*ClusterAuditLogging, error)
	Updated(obj *ClusterAuditLogging) (*ClusterAuditLogging, error)
}

type clusterAuditLoggingLifecycleAdapter struct {
	lifecycle ClusterAuditLoggingLifecycle
}

func (w *clusterAuditLoggingLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ClusterAuditLogging))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterAuditLoggingLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ClusterAuditLogging))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterAuditLoggingLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ClusterAuditLogging))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewClusterAuditLoggingLifecycleAdapter(name string, clusterScoped bool, client ClusterAuditLoggingInterface, l ClusterAuditLoggingLifecycle) ClusterAuditLoggingHandlerFunc {
	adapter := &clusterAuditLoggingLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ClusterAuditLogging) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
