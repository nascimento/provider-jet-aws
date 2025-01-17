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

type HostedConfigurationVersionObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	VersionNumber *float64 `json:"versionNumber,omitempty" tf:"version_number,omitempty"`
}

type HostedConfigurationVersionParameters struct {

	// +kubebuilder:validation:Required
	ApplicationID *string `json:"applicationId" tf:"application_id,omitempty"`

	// +kubebuilder:validation:Required
	ConfigurationProfileID *string `json:"configurationProfileId" tf:"configuration_profile_id,omitempty"`

	// +kubebuilder:validation:Required
	ContentSecretRef v1.SecretKeySelector `json:"contentSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// HostedConfigurationVersionSpec defines the desired state of HostedConfigurationVersion
type HostedConfigurationVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostedConfigurationVersionParameters `json:"forProvider"`
}

// HostedConfigurationVersionStatus defines the observed state of HostedConfigurationVersion.
type HostedConfigurationVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostedConfigurationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostedConfigurationVersion is the Schema for the HostedConfigurationVersions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type HostedConfigurationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostedConfigurationVersionSpec   `json:"spec"`
	Status            HostedConfigurationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostedConfigurationVersionList contains a list of HostedConfigurationVersions
type HostedConfigurationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostedConfigurationVersion `json:"items"`
}

// Repository type metadata.
var (
	HostedConfigurationVersion_Kind             = "HostedConfigurationVersion"
	HostedConfigurationVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostedConfigurationVersion_Kind}.String()
	HostedConfigurationVersion_KindAPIVersion   = HostedConfigurationVersion_Kind + "." + CRDGroupVersion.String()
	HostedConfigurationVersion_GroupVersionKind = CRDGroupVersion.WithKind(HostedConfigurationVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&HostedConfigurationVersion{}, &HostedConfigurationVersionList{})
}
