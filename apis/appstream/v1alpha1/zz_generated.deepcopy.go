//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessEndpointsObservation) DeepCopyInto(out *AccessEndpointsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessEndpointsObservation.
func (in *AccessEndpointsObservation) DeepCopy() *AccessEndpointsObservation {
	if in == nil {
		return nil
	}
	out := new(AccessEndpointsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessEndpointsParameters) DeepCopyInto(out *AccessEndpointsParameters) {
	*out = *in
	if in.EndpointType != nil {
		in, out := &in.EndpointType, &out.EndpointType
		*out = new(string)
		**out = **in
	}
	if in.VpceID != nil {
		in, out := &in.VpceID, &out.VpceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessEndpointsParameters.
func (in *AccessEndpointsParameters) DeepCopy() *AccessEndpointsParameters {
	if in == nil {
		return nil
	}
	out := new(AccessEndpointsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSettingsObservation) DeepCopyInto(out *ApplicationSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSettingsObservation.
func (in *ApplicationSettingsObservation) DeepCopy() *ApplicationSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(ApplicationSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSettingsParameters) DeepCopyInto(out *ApplicationSettingsParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.SettingsGroup != nil {
		in, out := &in.SettingsGroup, &out.SettingsGroup
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSettingsParameters.
func (in *ApplicationSettingsParameters) DeepCopy() *ApplicationSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(ApplicationSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stack) DeepCopyInto(out *Stack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stack.
func (in *Stack) DeepCopy() *Stack {
	if in == nil {
		return nil
	}
	out := new(Stack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackList) DeepCopyInto(out *StackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackList.
func (in *StackList) DeepCopy() *StackList {
	if in == nil {
		return nil
	}
	out := new(StackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackObservation) DeepCopyInto(out *StackObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CreatedTime != nil {
		in, out := &in.CreatedTime, &out.CreatedTime
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackObservation.
func (in *StackObservation) DeepCopy() *StackObservation {
	if in == nil {
		return nil
	}
	out := new(StackObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackParameters) DeepCopyInto(out *StackParameters) {
	*out = *in
	if in.AccessEndpoints != nil {
		in, out := &in.AccessEndpoints, &out.AccessEndpoints
		*out = make([]AccessEndpointsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ApplicationSettings != nil {
		in, out := &in.ApplicationSettings, &out.ApplicationSettings
		*out = make([]ApplicationSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
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
	if in.EmbedHostDomains != nil {
		in, out := &in.EmbedHostDomains, &out.EmbedHostDomains
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.FeedbackURL != nil {
		in, out := &in.FeedbackURL, &out.FeedbackURL
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RedirectURL != nil {
		in, out := &in.RedirectURL, &out.RedirectURL
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.StorageConnectors != nil {
		in, out := &in.StorageConnectors, &out.StorageConnectors
		*out = make([]StorageConnectorsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.UserSettings != nil {
		in, out := &in.UserSettings, &out.UserSettings
		*out = make([]UserSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackParameters.
func (in *StackParameters) DeepCopy() *StackParameters {
	if in == nil {
		return nil
	}
	out := new(StackParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSpec) DeepCopyInto(out *StackSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSpec.
func (in *StackSpec) DeepCopy() *StackSpec {
	if in == nil {
		return nil
	}
	out := new(StackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackStatus) DeepCopyInto(out *StackStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackStatus.
func (in *StackStatus) DeepCopy() *StackStatus {
	if in == nil {
		return nil
	}
	out := new(StackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConnectorsObservation) DeepCopyInto(out *StorageConnectorsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConnectorsObservation.
func (in *StorageConnectorsObservation) DeepCopy() *StorageConnectorsObservation {
	if in == nil {
		return nil
	}
	out := new(StorageConnectorsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConnectorsParameters) DeepCopyInto(out *StorageConnectorsParameters) {
	*out = *in
	if in.ConnectorType != nil {
		in, out := &in.ConnectorType, &out.ConnectorType
		*out = new(string)
		**out = **in
	}
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ResourceIdentifier != nil {
		in, out := &in.ResourceIdentifier, &out.ResourceIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConnectorsParameters.
func (in *StorageConnectorsParameters) DeepCopy() *StorageConnectorsParameters {
	if in == nil {
		return nil
	}
	out := new(StorageConnectorsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSettingsObservation) DeepCopyInto(out *UserSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSettingsObservation.
func (in *UserSettingsObservation) DeepCopy() *UserSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(UserSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSettingsParameters) DeepCopyInto(out *UserSettingsParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Permission != nil {
		in, out := &in.Permission, &out.Permission
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSettingsParameters.
func (in *UserSettingsParameters) DeepCopy() *UserSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(UserSettingsParameters)
	in.DeepCopyInto(out)
	return out
}
