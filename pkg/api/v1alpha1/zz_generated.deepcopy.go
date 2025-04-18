//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in EgressNetworkPolicies) DeepCopyInto(out *EgressNetworkPolicies) {
	{
		in := &in
		*out = make(EgressNetworkPolicies, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressNetworkPolicies.
func (in EgressNetworkPolicies) DeepCopy() EgressNetworkPolicies {
	if in == nil {
		return nil
	}
	out := new(EgressNetworkPolicies)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressNetworkPolicy) DeepCopyInto(out *EgressNetworkPolicy) {
	*out = *in
	out.To = in.To
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressNetworkPolicy.
func (in *EgressNetworkPolicy) DeepCopy() *EgressNetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(EgressNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressNetworkPolicyTo) DeepCopyInto(out *EgressNetworkPolicyTo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressNetworkPolicyTo.
func (in *EgressNetworkPolicyTo) DeepCopy() *EgressNetworkPolicyTo {
	if in == nil {
		return nil
	}
	out := new(EgressNetworkPolicyTo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateKey) DeepCopyInto(out *PrivateKey) {
	*out = *in
	in.SecretKeyRef.DeepCopyInto(&out.SecretKeyRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateKey.
func (in *PrivateKey) DeepCopy() *PrivateKey {
	if in == nil {
		return nil
	}
	out := new(PrivateKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Speed) DeepCopyInto(out *Speed) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Speed.
func (in *Speed) DeepCopy() *Speed {
	if in == nil {
		return nil
	}
	out := new(Speed)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Status) DeepCopyInto(out *Status) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WgStatusReport) DeepCopyInto(out *WgStatusReport) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WgStatusReport.
func (in *WgStatusReport) DeepCopy() *WgStatusReport {
	if in == nil {
		return nil
	}
	out := new(WgStatusReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Wireguard) DeepCopyInto(out *Wireguard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Wireguard.
func (in *Wireguard) DeepCopy() *Wireguard {
	if in == nil {
		return nil
	}
	out := new(Wireguard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Wireguard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WireguardList) DeepCopyInto(out *WireguardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Wireguard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WireguardList.
func (in *WireguardList) DeepCopy() *WireguardList {
	if in == nil {
		return nil
	}
	out := new(WireguardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WireguardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WireguardPeer) DeepCopyInto(out *WireguardPeer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WireguardPeer.
func (in *WireguardPeer) DeepCopy() *WireguardPeer {
	if in == nil {
		return nil
	}
	out := new(WireguardPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WireguardPeer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WireguardPeerList) DeepCopyInto(out *WireguardPeerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WireguardPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WireguardPeerList.
func (in *WireguardPeerList) DeepCopy() *WireguardPeerList {
	if in == nil {
		return nil
	}
	out := new(WireguardPeerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WireguardPeerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WireguardPeerSpec) DeepCopyInto(out *WireguardPeerSpec) {
	*out = *in
	in.PrivateKey.DeepCopyInto(&out.PrivateKey)
	if in.EgressNetworkPolicies != nil {
		in, out := &in.EgressNetworkPolicies, &out.EgressNetworkPolicies
		*out = make(EgressNetworkPolicies, len(*in))
		copy(*out, *in)
	}
	out.DownloadSpeed = in.DownloadSpeed
	out.UploadSpeed = in.UploadSpeed
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WireguardPeerSpec.
func (in *WireguardPeerSpec) DeepCopy() *WireguardPeerSpec {
	if in == nil {
		return nil
	}
	out := new(WireguardPeerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WireguardPeerStatus) DeepCopyInto(out *WireguardPeerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WireguardPeerStatus.
func (in *WireguardPeerStatus) DeepCopy() *WireguardPeerStatus {
	if in == nil {
		return nil
	}
	out := new(WireguardPeerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WireguardSpec) DeepCopyInto(out *WireguardSpec) {
	*out = *in
	if in.ServiceAnnotations != nil {
		in, out := &in.ServiceAnnotations, &out.ServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeploymentAffinity != nil {
		in, out := &in.DeploymentAffinity, &out.DeploymentAffinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.DeploymentNodeSelector != nil {
		in, out := &in.DeploymentNodeSelector, &out.DeploymentNodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeploymentTolerations != nil {
		in, out := &in.DeploymentTolerations, &out.DeploymentTolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WireguardSpec.
func (in *WireguardSpec) DeepCopy() *WireguardSpec {
	if in == nil {
		return nil
	}
	out := new(WireguardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WireguardStatus) DeepCopyInto(out *WireguardStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WireguardStatus.
func (in *WireguardStatus) DeepCopy() *WireguardStatus {
	if in == nil {
		return nil
	}
	out := new(WireguardStatus)
	in.DeepCopyInto(out)
	return out
}
