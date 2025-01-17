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

type DNSTargetResourceObservation struct {
}

type DNSTargetResourceParameters struct {

	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Optional
	HostedZoneArn *string `json:"hostedZoneArn,omitempty" tf:"hosted_zone_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RecordSetID *string `json:"recordSetId,omitempty" tf:"record_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	RecordType *string `json:"recordType,omitempty" tf:"record_type,omitempty"`

	// +kubebuilder:validation:Optional
	TargetResource []TargetResourceParameters `json:"targetResource,omitempty" tf:"target_resource,omitempty"`
}

type NlbResourceObservation struct {
}

type NlbResourceParameters struct {

	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type R53ResourceObservation struct {
}

type R53ResourceParameters struct {

	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Optional
	RecordSetID *string `json:"recordSetId,omitempty" tf:"record_set_id,omitempty"`
}

type ResourceSetObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ResourceSetParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceSetName *string `json:"resourceSetName" tf:"resource_set_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceSetType *string `json:"resourceSetType" tf:"resource_set_type,omitempty"`

	// +kubebuilder:validation:Required
	Resources []ResourcesParameters `json:"resources" tf:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ResourcesObservation struct {
	ComponentID *string `json:"componentId,omitempty" tf:"component_id,omitempty"`
}

type ResourcesParameters struct {

	// +kubebuilder:validation:Optional
	DNSTargetResource []DNSTargetResourceParameters `json:"dnsTargetResource,omitempty" tf:"dns_target_resource,omitempty"`

	// +kubebuilder:validation:Optional
	ReadinessScopes []*string `json:"readinessScopes,omitempty" tf:"readiness_scopes,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`
}

type TargetResourceObservation struct {
}

type TargetResourceParameters struct {

	// +kubebuilder:validation:Optional
	NlbResource []NlbResourceParameters `json:"nlbResource,omitempty" tf:"nlb_resource,omitempty"`

	// +kubebuilder:validation:Optional
	R53Resource []R53ResourceParameters `json:"r53Resource,omitempty" tf:"r53_resource,omitempty"`
}

// ResourceSetSpec defines the desired state of ResourceSet
type ResourceSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceSetParameters `json:"forProvider"`
}

// ResourceSetStatus defines the observed state of ResourceSet.
type ResourceSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceSet is the Schema for the ResourceSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ResourceSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceSetSpec   `json:"spec"`
	Status            ResourceSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceSetList contains a list of ResourceSets
type ResourceSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceSet `json:"items"`
}

// Repository type metadata.
var (
	ResourceSet_Kind             = "ResourceSet"
	ResourceSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceSet_Kind}.String()
	ResourceSet_KindAPIVersion   = ResourceSet_Kind + "." + CRDGroupVersion.String()
	ResourceSet_GroupVersionKind = CRDGroupVersion.WithKind(ResourceSet_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceSet{}, &ResourceSetList{})
}
