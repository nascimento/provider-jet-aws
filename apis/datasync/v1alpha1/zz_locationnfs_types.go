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

type LocationNFSObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type LocationNFSParameters struct {

	// +kubebuilder:validation:Optional
	MountOptions []MountOptionsParameters `json:"mountOptions,omitempty" tf:"mount_options,omitempty"`

	// +kubebuilder:validation:Required
	OnPremConfig []OnPremConfigParameters `json:"onPremConfig" tf:"on_prem_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ServerHostname *string `json:"serverHostname" tf:"server_hostname,omitempty"`

	// +kubebuilder:validation:Required
	Subdirectory *string `json:"subdirectory" tf:"subdirectory,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MountOptionsObservation struct {
}

type MountOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type OnPremConfigObservation struct {
}

type OnPremConfigParameters struct {

	// +kubebuilder:validation:Required
	AgentArns []*string `json:"agentArns" tf:"agent_arns,omitempty"`
}

// LocationNFSSpec defines the desired state of LocationNFS
type LocationNFSSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LocationNFSParameters `json:"forProvider"`
}

// LocationNFSStatus defines the observed state of LocationNFS.
type LocationNFSStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LocationNFSObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LocationNFS is the Schema for the LocationNFSs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type LocationNFS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocationNFSSpec   `json:"spec"`
	Status            LocationNFSStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocationNFSList contains a list of LocationNFSs
type LocationNFSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocationNFS `json:"items"`
}

// Repository type metadata.
var (
	LocationNFS_Kind             = "LocationNFS"
	LocationNFS_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LocationNFS_Kind}.String()
	LocationNFS_KindAPIVersion   = LocationNFS_Kind + "." + CRDGroupVersion.String()
	LocationNFS_GroupVersionKind = CRDGroupVersion.WithKind(LocationNFS_Kind)
)

func init() {
	SchemeBuilder.Register(&LocationNFS{}, &LocationNFSList{})
}
