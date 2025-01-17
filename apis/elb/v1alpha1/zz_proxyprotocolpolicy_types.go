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

type ProxyProtocolPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProxyProtocolPolicyParameters struct {

	// +kubebuilder:validation:Required
	InstancePorts []*string `json:"instancePorts" tf:"instance_ports,omitempty"`

	// +kubebuilder:validation:Required
	LoadBalancer *string `json:"loadBalancer" tf:"load_balancer,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ProxyProtocolPolicySpec defines the desired state of ProxyProtocolPolicy
type ProxyProtocolPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProxyProtocolPolicyParameters `json:"forProvider"`
}

// ProxyProtocolPolicyStatus defines the observed state of ProxyProtocolPolicy.
type ProxyProtocolPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProxyProtocolPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProxyProtocolPolicy is the Schema for the ProxyProtocolPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ProxyProtocolPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProxyProtocolPolicySpec   `json:"spec"`
	Status            ProxyProtocolPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProxyProtocolPolicyList contains a list of ProxyProtocolPolicys
type ProxyProtocolPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProxyProtocolPolicy `json:"items"`
}

// Repository type metadata.
var (
	ProxyProtocolPolicy_Kind             = "ProxyProtocolPolicy"
	ProxyProtocolPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProxyProtocolPolicy_Kind}.String()
	ProxyProtocolPolicy_KindAPIVersion   = ProxyProtocolPolicy_Kind + "." + CRDGroupVersion.String()
	ProxyProtocolPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ProxyProtocolPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ProxyProtocolPolicy{}, &ProxyProtocolPolicyList{})
}
