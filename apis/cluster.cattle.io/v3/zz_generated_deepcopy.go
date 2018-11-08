package v3

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGraph) DeepCopyInto(out *ClusterGraph) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGraph.
func (in *ClusterGraph) DeepCopy() *ClusterGraph {
	if in == nil {
		return nil
	}
	out := new(ClusterGraph)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterGraph) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGraphList) DeepCopyInto(out *ClusterGraphList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterGraph, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGraphList.
func (in *ClusterGraphList) DeepCopy() *ClusterGraphList {
	if in == nil {
		return nil
	}
	out := new(ClusterGraphList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterGraphList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGraphSpec) DeepCopyInto(out *ClusterGraphSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGraphSpec.
func (in *ClusterGraphSpec) DeepCopy() *ClusterGraphSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterGraphSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorGraph) DeepCopyInto(out *MonitorGraph) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorGraph.
func (in *MonitorGraph) DeepCopy() *MonitorGraph {
	if in == nil {
		return nil
	}
	out := new(MonitorGraph)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitorGraph) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorGraphList) DeepCopyInto(out *MonitorGraphList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonitorGraph, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorGraphList.
func (in *MonitorGraphList) DeepCopy() *MonitorGraphList {
	if in == nil {
		return nil
	}
	out := new(MonitorGraphList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitorGraphList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorGraphSpec) DeepCopyInto(out *MonitorGraphSpec) {
	*out = *in
	if in.MetricsSelector != nil {
		in, out := &in.MetricsSelector, &out.MetricsSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.XAxis = in.XAxis
	out.YAxis = in.YAxis
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorGraphSpec.
func (in *MonitorGraphSpec) DeepCopy() *MonitorGraphSpec {
	if in == nil {
		return nil
	}
	out := new(MonitorGraphSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorGraphStatus) DeepCopyInto(out *MonitorGraphStatus) {
	*out = *in
	if in.MonitorMetricIDs != nil {
		in, out := &in.MonitorMetricIDs, &out.MonitorMetricIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorGraphStatus.
func (in *MonitorGraphStatus) DeepCopy() *MonitorGraphStatus {
	if in == nil {
		return nil
	}
	out := new(MonitorGraphStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorMetric) DeepCopyInto(out *MonitorMetric) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorMetric.
func (in *MonitorMetric) DeepCopy() *MonitorMetric {
	if in == nil {
		return nil
	}
	out := new(MonitorMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitorMetric) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorMetricList) DeepCopyInto(out *MonitorMetricList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonitorMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorMetricList.
func (in *MonitorMetricList) DeepCopy() *MonitorMetricList {
	if in == nil {
		return nil
	}
	out := new(MonitorMetricList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitorMetricList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorMetricSpec) DeepCopyInto(out *MonitorMetricSpec) {
	*out = *in
	if in.ExtraAddedTags != nil {
		in, out := &in.ExtraAddedTags, &out.ExtraAddedTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorMetricSpec.
func (in *MonitorMetricSpec) DeepCopy() *MonitorMetricSpec {
	if in == nil {
		return nil
	}
	out := new(MonitorMetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectGraph) DeepCopyInto(out *ProjectGraph) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectGraph.
func (in *ProjectGraph) DeepCopy() *ProjectGraph {
	if in == nil {
		return nil
	}
	out := new(ProjectGraph)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectGraph) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectGraphList) DeepCopyInto(out *ProjectGraphList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProjectGraph, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectGraphList.
func (in *ProjectGraphList) DeepCopy() *ProjectGraphList {
	if in == nil {
		return nil
	}
	out := new(ProjectGraphList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectGraphList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectGraphSpec) DeepCopyInto(out *ProjectGraphSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectGraphSpec.
func (in *ProjectGraphSpec) DeepCopy() *ProjectGraphSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectGraphSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatsStatus) DeepCopyInto(out *StatsStatus) {
	*out = *in
	if in.MonitorGraphIDs != nil {
		in, out := &in.MonitorGraphIDs, &out.MonitorGraphIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatsStatus.
func (in *StatsStatus) DeepCopy() *StatsStatus {
	if in == nil {
		return nil
	}
	out := new(StatsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XAxis) DeepCopyInto(out *XAxis) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XAxis.
func (in *XAxis) DeepCopy() *XAxis {
	if in == nil {
		return nil
	}
	out := new(XAxis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YAxis) DeepCopyInto(out *YAxis) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YAxis.
func (in *YAxis) DeepCopy() *YAxis {
	if in == nil {
		return nil
	}
	out := new(YAxis)
	in.DeepCopyInto(out)
	return out
}
