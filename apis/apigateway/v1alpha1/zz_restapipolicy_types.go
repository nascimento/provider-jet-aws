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

type RestAPIPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RestAPIPolicyParameters struct {

	// +kubebuilder:validation:Required
	Policy *string `json:"policy" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	RestAPIID *string `json:"restApiId" tf:"rest_api_id,omitempty"`
}

// RestAPIPolicySpec defines the desired state of RestAPIPolicy
type RestAPIPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RestAPIPolicyParameters `json:"forProvider"`
}

// RestAPIPolicyStatus defines the observed state of RestAPIPolicy.
type RestAPIPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RestAPIPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RestAPIPolicy is the Schema for the RestAPIPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type RestAPIPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RestAPIPolicySpec   `json:"spec"`
	Status            RestAPIPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RestAPIPolicyList contains a list of RestAPIPolicys
type RestAPIPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RestAPIPolicy `json:"items"`
}

// Repository type metadata.
var (
	RestAPIPolicy_Kind             = "RestAPIPolicy"
	RestAPIPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RestAPIPolicy_Kind}.String()
	RestAPIPolicy_KindAPIVersion   = RestAPIPolicy_Kind + "." + CRDGroupVersion.String()
	RestAPIPolicy_GroupVersionKind = CRDGroupVersion.WithKind(RestAPIPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&RestAPIPolicy{}, &RestAPIPolicyList{})
}