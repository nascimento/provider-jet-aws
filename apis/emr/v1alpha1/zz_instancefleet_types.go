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

type InstanceFleetInstanceTypeConfigsConfigurationsObservation struct {
}

type InstanceFleetInstanceTypeConfigsConfigurationsParameters struct {

	// +kubebuilder:validation:Optional
	Classification *string `json:"classification,omitempty" tf:"classification,omitempty"`

	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`
}

type InstanceFleetInstanceTypeConfigsEBSConfigObservation struct {
}

type InstanceFleetInstanceTypeConfigsEBSConfigParameters struct {

	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Required
	Size *float64 `json:"size" tf:"size,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	VolumesPerInstance *float64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance,omitempty"`
}

type InstanceFleetInstanceTypeConfigsObservation struct {
}

type InstanceFleetInstanceTypeConfigsParameters struct {

	// +kubebuilder:validation:Optional
	BidPrice *string `json:"bidPrice,omitempty" tf:"bid_price,omitempty"`

	// +kubebuilder:validation:Optional
	BidPriceAsPercentageOfOnDemandPrice *float64 `json:"bidPriceAsPercentageOfOnDemandPrice,omitempty" tf:"bid_price_as_percentage_of_on_demand_price,omitempty"`

	// +kubebuilder:validation:Optional
	Configurations []InstanceFleetInstanceTypeConfigsConfigurationsParameters `json:"configurations,omitempty" tf:"configurations,omitempty"`

	// +kubebuilder:validation:Optional
	EBSConfig []InstanceFleetInstanceTypeConfigsEBSConfigParameters `json:"ebsConfig,omitempty" tf:"ebs_config,omitempty"`

	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	WeightedCapacity *float64 `json:"weightedCapacity,omitempty" tf:"weighted_capacity,omitempty"`
}

type InstanceFleetLaunchSpecificationsObservation struct {
}

type InstanceFleetLaunchSpecificationsOnDemandSpecificationObservation struct {
}

type InstanceFleetLaunchSpecificationsOnDemandSpecificationParameters struct {

	// +kubebuilder:validation:Required
	AllocationStrategy *string `json:"allocationStrategy" tf:"allocation_strategy,omitempty"`
}

type InstanceFleetLaunchSpecificationsParameters struct {

	// +kubebuilder:validation:Optional
	OnDemandSpecification []InstanceFleetLaunchSpecificationsOnDemandSpecificationParameters `json:"onDemandSpecification,omitempty" tf:"on_demand_specification,omitempty"`

	// +kubebuilder:validation:Optional
	SpotSpecification []InstanceFleetLaunchSpecificationsSpotSpecificationParameters `json:"spotSpecification,omitempty" tf:"spot_specification,omitempty"`
}

type InstanceFleetLaunchSpecificationsSpotSpecificationObservation struct {
}

type InstanceFleetLaunchSpecificationsSpotSpecificationParameters struct {

	// +kubebuilder:validation:Required
	AllocationStrategy *string `json:"allocationStrategy" tf:"allocation_strategy,omitempty"`

	// +kubebuilder:validation:Optional
	BlockDurationMinutes *float64 `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes,omitempty"`

	// +kubebuilder:validation:Required
	TimeoutAction *string `json:"timeoutAction" tf:"timeout_action,omitempty"`

	// +kubebuilder:validation:Required
	TimeoutDurationMinutes *float64 `json:"timeoutDurationMinutes" tf:"timeout_duration_minutes,omitempty"`
}

type InstanceFleetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ProvisionedOnDemandCapacity *float64 `json:"provisionedOnDemandCapacity,omitempty" tf:"provisioned_on_demand_capacity,omitempty"`

	ProvisionedSpotCapacity *float64 `json:"provisionedSpotCapacity,omitempty" tf:"provisioned_spot_capacity,omitempty"`
}

type InstanceFleetParameters struct {

	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceTypeConfigs []InstanceFleetInstanceTypeConfigsParameters `json:"instanceTypeConfigs,omitempty" tf:"instance_type_configs,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchSpecifications []InstanceFleetLaunchSpecificationsParameters `json:"launchSpecifications,omitempty" tf:"launch_specifications,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	TargetOnDemandCapacity *float64 `json:"targetOnDemandCapacity,omitempty" tf:"target_on_demand_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	TargetSpotCapacity *float64 `json:"targetSpotCapacity,omitempty" tf:"target_spot_capacity,omitempty"`
}

// InstanceFleetSpec defines the desired state of InstanceFleet
type InstanceFleetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceFleetParameters `json:"forProvider"`
}

// InstanceFleetStatus defines the observed state of InstanceFleet.
type InstanceFleetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceFleetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceFleet is the Schema for the InstanceFleets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type InstanceFleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceFleetSpec   `json:"spec"`
	Status            InstanceFleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceFleetList contains a list of InstanceFleets
type InstanceFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceFleet `json:"items"`
}

// Repository type metadata.
var (
	InstanceFleet_Kind             = "InstanceFleet"
	InstanceFleet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceFleet_Kind}.String()
	InstanceFleet_KindAPIVersion   = InstanceFleet_Kind + "." + CRDGroupVersion.String()
	InstanceFleet_GroupVersionKind = CRDGroupVersion.WithKind(InstanceFleet_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceFleet{}, &InstanceFleetList{})
}
