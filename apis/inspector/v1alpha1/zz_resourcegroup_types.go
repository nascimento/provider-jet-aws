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

type ResourceGroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ResourceGroupParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Tags map[string]*string `json:"tags" tf:"tags,omitempty"`
}

// ResourceGroupSpec defines the desired state of ResourceGroup
type ResourceGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceGroupParameters `json:"forProvider"`
}

// ResourceGroupStatus defines the observed state of ResourceGroup.
type ResourceGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroup is the Schema for the ResourceGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceGroupSpec   `json:"spec"`
	Status            ResourceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupList contains a list of ResourceGroups
type ResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroup `json:"items"`
}

// Repository type metadata.
var (
	ResourceGroup_Kind             = "ResourceGroup"
	ResourceGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceGroup_Kind}.String()
	ResourceGroup_KindAPIVersion   = ResourceGroup_Kind + "." + CRDGroupVersion.String()
	ResourceGroup_GroupVersionKind = CRDGroupVersion.WithKind(ResourceGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceGroup{}, &ResourceGroupList{})
}