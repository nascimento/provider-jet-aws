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

type ResourceDataSyncObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ResourceDataSyncParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	S3Destination []S3DestinationParameters `json:"s3Destination" tf:"s3_destination,omitempty"`
}

type S3DestinationObservation struct {
}

type S3DestinationParameters struct {

	// +kubebuilder:validation:Required
	BucketName *string `json:"bucketName" tf:"bucket_name,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	SyncFormat *string `json:"syncFormat,omitempty" tf:"sync_format,omitempty"`
}

// ResourceDataSyncSpec defines the desired state of ResourceDataSync
type ResourceDataSyncSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceDataSyncParameters `json:"forProvider"`
}

// ResourceDataSyncStatus defines the observed state of ResourceDataSync.
type ResourceDataSyncStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceDataSyncObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceDataSync is the Schema for the ResourceDataSyncs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ResourceDataSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceDataSyncSpec   `json:"spec"`
	Status            ResourceDataSyncStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceDataSyncList contains a list of ResourceDataSyncs
type ResourceDataSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceDataSync `json:"items"`
}

// Repository type metadata.
var (
	ResourceDataSync_Kind             = "ResourceDataSync"
	ResourceDataSync_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceDataSync_Kind}.String()
	ResourceDataSync_KindAPIVersion   = ResourceDataSync_Kind + "." + CRDGroupVersion.String()
	ResourceDataSync_GroupVersionKind = CRDGroupVersion.WithKind(ResourceDataSync_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceDataSync{}, &ResourceDataSyncList{})
}
