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

type CapacityReservationSpecificationObservation struct {
}

type CapacityReservationSpecificationParameters struct {

	// +kubebuilder:validation:Optional
	CapacityReservationPreference *string `json:"capacityReservationPreference,omitempty" tf:"capacity_reservation_preference,omitempty"`

	// +kubebuilder:validation:Optional
	CapacityReservationTarget []CapacityReservationTargetParameters `json:"capacityReservationTarget,omitempty" tf:"capacity_reservation_target,omitempty"`
}

type CapacityReservationTargetObservation struct {
}

type CapacityReservationTargetParameters struct {

	// +kubebuilder:validation:Optional
	CapacityReservationID *string `json:"capacityReservationId,omitempty" tf:"capacity_reservation_id,omitempty"`
}

type CreditSpecificationObservation struct {
}

type CreditSpecificationParameters struct {

	// +kubebuilder:validation:Optional
	CPUCredits *string `json:"cpuCredits,omitempty" tf:"cpu_credits,omitempty"`
}

type EnclaveOptionsObservation struct {
}

type EnclaveOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type LaunchTemplateObservation struct {
}

type LaunchTemplateParameters struct {

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MetadataOptionsObservation struct {
}

type MetadataOptionsParameters struct {

	// +kubebuilder:validation:Optional
	HTTPEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPPutResponseHopLimit *float64 `json:"httpPutResponseHopLimit,omitempty" tf:"http_put_response_hop_limit,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPTokens *string `json:"httpTokens,omitempty" tf:"http_tokens,omitempty"`
}

type NetworkInterfaceObservation struct {
}

type NetworkInterfaceParameters struct {

	// +kubebuilder:validation:Optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// +kubebuilder:validation:Required
	DeviceIndex *float64 `json:"deviceIndex" tf:"device_index,omitempty"`

	// +kubebuilder:validation:Required
	NetworkInterfaceID *string `json:"networkInterfaceId" tf:"network_interface_id,omitempty"`
}

type SpotInstanceRequestEBSBlockDeviceObservation struct {
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

type SpotInstanceRequestEBSBlockDeviceParameters struct {

	// +kubebuilder:validation:Optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type SpotInstanceRequestEphemeralBlockDeviceObservation struct {
}

type SpotInstanceRequestEphemeralBlockDeviceParameters struct {

	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Optional
	NoDevice *bool `json:"noDevice,omitempty" tf:"no_device,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type SpotInstanceRequestObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceState *string `json:"instanceState,omitempty" tf:"instance_state,omitempty"`

	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	PasswordData *string `json:"passwordData,omitempty" tf:"password_data,omitempty"`

	PrimaryNetworkInterfaceID *string `json:"primaryNetworkInterfaceId,omitempty" tf:"primary_network_interface_id,omitempty"`

	PrivateDNS *string `json:"privateDns,omitempty" tf:"private_dns,omitempty"`

	PublicDNS *string `json:"publicDns,omitempty" tf:"public_dns,omitempty"`

	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	SpotBidStatus *string `json:"spotBidStatus,omitempty" tf:"spot_bid_status,omitempty"`

	SpotInstanceID *string `json:"spotInstanceId,omitempty" tf:"spot_instance_id,omitempty"`

	SpotRequestState *string `json:"spotRequestState,omitempty" tf:"spot_request_state,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SpotInstanceRequestParameters struct {

	// +kubebuilder:validation:Optional
	AMI *string `json:"ami,omitempty" tf:"ami,omitempty"`

	// +kubebuilder:validation:Optional
	AssociatePublicIPAddress *bool `json:"associatePublicIpAddress,omitempty" tf:"associate_public_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Optional
	BlockDurationMinutes *float64 `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes,omitempty"`

	// +kubebuilder:validation:Optional
	CPUCoreCount *float64 `json:"cpuCoreCount,omitempty" tf:"cpu_core_count,omitempty"`

	// +kubebuilder:validation:Optional
	CPUThreadsPerCore *float64 `json:"cpuThreadsPerCore,omitempty" tf:"cpu_threads_per_core,omitempty"`

	// +kubebuilder:validation:Optional
	CapacityReservationSpecification []CapacityReservationSpecificationParameters `json:"capacityReservationSpecification,omitempty" tf:"capacity_reservation_specification,omitempty"`

	// +kubebuilder:validation:Optional
	CreditSpecification []CreditSpecificationParameters `json:"creditSpecification,omitempty" tf:"credit_specification,omitempty"`

	// +kubebuilder:validation:Optional
	DisableAPITermination *bool `json:"disableApiTermination,omitempty" tf:"disable_api_termination,omitempty"`

	// +kubebuilder:validation:Optional
	EBSBlockDevice []SpotInstanceRequestEBSBlockDeviceParameters `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`

	// +kubebuilder:validation:Optional
	EBSOptimized *bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`

	// +kubebuilder:validation:Optional
	EnclaveOptions []EnclaveOptionsParameters `json:"enclaveOptions,omitempty" tf:"enclave_options,omitempty"`

	// +kubebuilder:validation:Optional
	EphemeralBlockDevice []SpotInstanceRequestEphemeralBlockDeviceParameters `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`

	// +kubebuilder:validation:Optional
	GetPasswordData *bool `json:"getPasswordData,omitempty" tf:"get_password_data,omitempty"`

	// +kubebuilder:validation:Optional
	Hibernation *bool `json:"hibernation,omitempty" tf:"hibernation,omitempty"`

	// +kubebuilder:validation:Optional
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// +kubebuilder:validation:Optional
	IAMInstanceProfile *string `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6AddressCount *float64 `json:"ipv6AddressCount,omitempty" tf:"ipv6_address_count,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6Addresses []*string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceInitiatedShutdownBehavior *string `json:"instanceInitiatedShutdownBehavior,omitempty" tf:"instance_initiated_shutdown_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceInterruptionBehavior *string `json:"instanceInterruptionBehavior,omitempty" tf:"instance_interruption_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceInterruptionBehaviour *string `json:"instanceInterruptionBehaviour,omitempty" tf:"instance_interruption_behaviour,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchGroup *string `json:"launchGroup,omitempty" tf:"launch_group,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplate []LaunchTemplateParameters `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`

	// +kubebuilder:validation:Optional
	MetadataOptions []MetadataOptionsParameters `json:"metadataOptions,omitempty" tf:"metadata_options,omitempty"`

	// +kubebuilder:validation:Optional
	Monitoring *bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterface []NetworkInterfaceParameters `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// +kubebuilder:validation:Optional
	PlacementGroup *string `json:"placementGroup,omitempty" tf:"placement_group,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RootBlockDevice []SpotInstanceRequestRootBlockDeviceParameters `json:"rootBlockDevice,omitempty" tf:"root_block_device,omitempty"`

	// +kubebuilder:validation:Optional
	SecondaryPrivateIps []*string `json:"secondaryPrivateIps,omitempty" tf:"secondary_private_ips,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// +kubebuilder:validation:Optional
	SourceDestCheck *bool `json:"sourceDestCheck,omitempty" tf:"source_dest_check,omitempty"`

	// +kubebuilder:validation:Optional
	SpotPrice *string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// +kubebuilder:validation:Optional
	SpotType *string `json:"spotType,omitempty" tf:"spot_type,omitempty"`

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
	Tenancy *string `json:"tenancy,omitempty" tf:"tenancy,omitempty"`

	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// +kubebuilder:validation:Optional
	UserDataBase64 *string `json:"userDataBase64,omitempty" tf:"user_data_base64,omitempty"`

	// +kubebuilder:validation:Optional
	VPCSecurityGroupIdRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCSecurityGroupIdSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIdRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIdSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	ValidFrom *string `json:"validFrom,omitempty" tf:"valid_from,omitempty"`

	// +kubebuilder:validation:Optional
	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeTags map[string]*string `json:"volumeTags,omitempty" tf:"volume_tags,omitempty"`

	// +kubebuilder:validation:Optional
	WaitForFulfillment *bool `json:"waitForFulfillment,omitempty" tf:"wait_for_fulfillment,omitempty"`
}

type SpotInstanceRequestRootBlockDeviceObservation struct {
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

type SpotInstanceRequestRootBlockDeviceParameters struct {

	// +kubebuilder:validation:Optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

// SpotInstanceRequestSpec defines the desired state of SpotInstanceRequest
type SpotInstanceRequestSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpotInstanceRequestParameters `json:"forProvider"`
}

// SpotInstanceRequestStatus defines the observed state of SpotInstanceRequest.
type SpotInstanceRequestStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpotInstanceRequestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpotInstanceRequest is the Schema for the SpotInstanceRequests API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type SpotInstanceRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpotInstanceRequestSpec   `json:"spec"`
	Status            SpotInstanceRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpotInstanceRequestList contains a list of SpotInstanceRequests
type SpotInstanceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotInstanceRequest `json:"items"`
}

// Repository type metadata.
var (
	SpotInstanceRequest_Kind             = "SpotInstanceRequest"
	SpotInstanceRequest_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpotInstanceRequest_Kind}.String()
	SpotInstanceRequest_KindAPIVersion   = SpotInstanceRequest_Kind + "." + CRDGroupVersion.String()
	SpotInstanceRequest_GroupVersionKind = CRDGroupVersion.WithKind(SpotInstanceRequest_Kind)
)

func init() {
	SchemeBuilder.Register(&SpotInstanceRequest{}, &SpotInstanceRequestList{})
}
