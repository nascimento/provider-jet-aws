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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LocationFSXWindowsFileSystemObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type LocationFSXWindowsFileSystemParameters struct {

	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// +kubebuilder:validation:Required
	FSXFilesystemArn *string `json:"fsxFilesystemArn" tf:"fsx_filesystem_arn,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SecurityGroupArns []*string `json:"securityGroupArns" tf:"security_group_arns,omitempty"`

	// +kubebuilder:validation:Optional
	Subdirectory *string `json:"subdirectory,omitempty" tf:"subdirectory,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	User *string `json:"user" tf:"user,omitempty"`
}

// LocationFSXWindowsFileSystemSpec defines the desired state of LocationFSXWindowsFileSystem
type LocationFSXWindowsFileSystemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LocationFSXWindowsFileSystemParameters `json:"forProvider"`
}

// LocationFSXWindowsFileSystemStatus defines the observed state of LocationFSXWindowsFileSystem.
type LocationFSXWindowsFileSystemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LocationFSXWindowsFileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LocationFSXWindowsFileSystem is the Schema for the LocationFSXWindowsFileSystems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type LocationFSXWindowsFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocationFSXWindowsFileSystemSpec   `json:"spec"`
	Status            LocationFSXWindowsFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocationFSXWindowsFileSystemList contains a list of LocationFSXWindowsFileSystems
type LocationFSXWindowsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocationFSXWindowsFileSystem `json:"items"`
}

// Repository type metadata.
var (
	LocationFSXWindowsFileSystem_Kind             = "LocationFSXWindowsFileSystem"
	LocationFSXWindowsFileSystem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LocationFSXWindowsFileSystem_Kind}.String()
	LocationFSXWindowsFileSystem_KindAPIVersion   = LocationFSXWindowsFileSystem_Kind + "." + CRDGroupVersion.String()
	LocationFSXWindowsFileSystem_GroupVersionKind = CRDGroupVersion.WithKind(LocationFSXWindowsFileSystem_Kind)
)

func init() {
	SchemeBuilder.Register(&LocationFSXWindowsFileSystem{}, &LocationFSXWindowsFileSystemList{})
}
