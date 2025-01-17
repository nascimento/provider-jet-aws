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

type ReadinessCheckObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ReadinessCheckParameters struct {

	// +kubebuilder:validation:Required
	ReadinessCheckName *string `json:"readinessCheckName" tf:"readiness_check_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceSetName *string `json:"resourceSetName" tf:"resource_set_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ReadinessCheckSpec defines the desired state of ReadinessCheck
type ReadinessCheckSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReadinessCheckParameters `json:"forProvider"`
}

// ReadinessCheckStatus defines the observed state of ReadinessCheck.
type ReadinessCheckStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReadinessCheckObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReadinessCheck is the Schema for the ReadinessChecks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ReadinessCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReadinessCheckSpec   `json:"spec"`
	Status            ReadinessCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReadinessCheckList contains a list of ReadinessChecks
type ReadinessCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReadinessCheck `json:"items"`
}

// Repository type metadata.
var (
	ReadinessCheck_Kind             = "ReadinessCheck"
	ReadinessCheck_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReadinessCheck_Kind}.String()
	ReadinessCheck_KindAPIVersion   = ReadinessCheck_Kind + "." + CRDGroupVersion.String()
	ReadinessCheck_GroupVersionKind = CRDGroupVersion.WithKind(ReadinessCheck_Kind)
)

func init() {
	SchemeBuilder.Register(&ReadinessCheck{}, &ReadinessCheckList{})
}
