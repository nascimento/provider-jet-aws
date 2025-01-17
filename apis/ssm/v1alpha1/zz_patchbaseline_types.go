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

type ApprovalRuleObservation struct {
}

type ApprovalRuleParameters struct {

	// +kubebuilder:validation:Optional
	ApproveAfterDays *float64 `json:"approveAfterDays,omitempty" tf:"approve_after_days,omitempty"`

	// +kubebuilder:validation:Optional
	ApproveUntilDate *string `json:"approveUntilDate,omitempty" tf:"approve_until_date,omitempty"`

	// +kubebuilder:validation:Optional
	ComplianceLevel *string `json:"complianceLevel,omitempty" tf:"compliance_level,omitempty"`

	// +kubebuilder:validation:Optional
	EnableNonSecurity *bool `json:"enableNonSecurity,omitempty" tf:"enable_non_security,omitempty"`

	// +kubebuilder:validation:Required
	PatchFilter []PatchFilterParameters `json:"patchFilter" tf:"patch_filter,omitempty"`
}

type GlobalFilterObservation struct {
}

type GlobalFilterParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type PatchBaselineObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type PatchBaselineParameters struct {

	// +kubebuilder:validation:Optional
	ApprovalRule []ApprovalRuleParameters `json:"approvalRule,omitempty" tf:"approval_rule,omitempty"`

	// +kubebuilder:validation:Optional
	ApprovedPatches []*string `json:"approvedPatches,omitempty" tf:"approved_patches,omitempty"`

	// +kubebuilder:validation:Optional
	ApprovedPatchesComplianceLevel *string `json:"approvedPatchesComplianceLevel,omitempty" tf:"approved_patches_compliance_level,omitempty"`

	// +kubebuilder:validation:Optional
	ApprovedPatchesEnableNonSecurity *bool `json:"approvedPatchesEnableNonSecurity,omitempty" tf:"approved_patches_enable_non_security,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	GlobalFilter []GlobalFilterParameters `json:"globalFilter,omitempty" tf:"global_filter,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RejectedPatches []*string `json:"rejectedPatches,omitempty" tf:"rejected_patches,omitempty"`

	// +kubebuilder:validation:Optional
	RejectedPatchesAction *string `json:"rejectedPatchesAction,omitempty" tf:"rejected_patches_action,omitempty"`

	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PatchFilterObservation struct {
}

type PatchFilterParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type SourceObservation struct {
}

type SourceParameters struct {

	// +kubebuilder:validation:Required
	Configuration *string `json:"configuration" tf:"configuration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Products []*string `json:"products" tf:"products,omitempty"`
}

// PatchBaselineSpec defines the desired state of PatchBaseline
type PatchBaselineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PatchBaselineParameters `json:"forProvider"`
}

// PatchBaselineStatus defines the observed state of PatchBaseline.
type PatchBaselineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PatchBaselineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PatchBaseline is the Schema for the PatchBaselines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type PatchBaseline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PatchBaselineSpec   `json:"spec"`
	Status            PatchBaselineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PatchBaselineList contains a list of PatchBaselines
type PatchBaselineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PatchBaseline `json:"items"`
}

// Repository type metadata.
var (
	PatchBaseline_Kind             = "PatchBaseline"
	PatchBaseline_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PatchBaseline_Kind}.String()
	PatchBaseline_KindAPIVersion   = PatchBaseline_Kind + "." + CRDGroupVersion.String()
	PatchBaseline_GroupVersionKind = CRDGroupVersion.WithKind(PatchBaseline_Kind)
)

func init() {
	SchemeBuilder.Register(&PatchBaseline{}, &PatchBaselineList{})
}
