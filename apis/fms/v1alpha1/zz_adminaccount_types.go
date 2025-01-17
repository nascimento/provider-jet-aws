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

type AdminAccountObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AdminAccountParameters struct {

	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// AdminAccountSpec defines the desired state of AdminAccount
type AdminAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AdminAccountParameters `json:"forProvider"`
}

// AdminAccountStatus defines the observed state of AdminAccount.
type AdminAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AdminAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AdminAccount is the Schema for the AdminAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type AdminAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AdminAccountSpec   `json:"spec"`
	Status            AdminAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AdminAccountList contains a list of AdminAccounts
type AdminAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdminAccount `json:"items"`
}

// Repository type metadata.
var (
	AdminAccount_Kind             = "AdminAccount"
	AdminAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AdminAccount_Kind}.String()
	AdminAccount_KindAPIVersion   = AdminAccount_Kind + "." + CRDGroupVersion.String()
	AdminAccount_GroupVersionKind = CRDGroupVersion.WithKind(AdminAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&AdminAccount{}, &AdminAccountList{})
}
