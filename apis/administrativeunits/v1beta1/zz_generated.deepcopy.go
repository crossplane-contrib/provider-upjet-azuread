//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Member) DeepCopyInto(out *Member) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Member.
func (in *Member) DeepCopy() *Member {
	if in == nil {
		return nil
	}
	out := new(Member)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Member) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberInitParameters) DeepCopyInto(out *MemberInitParameters) {
	*out = *in
	if in.AdministrativeUnitObjectID != nil {
		in, out := &in.AdministrativeUnitObjectID, &out.AdministrativeUnitObjectID
		*out = new(string)
		**out = **in
	}
	if in.AdministrativeUnitObjectIDRef != nil {
		in, out := &in.AdministrativeUnitObjectIDRef, &out.AdministrativeUnitObjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AdministrativeUnitObjectIDSelector != nil {
		in, out := &in.AdministrativeUnitObjectIDSelector, &out.AdministrativeUnitObjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.MemberObjectID != nil {
		in, out := &in.MemberObjectID, &out.MemberObjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberInitParameters.
func (in *MemberInitParameters) DeepCopy() *MemberInitParameters {
	if in == nil {
		return nil
	}
	out := new(MemberInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberList) DeepCopyInto(out *MemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Member, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberList.
func (in *MemberList) DeepCopy() *MemberList {
	if in == nil {
		return nil
	}
	out := new(MemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberObservation) DeepCopyInto(out *MemberObservation) {
	*out = *in
	if in.AdministrativeUnitObjectID != nil {
		in, out := &in.AdministrativeUnitObjectID, &out.AdministrativeUnitObjectID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MemberObjectID != nil {
		in, out := &in.MemberObjectID, &out.MemberObjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberObservation.
func (in *MemberObservation) DeepCopy() *MemberObservation {
	if in == nil {
		return nil
	}
	out := new(MemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberParameters) DeepCopyInto(out *MemberParameters) {
	*out = *in
	if in.AdministrativeUnitObjectID != nil {
		in, out := &in.AdministrativeUnitObjectID, &out.AdministrativeUnitObjectID
		*out = new(string)
		**out = **in
	}
	if in.AdministrativeUnitObjectIDRef != nil {
		in, out := &in.AdministrativeUnitObjectIDRef, &out.AdministrativeUnitObjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AdministrativeUnitObjectIDSelector != nil {
		in, out := &in.AdministrativeUnitObjectIDSelector, &out.AdministrativeUnitObjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.MemberObjectID != nil {
		in, out := &in.MemberObjectID, &out.MemberObjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberParameters.
func (in *MemberParameters) DeepCopy() *MemberParameters {
	if in == nil {
		return nil
	}
	out := new(MemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberSpec) DeepCopyInto(out *MemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberSpec.
func (in *MemberSpec) DeepCopy() *MemberSpec {
	if in == nil {
		return nil
	}
	out := new(MemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberStatus) DeepCopyInto(out *MemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberStatus.
func (in *MemberStatus) DeepCopy() *MemberStatus {
	if in == nil {
		return nil
	}
	out := new(MemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Unit) DeepCopyInto(out *Unit) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Unit.
func (in *Unit) DeepCopy() *Unit {
	if in == nil {
		return nil
	}
	out := new(Unit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Unit) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnitInitParameters) DeepCopyInto(out *UnitInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.HiddenMembershipEnabled != nil {
		in, out := &in.HiddenMembershipEnabled, &out.HiddenMembershipEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PreventDuplicateNames != nil {
		in, out := &in.PreventDuplicateNames, &out.PreventDuplicateNames
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnitInitParameters.
func (in *UnitInitParameters) DeepCopy() *UnitInitParameters {
	if in == nil {
		return nil
	}
	out := new(UnitInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnitList) DeepCopyInto(out *UnitList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Unit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnitList.
func (in *UnitList) DeepCopy() *UnitList {
	if in == nil {
		return nil
	}
	out := new(UnitList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UnitList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnitObservation) DeepCopyInto(out *UnitObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.HiddenMembershipEnabled != nil {
		in, out := &in.HiddenMembershipEnabled, &out.HiddenMembershipEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ObjectID != nil {
		in, out := &in.ObjectID, &out.ObjectID
		*out = new(string)
		**out = **in
	}
	if in.PreventDuplicateNames != nil {
		in, out := &in.PreventDuplicateNames, &out.PreventDuplicateNames
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnitObservation.
func (in *UnitObservation) DeepCopy() *UnitObservation {
	if in == nil {
		return nil
	}
	out := new(UnitObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnitParameters) DeepCopyInto(out *UnitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.HiddenMembershipEnabled != nil {
		in, out := &in.HiddenMembershipEnabled, &out.HiddenMembershipEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PreventDuplicateNames != nil {
		in, out := &in.PreventDuplicateNames, &out.PreventDuplicateNames
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnitParameters.
func (in *UnitParameters) DeepCopy() *UnitParameters {
	if in == nil {
		return nil
	}
	out := new(UnitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnitSpec) DeepCopyInto(out *UnitSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnitSpec.
func (in *UnitSpec) DeepCopy() *UnitSpec {
	if in == nil {
		return nil
	}
	out := new(UnitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnitStatus) DeepCopyInto(out *UnitStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnitStatus.
func (in *UnitStatus) DeepCopy() *UnitStatus {
	if in == nil {
		return nil
	}
	out := new(UnitStatus)
	in.DeepCopyInto(out)
	return out
}
