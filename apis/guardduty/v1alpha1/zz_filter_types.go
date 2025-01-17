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

type CriterionObservation struct {
}

type CriterionParameters struct {

	// +kubebuilder:validation:Optional
	Equals []*string `json:"equals,omitempty" tf:"equals,omitempty"`

	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// +kubebuilder:validation:Optional
	GreaterThan *string `json:"greaterThan,omitempty" tf:"greater_than,omitempty"`

	// +kubebuilder:validation:Optional
	GreaterThanOrEqual *string `json:"greaterThanOrEqual,omitempty" tf:"greater_than_or_equal,omitempty"`

	// +kubebuilder:validation:Optional
	LessThan *string `json:"lessThan,omitempty" tf:"less_than,omitempty"`

	// +kubebuilder:validation:Optional
	LessThanOrEqual *string `json:"lessThanOrEqual,omitempty" tf:"less_than_or_equal,omitempty"`

	// +kubebuilder:validation:Optional
	NotEquals []*string `json:"notEquals,omitempty" tf:"not_equals,omitempty"`
}

type FilterObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type FilterParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	DetectorID *string `json:"detectorId" tf:"detector_id,omitempty"`

	// +kubebuilder:validation:Required
	FindingCriteria []FindingCriteriaParameters `json:"findingCriteria" tf:"finding_criteria,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Rank *float64 `json:"rank" tf:"rank,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FindingCriteriaObservation struct {
}

type FindingCriteriaParameters struct {

	// +kubebuilder:validation:Required
	Criterion []CriterionParameters `json:"criterion" tf:"criterion,omitempty"`
}

// FilterSpec defines the desired state of Filter
type FilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FilterParameters `json:"forProvider"`
}

// FilterStatus defines the observed state of Filter.
type FilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Filter is the Schema for the Filters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Filter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FilterSpec   `json:"spec"`
	Status            FilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FilterList contains a list of Filters
type FilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Filter `json:"items"`
}

// Repository type metadata.
var (
	Filter_Kind             = "Filter"
	Filter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Filter_Kind}.String()
	Filter_KindAPIVersion   = Filter_Kind + "." + CRDGroupVersion.String()
	Filter_GroupVersionKind = CRDGroupVersion.WithKind(Filter_Kind)
)

func init() {
	SchemeBuilder.Register(&Filter{}, &FilterList{})
}
