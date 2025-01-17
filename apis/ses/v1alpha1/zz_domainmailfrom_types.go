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

type DomainMailFromObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DomainMailFromParameters struct {

	// +kubebuilder:validation:Optional
	BehaviorOnMxFailure *string `json:"behaviorOnMxFailure,omitempty" tf:"behavior_on_mx_failure,omitempty"`

	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// +kubebuilder:validation:Required
	MailFromDomain *string `json:"mailFromDomain" tf:"mail_from_domain,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DomainMailFromSpec defines the desired state of DomainMailFrom
type DomainMailFromSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainMailFromParameters `json:"forProvider"`
}

// DomainMailFromStatus defines the observed state of DomainMailFrom.
type DomainMailFromStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainMailFromObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainMailFrom is the Schema for the DomainMailFroms API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type DomainMailFrom struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainMailFromSpec   `json:"spec"`
	Status            DomainMailFromStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainMailFromList contains a list of DomainMailFroms
type DomainMailFromList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainMailFrom `json:"items"`
}

// Repository type metadata.
var (
	DomainMailFrom_Kind             = "DomainMailFrom"
	DomainMailFrom_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainMailFrom_Kind}.String()
	DomainMailFrom_KindAPIVersion   = DomainMailFrom_Kind + "." + CRDGroupVersion.String()
	DomainMailFrom_GroupVersionKind = CRDGroupVersion.WithKind(DomainMailFrom_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainMailFrom{}, &DomainMailFromList{})
}
