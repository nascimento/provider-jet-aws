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

type ConditionalForwarderObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ConditionalForwarderParameters struct {

	// +kubebuilder:validation:Required
	DNSIps []*string `json:"dnsIps" tf:"dns_ips,omitempty"`

	// +kubebuilder:validation:Required
	DirectoryID *string `json:"directoryId" tf:"directory_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	RemoteDomainName *string `json:"remoteDomainName" tf:"remote_domain_name,omitempty"`
}

// ConditionalForwarderSpec defines the desired state of ConditionalForwarder
type ConditionalForwarderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConditionalForwarderParameters `json:"forProvider"`
}

// ConditionalForwarderStatus defines the observed state of ConditionalForwarder.
type ConditionalForwarderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConditionalForwarderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConditionalForwarder is the Schema for the ConditionalForwarders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ConditionalForwarder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConditionalForwarderSpec   `json:"spec"`
	Status            ConditionalForwarderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConditionalForwarderList contains a list of ConditionalForwarders
type ConditionalForwarderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConditionalForwarder `json:"items"`
}

// Repository type metadata.
var (
	ConditionalForwarder_Kind             = "ConditionalForwarder"
	ConditionalForwarder_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConditionalForwarder_Kind}.String()
	ConditionalForwarder_KindAPIVersion   = ConditionalForwarder_Kind + "." + CRDGroupVersion.String()
	ConditionalForwarder_GroupVersionKind = CRDGroupVersion.WithKind(ConditionalForwarder_Kind)
)

func init() {
	SchemeBuilder.Register(&ConditionalForwarder{}, &ConditionalForwarderList{})
}
