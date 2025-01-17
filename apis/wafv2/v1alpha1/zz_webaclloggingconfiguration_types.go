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

type ActionConditionObservation struct {
}

type ActionConditionParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`
}

type ConditionObservation struct {
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	ActionCondition []ActionConditionParameters `json:"actionCondition,omitempty" tf:"action_condition,omitempty"`

	// +kubebuilder:validation:Optional
	LabelNameCondition []LabelNameConditionParameters `json:"labelNameCondition,omitempty" tf:"label_name_condition,omitempty"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// +kubebuilder:validation:Required
	Behavior *string `json:"behavior" tf:"behavior,omitempty"`

	// +kubebuilder:validation:Required
	Condition []ConditionParameters `json:"condition" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Requirement *string `json:"requirement" tf:"requirement,omitempty"`
}

type LabelNameConditionObservation struct {
}

type LabelNameConditionParameters struct {

	// +kubebuilder:validation:Required
	LabelName *string `json:"labelName" tf:"label_name,omitempty"`
}

type LoggingFilterObservation struct {
}

type LoggingFilterParameters struct {

	// +kubebuilder:validation:Required
	DefaultBehavior *string `json:"defaultBehavior" tf:"default_behavior,omitempty"`

	// +kubebuilder:validation:Required
	Filter []FilterParameters `json:"filter" tf:"filter,omitempty"`
}

type RedactedFieldsAllQueryArgumentsObservation struct {
}

type RedactedFieldsAllQueryArgumentsParameters struct {
}

type RedactedFieldsBodyObservation struct {
}

type RedactedFieldsBodyParameters struct {
}

type RedactedFieldsMethodObservation struct {
}

type RedactedFieldsMethodParameters struct {
}

type RedactedFieldsObservation struct {
}

type RedactedFieldsParameters struct {

	// +kubebuilder:validation:Optional
	AllQueryArguments []RedactedFieldsAllQueryArgumentsParameters `json:"allQueryArguments,omitempty" tf:"all_query_arguments,omitempty"`

	// +kubebuilder:validation:Optional
	Body []RedactedFieldsBodyParameters `json:"body,omitempty" tf:"body,omitempty"`

	// +kubebuilder:validation:Optional
	Method []RedactedFieldsMethodParameters `json:"method,omitempty" tf:"method,omitempty"`

	// +kubebuilder:validation:Optional
	QueryString []RedactedFieldsQueryStringParameters `json:"queryString,omitempty" tf:"query_string,omitempty"`

	// +kubebuilder:validation:Optional
	SingleHeader []RedactedFieldsSingleHeaderParameters `json:"singleHeader,omitempty" tf:"single_header,omitempty"`

	// +kubebuilder:validation:Optional
	SingleQueryArgument []RedactedFieldsSingleQueryArgumentParameters `json:"singleQueryArgument,omitempty" tf:"single_query_argument,omitempty"`

	// +kubebuilder:validation:Optional
	URIPath []RedactedFieldsURIPathParameters `json:"uriPath,omitempty" tf:"uri_path,omitempty"`
}

type RedactedFieldsQueryStringObservation struct {
}

type RedactedFieldsQueryStringParameters struct {
}

type RedactedFieldsSingleHeaderObservation struct {
}

type RedactedFieldsSingleHeaderParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type RedactedFieldsSingleQueryArgumentObservation struct {
}

type RedactedFieldsSingleQueryArgumentParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type RedactedFieldsURIPathObservation struct {
}

type RedactedFieldsURIPathParameters struct {
}

type WebACLLoggingConfigurationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type WebACLLoggingConfigurationParameters struct {

	// AWS Kinesis Firehose Delivery Stream ARNs
	// +kubebuilder:validation:Required
	LogDestinationConfigs []*string `json:"logDestinationConfigs" tf:"log_destination_configs,omitempty"`

	// +kubebuilder:validation:Optional
	LoggingFilter []LoggingFilterParameters `json:"loggingFilter,omitempty" tf:"logging_filter,omitempty"`

	// Parts of the request to exclude from logs
	// +kubebuilder:validation:Optional
	RedactedFields []RedactedFieldsParameters `json:"redactedFields,omitempty" tf:"redacted_fields,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// AWS WebACL ARN
	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`
}

// WebACLLoggingConfigurationSpec defines the desired state of WebACLLoggingConfiguration
type WebACLLoggingConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebACLLoggingConfigurationParameters `json:"forProvider"`
}

// WebACLLoggingConfigurationStatus defines the observed state of WebACLLoggingConfiguration.
type WebACLLoggingConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebACLLoggingConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WebACLLoggingConfiguration is the Schema for the WebACLLoggingConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type WebACLLoggingConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebACLLoggingConfigurationSpec   `json:"spec"`
	Status            WebACLLoggingConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebACLLoggingConfigurationList contains a list of WebACLLoggingConfigurations
type WebACLLoggingConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebACLLoggingConfiguration `json:"items"`
}

// Repository type metadata.
var (
	WebACLLoggingConfiguration_Kind             = "WebACLLoggingConfiguration"
	WebACLLoggingConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebACLLoggingConfiguration_Kind}.String()
	WebACLLoggingConfiguration_KindAPIVersion   = WebACLLoggingConfiguration_Kind + "." + CRDGroupVersion.String()
	WebACLLoggingConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(WebACLLoggingConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&WebACLLoggingConfiguration{}, &WebACLLoggingConfigurationList{})
}
