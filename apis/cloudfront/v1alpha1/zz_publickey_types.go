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

type PublicKeyObservation struct {
	CallerReference *string `json:"callerReference,omitempty" tf:"caller_reference,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PublicKeyParameters struct {

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Required
	EncodedKey *string `json:"encodedKey" tf:"encoded_key,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// PublicKeySpec defines the desired state of PublicKey
type PublicKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicKeyParameters `json:"forProvider"`
}

// PublicKeyStatus defines the observed state of PublicKey.
type PublicKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PublicKey is the Schema for the PublicKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type PublicKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PublicKeySpec   `json:"spec"`
	Status            PublicKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicKeyList contains a list of PublicKeys
type PublicKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicKey `json:"items"`
}

// Repository type metadata.
var (
	PublicKey_Kind             = "PublicKey"
	PublicKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PublicKey_Kind}.String()
	PublicKey_KindAPIVersion   = PublicKey_Kind + "." + CRDGroupVersion.String()
	PublicKey_GroupVersionKind = CRDGroupVersion.WithKind(PublicKey_Kind)
)

func init() {
	SchemeBuilder.Register(&PublicKey{}, &PublicKeyList{})
}
