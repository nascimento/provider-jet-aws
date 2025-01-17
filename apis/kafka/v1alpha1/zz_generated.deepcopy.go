//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerLogsObservation) DeepCopyInto(out *BrokerLogsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerLogsObservation.
func (in *BrokerLogsObservation) DeepCopy() *BrokerLogsObservation {
	if in == nil {
		return nil
	}
	out := new(BrokerLogsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerLogsParameters) DeepCopyInto(out *BrokerLogsParameters) {
	*out = *in
	if in.CloudwatchLogs != nil {
		in, out := &in.CloudwatchLogs, &out.CloudwatchLogs
		*out = make([]CloudwatchLogsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Firehose != nil {
		in, out := &in.Firehose, &out.Firehose
		*out = make([]FirehoseParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = make([]S3Parameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerLogsParameters.
func (in *BrokerLogsParameters) DeepCopy() *BrokerLogsParameters {
	if in == nil {
		return nil
	}
	out := new(BrokerLogsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerNodeGroupInfoObservation) DeepCopyInto(out *BrokerNodeGroupInfoObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerNodeGroupInfoObservation.
func (in *BrokerNodeGroupInfoObservation) DeepCopy() *BrokerNodeGroupInfoObservation {
	if in == nil {
		return nil
	}
	out := new(BrokerNodeGroupInfoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerNodeGroupInfoParameters) DeepCopyInto(out *BrokerNodeGroupInfoParameters) {
	*out = *in
	if in.AzDistribution != nil {
		in, out := &in.AzDistribution, &out.AzDistribution
		*out = new(string)
		**out = **in
	}
	if in.ClientSubnets != nil {
		in, out := &in.ClientSubnets, &out.ClientSubnets
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EBSVolumeSize != nil {
		in, out := &in.EBSVolumeSize, &out.EBSVolumeSize
		*out = new(float64)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerNodeGroupInfoParameters.
func (in *BrokerNodeGroupInfoParameters) DeepCopy() *BrokerNodeGroupInfoParameters {
	if in == nil {
		return nil
	}
	out := new(BrokerNodeGroupInfoParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientAuthenticationObservation) DeepCopyInto(out *ClientAuthenticationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientAuthenticationObservation.
func (in *ClientAuthenticationObservation) DeepCopy() *ClientAuthenticationObservation {
	if in == nil {
		return nil
	}
	out := new(ClientAuthenticationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientAuthenticationParameters) DeepCopyInto(out *ClientAuthenticationParameters) {
	*out = *in
	if in.Sasl != nil {
		in, out := &in.Sasl, &out.Sasl
		*out = make([]SaslParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]TLSParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientAuthenticationParameters.
func (in *ClientAuthenticationParameters) DeepCopy() *ClientAuthenticationParameters {
	if in == nil {
		return nil
	}
	out := new(ClientAuthenticationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchLogsObservation) DeepCopyInto(out *CloudwatchLogsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogsObservation.
func (in *CloudwatchLogsObservation) DeepCopy() *CloudwatchLogsObservation {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchLogsParameters) DeepCopyInto(out *CloudwatchLogsParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LogGroup != nil {
		in, out := &in.LogGroup, &out.LogGroup
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogsParameters.
func (in *CloudwatchLogsParameters) DeepCopy() *CloudwatchLogsParameters {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObservation) DeepCopyInto(out *ClusterObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.BootstrapBrokers != nil {
		in, out := &in.BootstrapBrokers, &out.BootstrapBrokers
		*out = new(string)
		**out = **in
	}
	if in.BootstrapBrokersSaslIAM != nil {
		in, out := &in.BootstrapBrokersSaslIAM, &out.BootstrapBrokersSaslIAM
		*out = new(string)
		**out = **in
	}
	if in.BootstrapBrokersSaslScram != nil {
		in, out := &in.BootstrapBrokersSaslScram, &out.BootstrapBrokersSaslScram
		*out = new(string)
		**out = **in
	}
	if in.BootstrapBrokersTLS != nil {
		in, out := &in.BootstrapBrokersTLS, &out.BootstrapBrokersTLS
		*out = new(string)
		**out = **in
	}
	if in.CurrentVersion != nil {
		in, out := &in.CurrentVersion, &out.CurrentVersion
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ZookeeperConnectString != nil {
		in, out := &in.ZookeeperConnectString, &out.ZookeeperConnectString
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObservation.
func (in *ClusterObservation) DeepCopy() *ClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameters) DeepCopyInto(out *ClusterParameters) {
	*out = *in
	if in.BrokerNodeGroupInfo != nil {
		in, out := &in.BrokerNodeGroupInfo, &out.BrokerNodeGroupInfo
		*out = make([]BrokerNodeGroupInfoParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClientAuthentication != nil {
		in, out := &in.ClientAuthentication, &out.ClientAuthentication
		*out = make([]ClientAuthenticationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.ConfigurationInfo != nil {
		in, out := &in.ConfigurationInfo, &out.ConfigurationInfo
		*out = make([]ConfigurationInfoParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EncryptionInfo != nil {
		in, out := &in.EncryptionInfo, &out.EncryptionInfo
		*out = make([]EncryptionInfoParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnhancedMonitoring != nil {
		in, out := &in.EnhancedMonitoring, &out.EnhancedMonitoring
		*out = new(string)
		**out = **in
	}
	if in.KafkaVersion != nil {
		in, out := &in.KafkaVersion, &out.KafkaVersion
		*out = new(string)
		**out = **in
	}
	if in.LoggingInfo != nil {
		in, out := &in.LoggingInfo, &out.LoggingInfo
		*out = make([]LoggingInfoParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NumberOfBrokerNodes != nil {
		in, out := &in.NumberOfBrokerNodes, &out.NumberOfBrokerNodes
		*out = new(float64)
		**out = **in
	}
	if in.OpenMonitoring != nil {
		in, out := &in.OpenMonitoring, &out.OpenMonitoring
		*out = make([]OpenMonitoringParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameters.
func (in *ClusterParameters) DeepCopy() *ClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationInfoObservation) DeepCopyInto(out *ConfigurationInfoObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationInfoObservation.
func (in *ConfigurationInfoObservation) DeepCopy() *ConfigurationInfoObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigurationInfoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationInfoParameters) DeepCopyInto(out *ConfigurationInfoParameters) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationInfoParameters.
func (in *ConfigurationInfoParameters) DeepCopy() *ConfigurationInfoParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigurationInfoParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationList) DeepCopyInto(out *ConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Configuration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationList.
func (in *ConfigurationList) DeepCopy() *ConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationObservation) DeepCopyInto(out *ConfigurationObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LatestRevision != nil {
		in, out := &in.LatestRevision, &out.LatestRevision
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationObservation.
func (in *ConfigurationObservation) DeepCopy() *ConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationParameters) DeepCopyInto(out *ConfigurationParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.KafkaVersions != nil {
		in, out := &in.KafkaVersions, &out.KafkaVersions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ServerProperties != nil {
		in, out := &in.ServerProperties, &out.ServerProperties
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationParameters.
func (in *ConfigurationParameters) DeepCopy() *ConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpec) DeepCopyInto(out *ConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpec.
func (in *ConfigurationSpec) DeepCopy() *ConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStatus) DeepCopyInto(out *ConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStatus.
func (in *ConfigurationStatus) DeepCopy() *ConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionInTransitObservation) DeepCopyInto(out *EncryptionInTransitObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionInTransitObservation.
func (in *EncryptionInTransitObservation) DeepCopy() *EncryptionInTransitObservation {
	if in == nil {
		return nil
	}
	out := new(EncryptionInTransitObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionInTransitParameters) DeepCopyInto(out *EncryptionInTransitParameters) {
	*out = *in
	if in.ClientBroker != nil {
		in, out := &in.ClientBroker, &out.ClientBroker
		*out = new(string)
		**out = **in
	}
	if in.InCluster != nil {
		in, out := &in.InCluster, &out.InCluster
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionInTransitParameters.
func (in *EncryptionInTransitParameters) DeepCopy() *EncryptionInTransitParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptionInTransitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionInfoObservation) DeepCopyInto(out *EncryptionInfoObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionInfoObservation.
func (in *EncryptionInfoObservation) DeepCopy() *EncryptionInfoObservation {
	if in == nil {
		return nil
	}
	out := new(EncryptionInfoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionInfoParameters) DeepCopyInto(out *EncryptionInfoParameters) {
	*out = *in
	if in.EncryptionAtRestKMSKeyArn != nil {
		in, out := &in.EncryptionAtRestKMSKeyArn, &out.EncryptionAtRestKMSKeyArn
		*out = new(string)
		**out = **in
	}
	if in.EncryptionInTransit != nil {
		in, out := &in.EncryptionInTransit, &out.EncryptionInTransit
		*out = make([]EncryptionInTransitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionInfoParameters.
func (in *EncryptionInfoParameters) DeepCopy() *EncryptionInfoParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptionInfoParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirehoseObservation) DeepCopyInto(out *FirehoseObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirehoseObservation.
func (in *FirehoseObservation) DeepCopy() *FirehoseObservation {
	if in == nil {
		return nil
	}
	out := new(FirehoseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirehoseParameters) DeepCopyInto(out *FirehoseParameters) {
	*out = *in
	if in.DeliveryStream != nil {
		in, out := &in.DeliveryStream, &out.DeliveryStream
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirehoseParameters.
func (in *FirehoseParameters) DeepCopy() *FirehoseParameters {
	if in == nil {
		return nil
	}
	out := new(FirehoseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JmxExporterObservation) DeepCopyInto(out *JmxExporterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JmxExporterObservation.
func (in *JmxExporterObservation) DeepCopy() *JmxExporterObservation {
	if in == nil {
		return nil
	}
	out := new(JmxExporterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JmxExporterParameters) DeepCopyInto(out *JmxExporterParameters) {
	*out = *in
	if in.EnabledInBroker != nil {
		in, out := &in.EnabledInBroker, &out.EnabledInBroker
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JmxExporterParameters.
func (in *JmxExporterParameters) DeepCopy() *JmxExporterParameters {
	if in == nil {
		return nil
	}
	out := new(JmxExporterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingInfoObservation) DeepCopyInto(out *LoggingInfoObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingInfoObservation.
func (in *LoggingInfoObservation) DeepCopy() *LoggingInfoObservation {
	if in == nil {
		return nil
	}
	out := new(LoggingInfoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingInfoParameters) DeepCopyInto(out *LoggingInfoParameters) {
	*out = *in
	if in.BrokerLogs != nil {
		in, out := &in.BrokerLogs, &out.BrokerLogs
		*out = make([]BrokerLogsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingInfoParameters.
func (in *LoggingInfoParameters) DeepCopy() *LoggingInfoParameters {
	if in == nil {
		return nil
	}
	out := new(LoggingInfoParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeExporterObservation) DeepCopyInto(out *NodeExporterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeExporterObservation.
func (in *NodeExporterObservation) DeepCopy() *NodeExporterObservation {
	if in == nil {
		return nil
	}
	out := new(NodeExporterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeExporterParameters) DeepCopyInto(out *NodeExporterParameters) {
	*out = *in
	if in.EnabledInBroker != nil {
		in, out := &in.EnabledInBroker, &out.EnabledInBroker
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeExporterParameters.
func (in *NodeExporterParameters) DeepCopy() *NodeExporterParameters {
	if in == nil {
		return nil
	}
	out := new(NodeExporterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenMonitoringObservation) DeepCopyInto(out *OpenMonitoringObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenMonitoringObservation.
func (in *OpenMonitoringObservation) DeepCopy() *OpenMonitoringObservation {
	if in == nil {
		return nil
	}
	out := new(OpenMonitoringObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenMonitoringParameters) DeepCopyInto(out *OpenMonitoringParameters) {
	*out = *in
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = make([]PrometheusParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenMonitoringParameters.
func (in *OpenMonitoringParameters) DeepCopy() *OpenMonitoringParameters {
	if in == nil {
		return nil
	}
	out := new(OpenMonitoringParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusObservation) DeepCopyInto(out *PrometheusObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusObservation.
func (in *PrometheusObservation) DeepCopy() *PrometheusObservation {
	if in == nil {
		return nil
	}
	out := new(PrometheusObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusParameters) DeepCopyInto(out *PrometheusParameters) {
	*out = *in
	if in.JmxExporter != nil {
		in, out := &in.JmxExporter, &out.JmxExporter
		*out = make([]JmxExporterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeExporter != nil {
		in, out := &in.NodeExporter, &out.NodeExporter
		*out = make([]NodeExporterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusParameters.
func (in *PrometheusParameters) DeepCopy() *PrometheusParameters {
	if in == nil {
		return nil
	}
	out := new(PrometheusParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Observation) DeepCopyInto(out *S3Observation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Observation.
func (in *S3Observation) DeepCopy() *S3Observation {
	if in == nil {
		return nil
	}
	out := new(S3Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Parameters) DeepCopyInto(out *S3Parameters) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Parameters.
func (in *S3Parameters) DeepCopy() *S3Parameters {
	if in == nil {
		return nil
	}
	out := new(S3Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SaslObservation) DeepCopyInto(out *SaslObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SaslObservation.
func (in *SaslObservation) DeepCopy() *SaslObservation {
	if in == nil {
		return nil
	}
	out := new(SaslObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SaslParameters) DeepCopyInto(out *SaslParameters) {
	*out = *in
	if in.IAM != nil {
		in, out := &in.IAM, &out.IAM
		*out = new(bool)
		**out = **in
	}
	if in.Scram != nil {
		in, out := &in.Scram, &out.Scram
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SaslParameters.
func (in *SaslParameters) DeepCopy() *SaslParameters {
	if in == nil {
		return nil
	}
	out := new(SaslParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScramSecretAssociation) DeepCopyInto(out *ScramSecretAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScramSecretAssociation.
func (in *ScramSecretAssociation) DeepCopy() *ScramSecretAssociation {
	if in == nil {
		return nil
	}
	out := new(ScramSecretAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScramSecretAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScramSecretAssociationList) DeepCopyInto(out *ScramSecretAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScramSecretAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScramSecretAssociationList.
func (in *ScramSecretAssociationList) DeepCopy() *ScramSecretAssociationList {
	if in == nil {
		return nil
	}
	out := new(ScramSecretAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScramSecretAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScramSecretAssociationObservation) DeepCopyInto(out *ScramSecretAssociationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScramSecretAssociationObservation.
func (in *ScramSecretAssociationObservation) DeepCopy() *ScramSecretAssociationObservation {
	if in == nil {
		return nil
	}
	out := new(ScramSecretAssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScramSecretAssociationParameters) DeepCopyInto(out *ScramSecretAssociationParameters) {
	*out = *in
	if in.ClusterArn != nil {
		in, out := &in.ClusterArn, &out.ClusterArn
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SecretArnList != nil {
		in, out := &in.SecretArnList, &out.SecretArnList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScramSecretAssociationParameters.
func (in *ScramSecretAssociationParameters) DeepCopy() *ScramSecretAssociationParameters {
	if in == nil {
		return nil
	}
	out := new(ScramSecretAssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScramSecretAssociationSpec) DeepCopyInto(out *ScramSecretAssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScramSecretAssociationSpec.
func (in *ScramSecretAssociationSpec) DeepCopy() *ScramSecretAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(ScramSecretAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScramSecretAssociationStatus) DeepCopyInto(out *ScramSecretAssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScramSecretAssociationStatus.
func (in *ScramSecretAssociationStatus) DeepCopy() *ScramSecretAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(ScramSecretAssociationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSObservation) DeepCopyInto(out *TLSObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSObservation.
func (in *TLSObservation) DeepCopy() *TLSObservation {
	if in == nil {
		return nil
	}
	out := new(TLSObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSParameters) DeepCopyInto(out *TLSParameters) {
	*out = *in
	if in.CertificateAuthorityArns != nil {
		in, out := &in.CertificateAuthorityArns, &out.CertificateAuthorityArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSParameters.
func (in *TLSParameters) DeepCopy() *TLSParameters {
	if in == nil {
		return nil
	}
	out := new(TLSParameters)
	in.DeepCopyInto(out)
	return out
}
