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

type LBListenerCertificateObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LBListenerCertificateParameters struct {

	// +kubebuilder:validation:Required
	CertificateArn *string `json:"certificateArn" tf:"certificate_arn,omitempty"`

	// +kubebuilder:validation:Required
	ListenerArn *string `json:"listenerArn" tf:"listener_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// LBListenerCertificateSpec defines the desired state of LBListenerCertificate
type LBListenerCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBListenerCertificateParameters `json:"forProvider"`
}

// LBListenerCertificateStatus defines the observed state of LBListenerCertificate.
type LBListenerCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBListenerCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBListenerCertificate is the Schema for the LBListenerCertificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type LBListenerCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LBListenerCertificateSpec   `json:"spec"`
	Status            LBListenerCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBListenerCertificateList contains a list of LBListenerCertificates
type LBListenerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBListenerCertificate `json:"items"`
}

// Repository type metadata.
var (
	LBListenerCertificate_Kind             = "LBListenerCertificate"
	LBListenerCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBListenerCertificate_Kind}.String()
	LBListenerCertificate_KindAPIVersion   = LBListenerCertificate_Kind + "." + CRDGroupVersion.String()
	LBListenerCertificate_GroupVersionKind = CRDGroupVersion.WithKind(LBListenerCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&LBListenerCertificate{}, &LBListenerCertificateList{})
}