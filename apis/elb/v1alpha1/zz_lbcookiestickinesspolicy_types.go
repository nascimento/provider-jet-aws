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

type LBCookieStickinessPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LBCookieStickinessPolicyParameters struct {

	// +kubebuilder:validation:Optional
	CookieExpirationPeriod *int64 `json:"cookieExpirationPeriod,omitempty" tf:"cookie_expiration_period,omitempty"`

	// +kubebuilder:validation:Required
	LBPort *int64 `json:"lbPort" tf:"lb_port,omitempty"`

	// +kubebuilder:validation:Required
	LoadBalancer *string `json:"loadBalancer" tf:"load_balancer,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// LBCookieStickinessPolicySpec defines the desired state of LBCookieStickinessPolicy
type LBCookieStickinessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBCookieStickinessPolicyParameters `json:"forProvider"`
}

// LBCookieStickinessPolicyStatus defines the observed state of LBCookieStickinessPolicy.
type LBCookieStickinessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBCookieStickinessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBCookieStickinessPolicy is the Schema for the LBCookieStickinessPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type LBCookieStickinessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LBCookieStickinessPolicySpec   `json:"spec"`
	Status            LBCookieStickinessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBCookieStickinessPolicyList contains a list of LBCookieStickinessPolicys
type LBCookieStickinessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBCookieStickinessPolicy `json:"items"`
}

// Repository type metadata.
var (
	LBCookieStickinessPolicy_Kind             = "LBCookieStickinessPolicy"
	LBCookieStickinessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBCookieStickinessPolicy_Kind}.String()
	LBCookieStickinessPolicy_KindAPIVersion   = LBCookieStickinessPolicy_Kind + "." + CRDGroupVersion.String()
	LBCookieStickinessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(LBCookieStickinessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&LBCookieStickinessPolicy{}, &LBCookieStickinessPolicyList{})
}