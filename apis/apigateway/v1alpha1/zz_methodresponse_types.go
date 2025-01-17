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

type MethodResponseObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MethodResponseParameters struct {

	// +kubebuilder:validation:Required
	HTTPMethod *string `json:"httpMethod" tf:"http_method,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceId" tf:"resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	ResponseModels map[string]*string `json:"responseModels,omitempty" tf:"response_models,omitempty"`

	// +kubebuilder:validation:Optional
	ResponseParameters map[string]*bool `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`

	// +kubebuilder:validation:Required
	RestAPIID *string `json:"restApiId" tf:"rest_api_id,omitempty"`

	// +kubebuilder:validation:Required
	StatusCode *string `json:"statusCode" tf:"status_code,omitempty"`
}

// MethodResponseSpec defines the desired state of MethodResponse
type MethodResponseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MethodResponseParameters `json:"forProvider"`
}

// MethodResponseStatus defines the observed state of MethodResponse.
type MethodResponseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MethodResponseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MethodResponse is the Schema for the MethodResponses API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type MethodResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MethodResponseSpec   `json:"spec"`
	Status            MethodResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MethodResponseList contains a list of MethodResponses
type MethodResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MethodResponse `json:"items"`
}

// Repository type metadata.
var (
	MethodResponse_Kind             = "MethodResponse"
	MethodResponse_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MethodResponse_Kind}.String()
	MethodResponse_KindAPIVersion   = MethodResponse_Kind + "." + CRDGroupVersion.String()
	MethodResponse_GroupVersionKind = CRDGroupVersion.WithKind(MethodResponse_Kind)
)

func init() {
	SchemeBuilder.Register(&MethodResponse{}, &MethodResponseList{})
}
