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

type RepositoryPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`
}

type RepositoryPolicyParameters struct {

	// +kubebuilder:validation:Required
	Policy *string `json:"policy" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Repository *string `json:"repository" tf:"repository,omitempty"`
}

// RepositoryPolicySpec defines the desired state of RepositoryPolicy
type RepositoryPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryPolicyParameters `json:"forProvider"`
}

// RepositoryPolicyStatus defines the observed state of RepositoryPolicy.
type RepositoryPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryPolicy is the Schema for the RepositoryPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type RepositoryPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RepositoryPolicySpec   `json:"spec"`
	Status            RepositoryPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryPolicyList contains a list of RepositoryPolicys
type RepositoryPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RepositoryPolicy `json:"items"`
}

// Repository type metadata.
var (
	RepositoryPolicy_Kind             = "RepositoryPolicy"
	RepositoryPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RepositoryPolicy_Kind}.String()
	RepositoryPolicy_KindAPIVersion   = RepositoryPolicy_Kind + "." + CRDGroupVersion.String()
	RepositoryPolicy_GroupVersionKind = CRDGroupVersion.WithKind(RepositoryPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&RepositoryPolicy{}, &RepositoryPolicyList{})
}
