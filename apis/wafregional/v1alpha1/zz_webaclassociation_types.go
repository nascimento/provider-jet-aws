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

type WebACLAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type WebACLAssociationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`

	// +kubebuilder:validation:Required
	WebACLID *string `json:"webAclId" tf:"web_acl_id,omitempty"`
}

// WebACLAssociationSpec defines the desired state of WebACLAssociation
type WebACLAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebACLAssociationParameters `json:"forProvider"`
}

// WebACLAssociationStatus defines the observed state of WebACLAssociation.
type WebACLAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebACLAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WebACLAssociation is the Schema for the WebACLAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type WebACLAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebACLAssociationSpec   `json:"spec"`
	Status            WebACLAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebACLAssociationList contains a list of WebACLAssociations
type WebACLAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebACLAssociation `json:"items"`
}

// Repository type metadata.
var (
	WebACLAssociation_Kind             = "WebACLAssociation"
	WebACLAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebACLAssociation_Kind}.String()
	WebACLAssociation_KindAPIVersion   = WebACLAssociation_Kind + "." + CRDGroupVersion.String()
	WebACLAssociation_GroupVersionKind = CRDGroupVersion.WithKind(WebACLAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&WebACLAssociation{}, &WebACLAssociationList{})
}
