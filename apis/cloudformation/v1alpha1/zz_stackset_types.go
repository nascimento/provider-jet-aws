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

type AutoDeploymentObservation struct {
}

type AutoDeploymentParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RetainStacksOnAccountRemoval *bool `json:"retainStacksOnAccountRemoval,omitempty" tf:"retain_stacks_on_account_removal,omitempty"`
}

type StackSetObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	StackSetID *string `json:"stackSetId,omitempty" tf:"stack_set_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type StackSetParameters struct {

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	AdministrationRoleArn *string `json:"administrationRoleArn,omitempty" tf:"administration_role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	AdministrationRoleArnRef *v1.Reference `json:"administrationRoleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AdministrationRoleArnSelector *v1.Selector `json:"administrationRoleArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AutoDeployment []AutoDeploymentParameters `json:"autoDeployment,omitempty" tf:"auto_deployment,omitempty"`

	// +kubebuilder:validation:Optional
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ExecutionRoleName *string `json:"executionRoleName,omitempty" tf:"execution_role_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	PermissionModel *string `json:"permissionModel,omitempty" tf:"permission_model,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`
}

// StackSetSpec defines the desired state of StackSet
type StackSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StackSetParameters `json:"forProvider"`
}

// StackSetStatus defines the observed state of StackSet.
type StackSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StackSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StackSet is the Schema for the StackSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type StackSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StackSetSpec   `json:"spec"`
	Status            StackSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StackSetList contains a list of StackSets
type StackSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StackSet `json:"items"`
}

// Repository type metadata.
var (
	StackSet_Kind             = "StackSet"
	StackSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StackSet_Kind}.String()
	StackSet_KindAPIVersion   = StackSet_Kind + "." + CRDGroupVersion.String()
	StackSet_GroupVersionKind = CRDGroupVersion.WithKind(StackSet_Kind)
)

func init() {
	SchemeBuilder.Register(&StackSet{}, &StackSetList{})
}
