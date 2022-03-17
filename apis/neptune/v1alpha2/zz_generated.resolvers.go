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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha2

import (
	"context"
	v1alpha22 "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2"
	v1alpha2 "github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2"
	v1alpha21 "github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha2"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.IAMRoles),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.IAMRoleIdRefs,
		Selector:      mg.Spec.ForProvider.IAMRoleIdSelector,
		To: reference.To{
			List:    &v1alpha2.RoleList{},
			Managed: &v1alpha2.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IAMRoles")
	}
	mg.Spec.ForProvider.IAMRoles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.IAMRoleIdRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyArn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyArnRef,
		Selector:     mg.Spec.ForProvider.KMSKeyArnSelector,
		To: reference.To{
			List:    &v1alpha21.KeyList{},
			Managed: &v1alpha21.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyArn")
	}
	mg.Spec.ForProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NeptuneClusterParameterGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NeptuneClusterParameterGroupNameRef,
		Selector:     mg.Spec.ForProvider.NeptuneClusterParameterGroupNameSelector,
		To: reference.To{
			List:    &ClusterParameterGroupList{},
			Managed: &ClusterParameterGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NeptuneClusterParameterGroupName")
	}
	mg.Spec.ForProvider.NeptuneClusterParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NeptuneClusterParameterGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NeptuneSubnetGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NeptuneSubnetGroupNameRef,
		Selector:     mg.Spec.ForProvider.NeptuneSubnetGroupNameSelector,
		To: reference.To{
			List:    &SubnetGroupList{},
			Managed: &SubnetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NeptuneSubnetGroupName")
	}
	mg.Spec.ForProvider.NeptuneSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NeptuneSubnetGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ReplicationSourceIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ReplicationSourceIdentifierRef,
		Selector:     mg.Spec.ForProvider.ReplicationSourceIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ReplicationSourceIdentifier")
	}
	mg.Spec.ForProvider.ReplicationSourceIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ReplicationSourceIdentifierRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnapshotIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SnapshotIdentifierRef,
		Selector:     mg.Spec.ForProvider.SnapshotIdentifierSelector,
		To: reference.To{
			List:    &ClusterSnapshotList{},
			Managed: &ClusterSnapshot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnapshotIdentifier")
	}
	mg.Spec.ForProvider.SnapshotIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnapshotIdentifierRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VPCSecurityGroupIdRefs,
		Selector:      mg.Spec.ForProvider.VPCSecurityGroupIdSelector,
		To: reference.To{
			List:    &v1alpha22.SecurityGroupList{},
			Managed: &v1alpha22.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCSecurityGroupIds")
	}
	mg.Spec.ForProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCSecurityGroupIdRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this ClusterEndpoint.
func (mg *ClusterEndpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.ClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterIdentifier")
	}
	mg.Spec.ForProvider.ClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClusterInstance.
func (mg *ClusterInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.ClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterIdentifier")
	}
	mg.Spec.ForProvider.ClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIdentifierRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NeptuneParameterGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NeptuneParameterGroupNameRef,
		Selector:     mg.Spec.ForProvider.NeptuneParameterGroupNameSelector,
		To: reference.To{
			List:    &ParameterGroupList{},
			Managed: &ParameterGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NeptuneParameterGroupName")
	}
	mg.Spec.ForProvider.NeptuneParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NeptuneParameterGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NeptuneSubnetGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NeptuneSubnetGroupNameRef,
		Selector:     mg.Spec.ForProvider.NeptuneSubnetGroupNameSelector,
		To: reference.To{
			List:    &SubnetGroupList{},
			Managed: &SubnetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NeptuneSubnetGroupName")
	}
	mg.Spec.ForProvider.NeptuneSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NeptuneSubnetGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClusterSnapshot.
func (mg *ClusterSnapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBClusterIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DBClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.DBClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBClusterIdentifier")
	}
	mg.Spec.ForProvider.DBClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetGroup.
func (mg *SubnetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetIdRefs,
		Selector:      mg.Spec.ForProvider.SubnetIdSelector,
		To: reference.To{
			List:    &v1alpha22.SubnetList{},
			Managed: &v1alpha22.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIdRefs = mrsp.ResolvedReferences

	return nil
}