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

type InfrastructureConfigurationObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created,omitempty"`

	DateUpdated *string `json:"dateUpdated,omitempty" tf:"date_updated,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type InfrastructureConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	InstanceProfileName *string `json:"instanceProfileName" tf:"instance_profile_name,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceTypes []*string `json:"instanceTypes,omitempty" tf:"instance_types,omitempty"`

	// +kubebuilder:validation:Optional
	KeyPair *string `json:"keyPair,omitempty" tf:"key_pair,omitempty"`

	// +kubebuilder:validation:Optional
	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceTags map[string]*string `json:"resourceTags,omitempty" tf:"resource_tags,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIdRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupIdSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIdRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIdSelector
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TerminateInstanceOnFailure *bool `json:"terminateInstanceOnFailure,omitempty" tf:"terminate_instance_on_failure,omitempty"`
}

type LoggingObservation struct {
}

type LoggingParameters struct {

	// +kubebuilder:validation:Required
	S3Logs []S3LogsParameters `json:"s3Logs" tf:"s3_logs,omitempty"`
}

type S3LogsObservation struct {
}

type S3LogsParameters struct {

	// +kubebuilder:validation:Required
	S3BucketName *string `json:"s3BucketName" tf:"s3_bucket_name,omitempty"`

	// +kubebuilder:validation:Optional
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

// InfrastructureConfigurationSpec defines the desired state of InfrastructureConfiguration
type InfrastructureConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InfrastructureConfigurationParameters `json:"forProvider"`
}

// InfrastructureConfigurationStatus defines the observed state of InfrastructureConfiguration.
type InfrastructureConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InfrastructureConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InfrastructureConfiguration is the Schema for the InfrastructureConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type InfrastructureConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InfrastructureConfigurationSpec   `json:"spec"`
	Status            InfrastructureConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InfrastructureConfigurationList contains a list of InfrastructureConfigurations
type InfrastructureConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InfrastructureConfiguration `json:"items"`
}

// Repository type metadata.
var (
	InfrastructureConfiguration_Kind             = "InfrastructureConfiguration"
	InfrastructureConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InfrastructureConfiguration_Kind}.String()
	InfrastructureConfiguration_KindAPIVersion   = InfrastructureConfiguration_Kind + "." + CRDGroupVersion.String()
	InfrastructureConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(InfrastructureConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&InfrastructureConfiguration{}, &InfrastructureConfigurationList{})
}
