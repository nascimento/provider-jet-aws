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
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Application
func (mg *Application) GetTerraformResourceType() string {
	return "aws_appconfig_application"
}

// GetConnectionDetailsMapping for this Application
func (tr *Application) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Application
func (tr *Application) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Application
func (tr *Application) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Application
func (tr *Application) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Application
func (tr *Application) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Application
func (tr *Application) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Application using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Application) LateInitialize(attrs []byte) (bool, error) {
	params := &ApplicationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Application) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ConfigurationProfile
func (mg *ConfigurationProfile) GetTerraformResourceType() string {
	return "aws_appconfig_configuration_profile"
}

// GetConnectionDetailsMapping for this ConfigurationProfile
func (tr *ConfigurationProfile) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"validator[*].content": "spec.forProvider.validator[*].contentSecretRef"}
}

// GetObservation of this ConfigurationProfile
func (tr *ConfigurationProfile) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ConfigurationProfile
func (tr *ConfigurationProfile) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ConfigurationProfile
func (tr *ConfigurationProfile) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ConfigurationProfile
func (tr *ConfigurationProfile) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ConfigurationProfile
func (tr *ConfigurationProfile) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ConfigurationProfile using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ConfigurationProfile) LateInitialize(attrs []byte) (bool, error) {
	params := &ConfigurationProfileParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ConfigurationProfile) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Deployment
func (mg *Deployment) GetTerraformResourceType() string {
	return "aws_appconfig_deployment"
}

// GetConnectionDetailsMapping for this Deployment
func (tr *Deployment) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Deployment
func (tr *Deployment) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Deployment
func (tr *Deployment) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Deployment
func (tr *Deployment) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Deployment
func (tr *Deployment) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Deployment
func (tr *Deployment) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Deployment using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Deployment) LateInitialize(attrs []byte) (bool, error) {
	params := &DeploymentParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Deployment) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this DeploymentStrategy
func (mg *DeploymentStrategy) GetTerraformResourceType() string {
	return "aws_appconfig_deployment_strategy"
}

// GetConnectionDetailsMapping for this DeploymentStrategy
func (tr *DeploymentStrategy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this DeploymentStrategy
func (tr *DeploymentStrategy) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this DeploymentStrategy
func (tr *DeploymentStrategy) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this DeploymentStrategy
func (tr *DeploymentStrategy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this DeploymentStrategy
func (tr *DeploymentStrategy) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this DeploymentStrategy
func (tr *DeploymentStrategy) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this DeploymentStrategy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *DeploymentStrategy) LateInitialize(attrs []byte) (bool, error) {
	params := &DeploymentStrategyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *DeploymentStrategy) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Environment
func (mg *Environment) GetTerraformResourceType() string {
	return "aws_appconfig_environment"
}

// GetConnectionDetailsMapping for this Environment
func (tr *Environment) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Environment
func (tr *Environment) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Environment
func (tr *Environment) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Environment
func (tr *Environment) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Environment
func (tr *Environment) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Environment
func (tr *Environment) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Environment using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Environment) LateInitialize(attrs []byte) (bool, error) {
	params := &EnvironmentParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Environment) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this HostedConfigurationVersion
func (mg *HostedConfigurationVersion) GetTerraformResourceType() string {
	return "aws_appconfig_hosted_configuration_version"
}

// GetConnectionDetailsMapping for this HostedConfigurationVersion
func (tr *HostedConfigurationVersion) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"content": "spec.forProvider.contentSecretRef"}
}

// GetObservation of this HostedConfigurationVersion
func (tr *HostedConfigurationVersion) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this HostedConfigurationVersion
func (tr *HostedConfigurationVersion) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this HostedConfigurationVersion
func (tr *HostedConfigurationVersion) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this HostedConfigurationVersion
func (tr *HostedConfigurationVersion) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this HostedConfigurationVersion
func (tr *HostedConfigurationVersion) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this HostedConfigurationVersion using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HostedConfigurationVersion) LateInitialize(attrs []byte) (bool, error) {
	params := &HostedConfigurationVersionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *HostedConfigurationVersion) GetTerraformSchemaVersion() int {
	return 0
}
