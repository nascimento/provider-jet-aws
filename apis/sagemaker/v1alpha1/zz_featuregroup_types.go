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

type DataCatalogConfigObservation struct {
}

type DataCatalogConfigParameters struct {

	// +kubebuilder:validation:Optional
	Catalog *string `json:"catalog,omitempty" tf:"catalog,omitempty"`

	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

type FeatureDefinitionObservation struct {
}

type FeatureDefinitionParameters struct {

	// +kubebuilder:validation:Optional
	FeatureName *string `json:"featureName,omitempty" tf:"feature_name,omitempty"`

	// +kubebuilder:validation:Optional
	FeatureType *string `json:"featureType,omitempty" tf:"feature_type,omitempty"`
}

type FeatureGroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type FeatureGroupParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	EventTimeFeatureName *string `json:"eventTimeFeatureName" tf:"event_time_feature_name,omitempty"`

	// +kubebuilder:validation:Required
	FeatureDefinition []FeatureDefinitionParameters `json:"featureDefinition" tf:"feature_definition,omitempty"`

	// +kubebuilder:validation:Required
	FeatureGroupName *string `json:"featureGroupName" tf:"feature_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	OfflineStoreConfig []OfflineStoreConfigParameters `json:"offlineStoreConfig,omitempty" tf:"offline_store_config,omitempty"`

	// +kubebuilder:validation:Optional
	OnlineStoreConfig []OnlineStoreConfigParameters `json:"onlineStoreConfig,omitempty" tf:"online_store_config,omitempty"`

	// +kubebuilder:validation:Required
	RecordIdentifierFeatureName *string `json:"recordIdentifierFeatureName" tf:"record_identifier_feature_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type OfflineStoreConfigObservation struct {
}

type OfflineStoreConfigParameters struct {

	// +kubebuilder:validation:Optional
	DataCatalogConfig []DataCatalogConfigParameters `json:"dataCatalogConfig,omitempty" tf:"data_catalog_config,omitempty"`

	// +kubebuilder:validation:Optional
	DisableGlueTableCreation *bool `json:"disableGlueTableCreation,omitempty" tf:"disable_glue_table_creation,omitempty"`

	// +kubebuilder:validation:Required
	S3StorageConfig []S3StorageConfigParameters `json:"s3StorageConfig" tf:"s3_storage_config,omitempty"`
}

type OnlineStoreConfigObservation struct {
}

type OnlineStoreConfigParameters struct {

	// +kubebuilder:validation:Optional
	EnableOnlineStore *bool `json:"enableOnlineStore,omitempty" tf:"enable_online_store,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityConfig []SecurityConfigParameters `json:"securityConfig,omitempty" tf:"security_config,omitempty"`
}

type S3StorageConfigObservation struct {
}

type S3StorageConfigParameters struct {

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Required
	S3URI *string `json:"s3Uri" tf:"s3_uri,omitempty"`
}

type SecurityConfigObservation struct {
}

type SecurityConfigParameters struct {

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

// FeatureGroupSpec defines the desired state of FeatureGroup
type FeatureGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FeatureGroupParameters `json:"forProvider"`
}

// FeatureGroupStatus defines the observed state of FeatureGroup.
type FeatureGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FeatureGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FeatureGroup is the Schema for the FeatureGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type FeatureGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FeatureGroupSpec   `json:"spec"`
	Status            FeatureGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FeatureGroupList contains a list of FeatureGroups
type FeatureGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FeatureGroup `json:"items"`
}

// Repository type metadata.
var (
	FeatureGroup_Kind             = "FeatureGroup"
	FeatureGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FeatureGroup_Kind}.String()
	FeatureGroup_KindAPIVersion   = FeatureGroup_Kind + "." + CRDGroupVersion.String()
	FeatureGroup_GroupVersionKind = CRDGroupVersion.WithKind(FeatureGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&FeatureGroup{}, &FeatureGroupList{})
}
