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

type LaunchSpecificationEBSBlockDeviceObservation struct {
}

type LaunchSpecificationEBSBlockDeviceParameters struct {

	// +kubebuilder:validation:Optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	Iops *int64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// +kubebuilder:validation:Optional
	Throughput *int64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeSize *int64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type LaunchSpecificationEphemeralBlockDeviceObservation struct {
}

type LaunchSpecificationEphemeralBlockDeviceParameters struct {

	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Required
	VirtualName *string `json:"virtualName" tf:"virtual_name,omitempty"`
}

type LaunchSpecificationObservation struct {
}

type LaunchSpecificationParameters struct {

	// +kubebuilder:validation:Required
	AMI *string `json:"ami" tf:"ami,omitempty"`

	// +kubebuilder:validation:Optional
	AssociatePublicIPAddress *bool `json:"associatePublicIpAddress,omitempty" tf:"associate_public_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Optional
	EBSBlockDevice []LaunchSpecificationEBSBlockDeviceParameters `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`

	// +kubebuilder:validation:Optional
	EBSOptimized *bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`

	// +kubebuilder:validation:Optional
	EphemeralBlockDevice []LaunchSpecificationEphemeralBlockDeviceParameters `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`

	// +kubebuilder:validation:Optional
	IAMInstanceProfile *string `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile,omitempty"`

	// +kubebuilder:validation:Optional
	IAMInstanceProfileArn *string `json:"iamInstanceProfileArn,omitempty" tf:"iam_instance_profile_arn,omitempty"`

	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// +kubebuilder:validation:Optional
	Monitoring *bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`

	// +kubebuilder:validation:Optional
	PlacementGroup *string `json:"placementGroup,omitempty" tf:"placement_group,omitempty"`

	// +kubebuilder:validation:Optional
	PlacementTenancy *string `json:"placementTenancy,omitempty" tf:"placement_tenancy,omitempty"`

	// +kubebuilder:validation:Optional
	RootBlockDevice []RootBlockDeviceParameters `json:"rootBlockDevice,omitempty" tf:"root_block_device,omitempty"`

	// +kubebuilder:validation:Optional
	SpotPrice *string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	WeightedCapacity *string `json:"weightedCapacity,omitempty" tf:"weighted_capacity,omitempty"`
}

type LaunchTemplateConfigLaunchTemplateSpecificationObservation struct {
}

type LaunchTemplateConfigLaunchTemplateSpecificationParameters struct {

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type OverridesObservation struct {
}

type OverridesParameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	SpotPrice *string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	WeightedCapacity *float64 `json:"weightedCapacity,omitempty" tf:"weighted_capacity,omitempty"`
}

type RootBlockDeviceObservation struct {
}

type RootBlockDeviceParameters struct {

	// +kubebuilder:validation:Optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	Iops *int64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	Throughput *int64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeSize *int64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type SpotFleetRequestLaunchTemplateConfigObservation struct {
}

type SpotFleetRequestLaunchTemplateConfigParameters struct {

	// +kubebuilder:validation:Required
	LaunchTemplateSpecification []LaunchTemplateConfigLaunchTemplateSpecificationParameters `json:"launchTemplateSpecification" tf:"launch_template_specification,omitempty"`

	// +kubebuilder:validation:Optional
	Overrides []OverridesParameters `json:"overrides,omitempty" tf:"overrides,omitempty"`
}

type SpotFleetRequestObservation struct {
	ClientToken *string `json:"clientToken,omitempty" tf:"client_token,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SpotRequestState *string `json:"spotRequestState,omitempty" tf:"spot_request_state,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SpotFleetRequestParameters struct {

	// +kubebuilder:validation:Optional
	AllocationStrategy *string `json:"allocationStrategy,omitempty" tf:"allocation_strategy,omitempty"`

	// +kubebuilder:validation:Optional
	ExcessCapacityTerminationPolicy *string `json:"excessCapacityTerminationPolicy,omitempty" tf:"excess_capacity_termination_policy,omitempty"`

	// +kubebuilder:validation:Optional
	FleetType *string `json:"fleetType,omitempty" tf:"fleet_type,omitempty"`

	// +kubebuilder:validation:Required
	IAMFleetRole *string `json:"iamFleetRole" tf:"iam_fleet_role,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceInterruptionBehaviour *string `json:"instanceInterruptionBehaviour,omitempty" tf:"instance_interruption_behaviour,omitempty"`

	// +kubebuilder:validation:Optional
	InstancePoolsToUseCount *int64 `json:"instancePoolsToUseCount,omitempty" tf:"instance_pools_to_use_count,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchSpecification []LaunchSpecificationParameters `json:"launchSpecification,omitempty" tf:"launch_specification,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplateConfig []SpotFleetRequestLaunchTemplateConfigParameters `json:"launchTemplateConfig,omitempty" tf:"launch_template_config,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancers []*string `json:"loadBalancers,omitempty" tf:"load_balancers,omitempty"`

	// +kubebuilder:validation:Optional
	OnDemandAllocationStrategy *string `json:"onDemandAllocationStrategy,omitempty" tf:"on_demand_allocation_strategy,omitempty"`

	// +kubebuilder:validation:Optional
	OnDemandMaxTotalPrice *string `json:"onDemandMaxTotalPrice,omitempty" tf:"on_demand_max_total_price,omitempty"`

	// +kubebuilder:validation:Optional
	OnDemandTargetCapacity *int64 `json:"onDemandTargetCapacity,omitempty" tf:"on_demand_target_capacity,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ReplaceUnhealthyInstances *bool `json:"replaceUnhealthyInstances,omitempty" tf:"replace_unhealthy_instances,omitempty"`

	// +kubebuilder:validation:Optional
	SpotMaintenanceStrategies []SpotMaintenanceStrategiesParameters `json:"spotMaintenanceStrategies,omitempty" tf:"spot_maintenance_strategies,omitempty"`

	// +kubebuilder:validation:Optional
	SpotPrice *string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TargetCapacity *int64 `json:"targetCapacity" tf:"target_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	TargetGroupArns []*string `json:"targetGroupArns,omitempty" tf:"target_group_arns,omitempty"`

	// +kubebuilder:validation:Optional
	TerminateInstancesWithExpiration *bool `json:"terminateInstancesWithExpiration,omitempty" tf:"terminate_instances_with_expiration,omitempty"`

	// +kubebuilder:validation:Optional
	ValidFrom *string `json:"validFrom,omitempty" tf:"valid_from,omitempty"`

	// +kubebuilder:validation:Optional
	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`

	// +kubebuilder:validation:Optional
	WaitForFulfillment *bool `json:"waitForFulfillment,omitempty" tf:"wait_for_fulfillment,omitempty"`
}

type SpotMaintenanceStrategiesCapacityRebalanceObservation struct {
}

type SpotMaintenanceStrategiesCapacityRebalanceParameters struct {

	// +kubebuilder:validation:Optional
	ReplacementStrategy *string `json:"replacementStrategy,omitempty" tf:"replacement_strategy,omitempty"`
}

type SpotMaintenanceStrategiesObservation struct {
}

type SpotMaintenanceStrategiesParameters struct {

	// +kubebuilder:validation:Optional
	CapacityRebalance []SpotMaintenanceStrategiesCapacityRebalanceParameters `json:"capacityRebalance,omitempty" tf:"capacity_rebalance,omitempty"`
}

// SpotFleetRequestSpec defines the desired state of SpotFleetRequest
type SpotFleetRequestSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpotFleetRequestParameters `json:"forProvider"`
}

// SpotFleetRequestStatus defines the observed state of SpotFleetRequest.
type SpotFleetRequestStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpotFleetRequestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpotFleetRequest is the Schema for the SpotFleetRequests API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type SpotFleetRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpotFleetRequestSpec   `json:"spec"`
	Status            SpotFleetRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpotFleetRequestList contains a list of SpotFleetRequests
type SpotFleetRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotFleetRequest `json:"items"`
}

// Repository type metadata.
var (
	SpotFleetRequest_Kind             = "SpotFleetRequest"
	SpotFleetRequest_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpotFleetRequest_Kind}.String()
	SpotFleetRequest_KindAPIVersion   = SpotFleetRequest_Kind + "." + CRDGroupVersion.String()
	SpotFleetRequest_GroupVersionKind = CRDGroupVersion.WithKind(SpotFleetRequest_Kind)
)

func init() {
	SchemeBuilder.Register(&SpotFleetRequest{}, &SpotFleetRequestList{})
}