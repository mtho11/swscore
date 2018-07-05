package kubernetes

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// solarwinds is one of the specific types defined by Istio as a Kubernetes extension.
// Preliminar rule is to define one file per type.
// types.go will collect common/shared types.
// This type is extracted from Istio Pilot as models used are not public and it doesn't make sense to fetch all
// Istio project as a dependency.
// Reference: https://github.com/istio/istio/blob/master/pilot/pkg/config/kube/crd/types.go

const solarwindses = "solarwindses"
const solarwindsType = "solarwinds"
const solarwindsLabel = "solarwinds"

// solarwinds is the generic Kubernetes API object wrapper
// solarwinds starts with lowercase as it maps a "kind":"solarwinds" Istio response.
type solarwinds struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata"`
	Spec               map[string]interface{} `json:"spec"`
}

// solarwindsList is the generic Kubernetes API list wrapper
// solarwindsList starts with lowercase as it maps a "kind":"solarwindsList" Istio response.
type solarwindsList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []solarwinds `json:"items"`
}

// GetSpec from a wrapper
func (in *solarwinds) GetSpec() map[string]interface{} {
	return in.Spec
}

// SetSpec for a wrapper
func (in *solarwinds) SetSpec(spec map[string]interface{}) {
	in.Spec = spec
}

// GetObjectMeta from a wrapper
func (in *solarwinds) GetObjectMeta() meta_v1.ObjectMeta {
	return in.ObjectMeta
}

// SetObjectMeta for a wrapper
func (in *solarwinds) SetObjectMeta(metadata meta_v1.ObjectMeta) {
	in.ObjectMeta = metadata
}

// GetItems from a wrapper
func (in *solarwindsList) GetItems() []IstioObject {
	out := make([]IstioObject, len(in.Items))
	for i := range in.Items {
		out[i] = &in.Items[i]
	}
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *solarwinds) DeepCopyInto(out *solarwinds) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new solarwinds.
func (in *solarwinds) DeepCopy() *solarwinds {
	if in == nil {
		return nil
	}
	out := new(solarwinds)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *solarwinds) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}

// DeepCopyIstioObject is an autogenerated deepcopy function, copying the receiver, creating a new IstioObject.
func (in *solarwinds) DeepCopyIstioObject() IstioObject {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *solarwindsList) DeepCopyInto(out *solarwindsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]solarwinds, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleList.
func (in *solarwindsList) DeepCopy() *solarwindsList {
	if in == nil {
		return nil
	}
	out := new(solarwindsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *solarwindsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}
