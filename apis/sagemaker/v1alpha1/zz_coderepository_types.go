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

type CodeRepositoryObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CodeRepositoryParameters struct {

	// +kubebuilder:validation:Required
	CodeRepositoryName *string `json:"codeRepositoryName" tf:"code_repository_name,omitempty"`

	// +kubebuilder:validation:Required
	GitConfig []GitConfigParameters `json:"gitConfig" tf:"git_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type GitConfigObservation struct {
}

type GitConfigParameters struct {

	// +kubebuilder:validation:Optional
	Branch *string `json:"branch,omitempty" tf:"branch,omitempty"`

	// +kubebuilder:validation:Required
	RepositoryURL *string `json:"repositoryUrl" tf:"repository_url,omitempty"`

	// +kubebuilder:validation:Optional
	SecretArn *string `json:"secretArn,omitempty" tf:"secret_arn,omitempty"`
}

// CodeRepositorySpec defines the desired state of CodeRepository
type CodeRepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CodeRepositoryParameters `json:"forProvider"`
}

// CodeRepositoryStatus defines the observed state of CodeRepository.
type CodeRepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CodeRepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodeRepository is the Schema for the CodeRepositorys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type CodeRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodeRepositorySpec   `json:"spec"`
	Status            CodeRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodeRepositoryList contains a list of CodeRepositorys
type CodeRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodeRepository `json:"items"`
}

// Repository type metadata.
var (
	CodeRepository_Kind             = "CodeRepository"
	CodeRepository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CodeRepository_Kind}.String()
	CodeRepository_KindAPIVersion   = CodeRepository_Kind + "." + CRDGroupVersion.String()
	CodeRepository_GroupVersionKind = CRDGroupVersion.WithKind(CodeRepository_Kind)
)

func init() {
	SchemeBuilder.Register(&CodeRepository{}, &CodeRepositoryList{})
}
