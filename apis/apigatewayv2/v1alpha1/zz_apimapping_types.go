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

type APIMappingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type APIMappingParameters struct {

	// +kubebuilder:validation:Required
	APIID *string `json:"apiId" tf:"api_id,omitempty"`

	// +kubebuilder:validation:Optional
	APIMappingKey *string `json:"apiMappingKey,omitempty" tf:"api_mapping_key,omitempty"`

	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Stage *string `json:"stage" tf:"stage,omitempty"`
}

// APIMappingSpec defines the desired state of APIMapping
type APIMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIMappingParameters `json:"forProvider"`
}

// APIMappingStatus defines the observed state of APIMapping.
type APIMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIMapping is the Schema for the APIMappings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type APIMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIMappingSpec   `json:"spec"`
	Status            APIMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIMappingList contains a list of APIMappings
type APIMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIMapping `json:"items"`
}

// Repository type metadata.
var (
	APIMapping_Kind             = "APIMapping"
	APIMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIMapping_Kind}.String()
	APIMapping_KindAPIVersion   = APIMapping_Kind + "." + CRDGroupVersion.String()
	APIMapping_GroupVersionKind = CRDGroupVersion.WithKind(APIMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&APIMapping{}, &APIMappingList{})
}
