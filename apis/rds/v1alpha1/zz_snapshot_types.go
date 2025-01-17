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

type SnapshotObservation struct {
	AllocatedStorage *float64 `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`

	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	DBSnapshotArn *string `json:"dbSnapshotArn,omitempty" tf:"db_snapshot_arn,omitempty"`

	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	LicenseModel *string `json:"licenseModel,omitempty" tf:"license_model,omitempty"`

	OptionGroupName *string `json:"optionGroupName,omitempty" tf:"option_group_name,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	SnapshotType *string `json:"snapshotType,omitempty" tf:"snapshot_type,omitempty"`

	SourceDBSnapshotIdentifier *string `json:"sourceDbSnapshotIdentifier,omitempty" tf:"source_db_snapshot_identifier,omitempty"`

	SourceRegion *string `json:"sourceRegion,omitempty" tf:"source_region,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type SnapshotParameters struct {

	// +kubebuilder:validation:Required
	DBInstanceIdentifier *string `json:"dbInstanceIdentifier" tf:"db_instance_identifier,omitempty"`

	// +kubebuilder:validation:Required
	DBSnapshotIdentifier *string `json:"dbSnapshotIdentifier" tf:"db_snapshot_identifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SnapshotSpec defines the desired state of Snapshot
type SnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotParameters `json:"forProvider"`
}

// SnapshotStatus defines the observed state of Snapshot.
type SnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Snapshot is the Schema for the Snapshots API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotList contains a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

// Repository type metadata.
var (
	Snapshot_Kind             = "Snapshot"
	Snapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Snapshot_Kind}.String()
	Snapshot_KindAPIVersion   = Snapshot_Kind + "." + CRDGroupVersion.String()
	Snapshot_GroupVersionKind = CRDGroupVersion.WithKind(Snapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}
