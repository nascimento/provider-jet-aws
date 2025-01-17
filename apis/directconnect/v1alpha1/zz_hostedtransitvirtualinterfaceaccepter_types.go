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

type HostedTransitVirtualInterfaceAccepterObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type HostedTransitVirtualInterfaceAccepterParameters struct {

	// +kubebuilder:validation:Required
	DxGatewayID *string `json:"dxGatewayId" tf:"dx_gateway_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VirtualInterfaceID *string `json:"virtualInterfaceId" tf:"virtual_interface_id,omitempty"`
}

// HostedTransitVirtualInterfaceAccepterSpec defines the desired state of HostedTransitVirtualInterfaceAccepter
type HostedTransitVirtualInterfaceAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostedTransitVirtualInterfaceAccepterParameters `json:"forProvider"`
}

// HostedTransitVirtualInterfaceAccepterStatus defines the observed state of HostedTransitVirtualInterfaceAccepter.
type HostedTransitVirtualInterfaceAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostedTransitVirtualInterfaceAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostedTransitVirtualInterfaceAccepter is the Schema for the HostedTransitVirtualInterfaceAccepters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type HostedTransitVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostedTransitVirtualInterfaceAccepterSpec   `json:"spec"`
	Status            HostedTransitVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostedTransitVirtualInterfaceAccepterList contains a list of HostedTransitVirtualInterfaceAccepters
type HostedTransitVirtualInterfaceAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostedTransitVirtualInterfaceAccepter `json:"items"`
}

// Repository type metadata.
var (
	HostedTransitVirtualInterfaceAccepter_Kind             = "HostedTransitVirtualInterfaceAccepter"
	HostedTransitVirtualInterfaceAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostedTransitVirtualInterfaceAccepter_Kind}.String()
	HostedTransitVirtualInterfaceAccepter_KindAPIVersion   = HostedTransitVirtualInterfaceAccepter_Kind + "." + CRDGroupVersion.String()
	HostedTransitVirtualInterfaceAccepter_GroupVersionKind = CRDGroupVersion.WithKind(HostedTransitVirtualInterfaceAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&HostedTransitVirtualInterfaceAccepter{}, &HostedTransitVirtualInterfaceAccepterList{})
}
