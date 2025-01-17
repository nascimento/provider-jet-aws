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

type LagObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	HasLogicalRedundancy *string `json:"hasLogicalRedundancy,omitempty" tf:"has_logical_redundancy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	JumboFrameCapable *bool `json:"jumboFrameCapable,omitempty" tf:"jumbo_frame_capable,omitempty"`

	OwnerAccountID *string `json:"ownerAccountId,omitempty" tf:"owner_account_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LagParameters struct {

	// +kubebuilder:validation:Required
	ConnectionsBandwidth *string `json:"connectionsBandwidth" tf:"connections_bandwidth,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LagSpec defines the desired state of Lag
type LagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LagParameters `json:"forProvider"`
}

// LagStatus defines the observed state of Lag.
type LagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Lag is the Schema for the Lags API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Lag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LagSpec   `json:"spec"`
	Status            LagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LagList contains a list of Lags
type LagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Lag `json:"items"`
}

// Repository type metadata.
var (
	Lag_Kind             = "Lag"
	Lag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Lag_Kind}.String()
	Lag_KindAPIVersion   = Lag_Kind + "." + CRDGroupVersion.String()
	Lag_GroupVersionKind = CRDGroupVersion.WithKind(Lag_Kind)
)

func init() {
	SchemeBuilder.Register(&Lag{}, &LagList{})
}
