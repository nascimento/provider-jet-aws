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

type EBSSnapshotCopyObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DataEncryptionKeyID *string `json:"dataEncryptionKeyId,omitempty" tf:"data_encryption_key_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OwnerAlias *string `json:"ownerAlias,omitempty" tf:"owner_alias,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
}

type EBSSnapshotCopyParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha2.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SourceRegion *string `json:"sourceRegion" tf:"source_region,omitempty"`

	// +kubebuilder:validation:Required
	SourceSnapshotID *string `json:"sourceSnapshotId" tf:"source_snapshot_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// EBSSnapshotCopySpec defines the desired state of EBSSnapshotCopy
type EBSSnapshotCopySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EBSSnapshotCopyParameters `json:"forProvider"`
}

// EBSSnapshotCopyStatus defines the observed state of EBSSnapshotCopy.
type EBSSnapshotCopyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EBSSnapshotCopyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EBSSnapshotCopy is the Schema for the EBSSnapshotCopys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type EBSSnapshotCopy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EBSSnapshotCopySpec   `json:"spec"`
	Status            EBSSnapshotCopyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EBSSnapshotCopyList contains a list of EBSSnapshotCopys
type EBSSnapshotCopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EBSSnapshotCopy `json:"items"`
}

// Repository type metadata.
var (
	EBSSnapshotCopy_Kind             = "EBSSnapshotCopy"
	EBSSnapshotCopy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EBSSnapshotCopy_Kind}.String()
	EBSSnapshotCopy_KindAPIVersion   = EBSSnapshotCopy_Kind + "." + CRDGroupVersion.String()
	EBSSnapshotCopy_GroupVersionKind = CRDGroupVersion.WithKind(EBSSnapshotCopy_Kind)
)

func init() {
	SchemeBuilder.Register(&EBSSnapshotCopy{}, &EBSSnapshotCopyList{})
}
