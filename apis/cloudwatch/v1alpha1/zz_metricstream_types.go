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

type ExcludeFilterObservation struct {
}

type ExcludeFilterParameters struct {

	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

type IncludeFilterObservation struct {
}

type IncludeFilterParameters struct {

	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

type MetricStreamObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LastUpdateDate *string `json:"lastUpdateDate,omitempty" tf:"last_update_date,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type MetricStreamParameters struct {

	// +kubebuilder:validation:Optional
	ExcludeFilter []ExcludeFilterParameters `json:"excludeFilter,omitempty" tf:"exclude_filter,omitempty"`

	// +kubebuilder:validation:Required
	FirehoseArn *string `json:"firehoseArn" tf:"firehose_arn,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeFilter []IncludeFilterParameters `json:"includeFilter,omitempty" tf:"include_filter,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	OutputFormat *string `json:"outputFormat" tf:"output_format,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// MetricStreamSpec defines the desired state of MetricStream
type MetricStreamSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MetricStreamParameters `json:"forProvider"`
}

// MetricStreamStatus defines the observed state of MetricStream.
type MetricStreamStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MetricStreamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MetricStream is the Schema for the MetricStreams API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type MetricStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MetricStreamSpec   `json:"spec"`
	Status            MetricStreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MetricStreamList contains a list of MetricStreams
type MetricStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetricStream `json:"items"`
}

// Repository type metadata.
var (
	MetricStream_Kind             = "MetricStream"
	MetricStream_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MetricStream_Kind}.String()
	MetricStream_KindAPIVersion   = MetricStream_Kind + "." + CRDGroupVersion.String()
	MetricStream_GroupVersionKind = CRDGroupVersion.WithKind(MetricStream_Kind)
)

func init() {
	SchemeBuilder.Register(&MetricStream{}, &MetricStreamList{})
}
