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

type SnapshotCreateVolumePermissionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SnapshotCreateVolumePermissionParameters struct {

	// +kubebuilder:validation:Required
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SnapshotID *string `json:"snapshotId" tf:"snapshot_id,omitempty"`
}

// SnapshotCreateVolumePermissionSpec defines the desired state of SnapshotCreateVolumePermission
type SnapshotCreateVolumePermissionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotCreateVolumePermissionParameters `json:"forProvider"`
}

// SnapshotCreateVolumePermissionStatus defines the observed state of SnapshotCreateVolumePermission.
type SnapshotCreateVolumePermissionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotCreateVolumePermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotCreateVolumePermission is the Schema for the SnapshotCreateVolumePermissions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type SnapshotCreateVolumePermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotCreateVolumePermissionSpec   `json:"spec"`
	Status            SnapshotCreateVolumePermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotCreateVolumePermissionList contains a list of SnapshotCreateVolumePermissions
type SnapshotCreateVolumePermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotCreateVolumePermission `json:"items"`
}

// Repository type metadata.
var (
	SnapshotCreateVolumePermission_Kind             = "SnapshotCreateVolumePermission"
	SnapshotCreateVolumePermission_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnapshotCreateVolumePermission_Kind}.String()
	SnapshotCreateVolumePermission_KindAPIVersion   = SnapshotCreateVolumePermission_Kind + "." + CRDGroupVersion.String()
	SnapshotCreateVolumePermission_GroupVersionKind = CRDGroupVersion.WithKind(SnapshotCreateVolumePermission_Kind)
)

func init() {
	SchemeBuilder.Register(&SnapshotCreateVolumePermission{}, &SnapshotCreateVolumePermissionList{})
}
