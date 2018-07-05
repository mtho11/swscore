package kubernetes

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// QuotaSpecBinding is one of the specific types defined by Istio as a Kubernetes extension.
// Preliminar rule is to define one file per type.
// types.go will collect common/shared types.
// This type is extracted from Istio Pilot as models used are not public and it doesn't make sense to fetch all
// Istio project as a dependency.
// Reference: https://github.com/istio/istio/blob/master/pilot/pkg/config/kube/crd/types.go

const quotaspecbindings = "quotaspecbindings"
const quotaspecbindingType = "QuotaSpecBinding"
const quotaspecbindingLabel = "quotaspecbinding"

// QuotaSpecBinding is the generic Kubernetes API object wrapper
// QuotaSpecBinding starts with uppercase as it maps a "kind":"QuotaSpecBinding" Istio response.
type QuotaSpecBinding struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata"`
	Spec               map[string]interface{} `json:"spec"`
}

// QuotaSpecBindingList is the generic Kubernetes API list wrapper
// QuotaSpecBindingList starts with uppercase as it maps a "kind":"QuotaSpecBindingList" Istio response.
type QuotaSpecBindingList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []QuotaSpecBinding `json:"items"`
}

// GetSpec from a wrapper
func (in *QuotaSpecBinding) GetSpec() map[string]interface{} {
	return in.Spec
}

// SetSpec for a wrapper
func (in *QuotaSpecBinding) SetSpec(spec map[string]interface{}) {
	in.Spec = spec
}

// GetObjectMeta from a wrapper
func (in *QuotaSpecBinding) GetObjectMeta() meta_v1.ObjectMeta {
	return in.ObjectMeta
}

// SetObjectMeta for a wrapper
func (in *QuotaSpecBinding) SetObjectMeta(metadata meta_v1.ObjectMeta) {
	in.ObjectMeta = metadata
}

// GetItems from a wrapper
func (in *QuotaSpecBindingList) GetItems() []IstioObject {
	out := make([]IstioObject, len(in.Items))
	for i := range in.Items {
		out[i] = &in.Items[i]
	}
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaSpecBinding) DeepCopyInto(out *QuotaSpecBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaSpecBinding.
func (in *QuotaSpecBinding) DeepCopy() *QuotaSpecBinding {
	if in == nil {
		return nil
	}
	out := new(QuotaSpecBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuotaSpecBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}

// DeepCopyIstioObject is an autogenerated deepcopy function, copying the receiver, creating a new IstioObject.
func (in *QuotaSpecBinding) DeepCopyIstioObject() IstioObject {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaSpecBindingList) DeepCopyInto(out *QuotaSpecBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QuotaSpecBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleList.
func (in *QuotaSpecBindingList) DeepCopy() *QuotaSpecBindingList {
	if in == nil {
		return nil
	}
	out := new(QuotaSpecBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuotaSpecBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}
