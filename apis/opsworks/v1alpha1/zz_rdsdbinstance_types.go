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

type RDSDBInstanceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RDSDBInstanceParameters struct {

	// +kubebuilder:validation:Required
	DBPasswordSecretRef v1.SecretKeySelector `json:"dbPasswordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	DBUser *string `json:"dbUser" tf:"db_user,omitempty"`

	// +kubebuilder:validation:Required
	RDSDBInstanceArn *string `json:"rdsDbInstanceArn" tf:"rds_db_instance_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	StackID *string `json:"stackId" tf:"stack_id,omitempty"`
}

// RDSDBInstanceSpec defines the desired state of RDSDBInstance
type RDSDBInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RDSDBInstanceParameters `json:"forProvider"`
}

// RDSDBInstanceStatus defines the observed state of RDSDBInstance.
type RDSDBInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RDSDBInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RDSDBInstance is the Schema for the RDSDBInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type RDSDBInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RDSDBInstanceSpec   `json:"spec"`
	Status            RDSDBInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RDSDBInstanceList contains a list of RDSDBInstances
type RDSDBInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RDSDBInstance `json:"items"`
}

// Repository type metadata.
var (
	RDSDBInstance_Kind             = "RDSDBInstance"
	RDSDBInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RDSDBInstance_Kind}.String()
	RDSDBInstance_KindAPIVersion   = RDSDBInstance_Kind + "." + CRDGroupVersion.String()
	RDSDBInstance_GroupVersionKind = CRDGroupVersion.WithKind(RDSDBInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&RDSDBInstance{}, &RDSDBInstanceList{})
}
