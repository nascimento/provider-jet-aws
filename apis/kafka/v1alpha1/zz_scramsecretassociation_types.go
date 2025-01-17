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

type ScramSecretAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ScramSecretAssociationParameters struct {

	// +kubebuilder:validation:Required
	ClusterArn *string `json:"clusterArn" tf:"cluster_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SecretArnList []*string `json:"secretArnList" tf:"secret_arn_list,omitempty"`
}

// ScramSecretAssociationSpec defines the desired state of ScramSecretAssociation
type ScramSecretAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScramSecretAssociationParameters `json:"forProvider"`
}

// ScramSecretAssociationStatus defines the observed state of ScramSecretAssociation.
type ScramSecretAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScramSecretAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ScramSecretAssociation is the Schema for the ScramSecretAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ScramSecretAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScramSecretAssociationSpec   `json:"spec"`
	Status            ScramSecretAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScramSecretAssociationList contains a list of ScramSecretAssociations
type ScramSecretAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScramSecretAssociation `json:"items"`
}

// Repository type metadata.
var (
	ScramSecretAssociation_Kind             = "ScramSecretAssociation"
	ScramSecretAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScramSecretAssociation_Kind}.String()
	ScramSecretAssociation_KindAPIVersion   = ScramSecretAssociation_Kind + "." + CRDGroupVersion.String()
	ScramSecretAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ScramSecretAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ScramSecretAssociation{}, &ScramSecretAssociationList{})
}
