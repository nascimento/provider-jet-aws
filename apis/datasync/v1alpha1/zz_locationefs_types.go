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

type EC2ConfigObservation struct {
}

type EC2ConfigParameters struct {

	// +kubebuilder:validation:Required
	SecurityGroupArns []*string `json:"securityGroupArns" tf:"security_group_arns,omitempty"`

	// +kubebuilder:validation:Required
	SubnetArn *string `json:"subnetArn" tf:"subnet_arn,omitempty"`
}

type LocationEFSObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type LocationEFSParameters struct {

	// +kubebuilder:validation:Required
	EC2Config []EC2ConfigParameters `json:"ec2Config" tf:"ec2_config,omitempty"`

	// +kubebuilder:validation:Required
	EFSFileSystemArn *string `json:"efsFileSystemArn" tf:"efs_file_system_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Subdirectory *string `json:"subdirectory,omitempty" tf:"subdirectory,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LocationEFSSpec defines the desired state of LocationEFS
type LocationEFSSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LocationEFSParameters `json:"forProvider"`
}

// LocationEFSStatus defines the observed state of LocationEFS.
type LocationEFSStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LocationEFSObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LocationEFS is the Schema for the LocationEFSs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type LocationEFS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocationEFSSpec   `json:"spec"`
	Status            LocationEFSStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocationEFSList contains a list of LocationEFSs
type LocationEFSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocationEFS `json:"items"`
}

// Repository type metadata.
var (
	LocationEFS_Kind             = "LocationEFS"
	LocationEFS_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LocationEFS_Kind}.String()
	LocationEFS_KindAPIVersion   = LocationEFS_Kind + "." + CRDGroupVersion.String()
	LocationEFS_GroupVersionKind = CRDGroupVersion.WithKind(LocationEFS_Kind)
)

func init() {
	SchemeBuilder.Register(&LocationEFS{}, &LocationEFSList{})
}
