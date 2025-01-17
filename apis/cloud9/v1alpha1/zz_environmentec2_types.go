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

type EnvironmentEC2Observation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EnvironmentEC2Parameters struct {

	// +kubebuilder:validation:Optional
	AutomaticStopTimeMinutes *float64 `json:"automaticStopTimeMinutes,omitempty" tf:"automatic_stop_time_minutes,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	OwnerArn *string `json:"ownerArn,omitempty" tf:"owner_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// EnvironmentEC2Spec defines the desired state of EnvironmentEC2
type EnvironmentEC2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentEC2Parameters `json:"forProvider"`
}

// EnvironmentEC2Status defines the observed state of EnvironmentEC2.
type EnvironmentEC2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentEC2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentEC2 is the Schema for the EnvironmentEC2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type EnvironmentEC2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EnvironmentEC2Spec   `json:"spec"`
	Status            EnvironmentEC2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentEC2List contains a list of EnvironmentEC2s
type EnvironmentEC2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnvironmentEC2 `json:"items"`
}

// Repository type metadata.
var (
	EnvironmentEC2_Kind             = "EnvironmentEC2"
	EnvironmentEC2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EnvironmentEC2_Kind}.String()
	EnvironmentEC2_KindAPIVersion   = EnvironmentEC2_Kind + "." + CRDGroupVersion.String()
	EnvironmentEC2_GroupVersionKind = CRDGroupVersion.WithKind(EnvironmentEC2_Kind)
)

func init() {
	SchemeBuilder.Register(&EnvironmentEC2{}, &EnvironmentEC2List{})
}
