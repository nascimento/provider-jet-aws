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

type ExternalKeyObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ExpirationModel *string `json:"expirationModel,omitempty" tf:"expiration_model,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KeyState *string `json:"keyState,omitempty" tf:"key_state,omitempty"`

	KeyUsage *string `json:"keyUsage,omitempty" tf:"key_usage,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ExternalKeyParameters struct {

	// +kubebuilder:validation:Optional
	BypassPolicyLockoutSafetyCheck *bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" tf:"bypass_policy_lockout_safety_check,omitempty"`

	// +kubebuilder:validation:Optional
	DeletionWindowInDays *float64 `json:"deletionWindowInDays,omitempty" tf:"deletion_window_in_days,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	KeyMaterialBase64SecretRef *v1.SecretKeySelector `json:"keyMaterialBase64SecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	ValidTo *string `json:"validTo,omitempty" tf:"valid_to,omitempty"`
}

// ExternalKeySpec defines the desired state of ExternalKey
type ExternalKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExternalKeyParameters `json:"forProvider"`
}

// ExternalKeyStatus defines the observed state of ExternalKey.
type ExternalKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExternalKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalKey is the Schema for the ExternalKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ExternalKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExternalKeySpec   `json:"spec"`
	Status            ExternalKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalKeyList contains a list of ExternalKeys
type ExternalKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalKey `json:"items"`
}

// Repository type metadata.
var (
	ExternalKey_Kind             = "ExternalKey"
	ExternalKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExternalKey_Kind}.String()
	ExternalKey_KindAPIVersion   = ExternalKey_Kind + "." + CRDGroupVersion.String()
	ExternalKey_GroupVersionKind = CRDGroupVersion.WithKind(ExternalKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ExternalKey{}, &ExternalKeyList{})
}
