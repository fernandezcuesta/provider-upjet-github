/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ActionsVariableInitParameters struct {

	// Name of the repository
	// Name of the repository.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-github/apis/repo/v1alpha1.Repository
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// Value of the variable
	// Value of the variable.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Name of the variable
	// Name of the variable.
	VariableName *string `json:"variableName,omitempty" tf:"variable_name,omitempty"`
}

type ActionsVariableObservation struct {

	// Date of actions_variable creation.
	// Date of 'actions_variable' creation.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the repository
	// Name of the repository.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Date of actions_variable update.
	// Date of 'actions_variable' update.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// Value of the variable
	// Value of the variable.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Name of the variable
	// Name of the variable.
	VariableName *string `json:"variableName,omitempty" tf:"variable_name,omitempty"`
}

type ActionsVariableParameters struct {

	// Name of the repository
	// Name of the repository.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-github/apis/repo/v1alpha1.Repository
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// Value of the variable
	// Value of the variable.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Name of the variable
	// Name of the variable.
	// +kubebuilder:validation:Optional
	VariableName *string `json:"variableName,omitempty" tf:"variable_name,omitempty"`
}

// ActionsVariableSpec defines the desired state of ActionsVariable
type ActionsVariableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActionsVariableParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ActionsVariableInitParameters `json:"initProvider,omitempty"`
}

// ActionsVariableStatus defines the observed state of ActionsVariable.
type ActionsVariableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActionsVariableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ActionsVariable is the Schema for the ActionsVariables API. Creates and manages an Action variable within a GitHub repository
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type ActionsVariable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.variableName) || (has(self.initProvider) && has(self.initProvider.variableName))",message="spec.forProvider.variableName is a required parameter"
	Spec   ActionsVariableSpec   `json:"spec"`
	Status ActionsVariableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActionsVariableList contains a list of ActionsVariables
type ActionsVariableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActionsVariable `json:"items"`
}

// Repository type metadata.
var (
	ActionsVariable_Kind             = "ActionsVariable"
	ActionsVariable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ActionsVariable_Kind}.String()
	ActionsVariable_KindAPIVersion   = ActionsVariable_Kind + "." + CRDGroupVersion.String()
	ActionsVariable_GroupVersionKind = CRDGroupVersion.WithKind(ActionsVariable_Kind)
)

func init() {
	SchemeBuilder.Register(&ActionsVariable{}, &ActionsVariableList{})
}
