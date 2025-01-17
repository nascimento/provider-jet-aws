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

type RegexPatternSetObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LockToken *string `json:"lockToken,omitempty" tf:"lock_token,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RegexPatternSetParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RegularExpression []RegularExpressionParameters `json:"regularExpression,omitempty" tf:"regular_expression,omitempty"`

	// +kubebuilder:validation:Required
	Scope *string `json:"scope" tf:"scope,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RegularExpressionObservation struct {
}

type RegularExpressionParameters struct {

	// +kubebuilder:validation:Required
	RegexString *string `json:"regexString" tf:"regex_string,omitempty"`
}

// RegexPatternSetSpec defines the desired state of RegexPatternSet
type RegexPatternSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegexPatternSetParameters `json:"forProvider"`
}

// RegexPatternSetStatus defines the observed state of RegexPatternSet.
type RegexPatternSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegexPatternSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegexPatternSet is the Schema for the RegexPatternSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type RegexPatternSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegexPatternSetSpec   `json:"spec"`
	Status            RegexPatternSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegexPatternSetList contains a list of RegexPatternSets
type RegexPatternSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegexPatternSet `json:"items"`
}

// Repository type metadata.
var (
	RegexPatternSet_Kind             = "RegexPatternSet"
	RegexPatternSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegexPatternSet_Kind}.String()
	RegexPatternSet_KindAPIVersion   = RegexPatternSet_Kind + "." + CRDGroupVersion.String()
	RegexPatternSet_GroupVersionKind = CRDGroupVersion.WithKind(RegexPatternSet_Kind)
)

func init() {
	SchemeBuilder.Register(&RegexPatternSet{}, &RegexPatternSetList{})
}
