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

type UserSSHKeyObservation struct {
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SSHPublicKeyID *string `json:"sshPublicKeyId,omitempty" tf:"ssh_public_key_id,omitempty"`
}

type UserSSHKeyParameters struct {

	// +kubebuilder:validation:Required
	Encoding *string `json:"encoding" tf:"encoding,omitempty"`

	// +kubebuilder:validation:Required
	PublicKey *string `json:"publicKey" tf:"public_key,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

// UserSSHKeySpec defines the desired state of UserSSHKey
type UserSSHKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserSSHKeyParameters `json:"forProvider"`
}

// UserSSHKeyStatus defines the observed state of UserSSHKey.
type UserSSHKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserSSHKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserSSHKey is the Schema for the UserSSHKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type UserSSHKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSSHKeySpec   `json:"spec"`
	Status            UserSSHKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserSSHKeyList contains a list of UserSSHKeys
type UserSSHKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserSSHKey `json:"items"`
}

// Repository type metadata.
var (
	UserSSHKey_Kind             = "UserSSHKey"
	UserSSHKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserSSHKey_Kind}.String()
	UserSSHKey_KindAPIVersion   = UserSSHKey_Kind + "." + CRDGroupVersion.String()
	UserSSHKey_GroupVersionKind = CRDGroupVersion.WithKind(UserSSHKey_Kind)
)

func init() {
	SchemeBuilder.Register(&UserSSHKey{}, &UserSSHKeyList{})
}
