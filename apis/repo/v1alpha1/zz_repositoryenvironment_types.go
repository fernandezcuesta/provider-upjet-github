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

type DeploymentBranchPolicyInitParameters struct {

	// Whether only branches that match the specified name patterns can deploy to this environment.
	// Whether only branches that match the specified name patterns can deploy to this environment.
	CustomBranchPolicies *bool `json:"customBranchPolicies,omitempty" tf:"custom_branch_policies,omitempty"`

	// Whether only branches with branch protection rules can deploy to this environment.
	// Whether only branches with branch protection rules can deploy to this environment.
	ProtectedBranches *bool `json:"protectedBranches,omitempty" tf:"protected_branches,omitempty"`
}

type DeploymentBranchPolicyObservation struct {

	// Whether only branches that match the specified name patterns can deploy to this environment.
	// Whether only branches that match the specified name patterns can deploy to this environment.
	CustomBranchPolicies *bool `json:"customBranchPolicies,omitempty" tf:"custom_branch_policies,omitempty"`

	// Whether only branches with branch protection rules can deploy to this environment.
	// Whether only branches with branch protection rules can deploy to this environment.
	ProtectedBranches *bool `json:"protectedBranches,omitempty" tf:"protected_branches,omitempty"`
}

type DeploymentBranchPolicyParameters struct {

	// Whether only branches that match the specified name patterns can deploy to this environment.
	// Whether only branches that match the specified name patterns can deploy to this environment.
	// +kubebuilder:validation:Optional
	CustomBranchPolicies *bool `json:"customBranchPolicies" tf:"custom_branch_policies,omitempty"`

	// Whether only branches with branch protection rules can deploy to this environment.
	// Whether only branches with branch protection rules can deploy to this environment.
	// +kubebuilder:validation:Optional
	ProtectedBranches *bool `json:"protectedBranches" tf:"protected_branches,omitempty"`
}

type RepositoryEnvironmentInitParameters struct {

	// Can repository admins bypass the environment protections.  Defaults to true.
	// Can Admins bypass deployment protections
	CanAdminsBypass *bool `json:"canAdminsBypass,omitempty" tf:"can_admins_bypass,omitempty"`

	// The deployment branch policy configuration
	DeploymentBranchPolicy []DeploymentBranchPolicyInitParameters `json:"deploymentBranchPolicy,omitempty" tf:"deployment_branch_policy,omitempty"`

	// The name of the environment.
	// The name of the environment.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// Whether or not a user who created the job is prevented from approving their own job. Defaults to false.
	// Prevent users from approving workflows runs that they triggered.
	PreventSelfReview *bool `json:"preventSelfReview,omitempty" tf:"prevent_self_review,omitempty"`

	// The repository of the environment.
	// The repository of the environment.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-github/apis/repo/v1alpha1.Repository
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// The environment reviewers configuration.
	Reviewers []ReviewersInitParameters `json:"reviewers,omitempty" tf:"reviewers,omitempty"`

	// Amount of time to delay a job after the job is initially triggered.
	// Amount of time to delay a job after the job is initially triggered.
	WaitTimer *int64 `json:"waitTimer,omitempty" tf:"wait_timer,omitempty"`
}

type RepositoryEnvironmentObservation struct {

	// Can repository admins bypass the environment protections.  Defaults to true.
	// Can Admins bypass deployment protections
	CanAdminsBypass *bool `json:"canAdminsBypass,omitempty" tf:"can_admins_bypass,omitempty"`

	// The deployment branch policy configuration
	DeploymentBranchPolicy []DeploymentBranchPolicyObservation `json:"deploymentBranchPolicy,omitempty" tf:"deployment_branch_policy,omitempty"`

	// The name of the environment.
	// The name of the environment.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether or not a user who created the job is prevented from approving their own job. Defaults to false.
	// Prevent users from approving workflows runs that they triggered.
	PreventSelfReview *bool `json:"preventSelfReview,omitempty" tf:"prevent_self_review,omitempty"`

	// The repository of the environment.
	// The repository of the environment.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// The environment reviewers configuration.
	Reviewers []ReviewersObservation `json:"reviewers,omitempty" tf:"reviewers,omitempty"`

	// Amount of time to delay a job after the job is initially triggered.
	// Amount of time to delay a job after the job is initially triggered.
	WaitTimer *int64 `json:"waitTimer,omitempty" tf:"wait_timer,omitempty"`
}

type RepositoryEnvironmentParameters struct {

	// Can repository admins bypass the environment protections.  Defaults to true.
	// Can Admins bypass deployment protections
	// +kubebuilder:validation:Optional
	CanAdminsBypass *bool `json:"canAdminsBypass,omitempty" tf:"can_admins_bypass,omitempty"`

	// The deployment branch policy configuration
	// +kubebuilder:validation:Optional
	DeploymentBranchPolicy []DeploymentBranchPolicyParameters `json:"deploymentBranchPolicy,omitempty" tf:"deployment_branch_policy,omitempty"`

	// The name of the environment.
	// The name of the environment.
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// Whether or not a user who created the job is prevented from approving their own job. Defaults to false.
	// Prevent users from approving workflows runs that they triggered.
	// +kubebuilder:validation:Optional
	PreventSelfReview *bool `json:"preventSelfReview,omitempty" tf:"prevent_self_review,omitempty"`

	// The repository of the environment.
	// The repository of the environment.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-github/apis/repo/v1alpha1.Repository
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// The environment reviewers configuration.
	// +kubebuilder:validation:Optional
	Reviewers []ReviewersParameters `json:"reviewers,omitempty" tf:"reviewers,omitempty"`

	// Amount of time to delay a job after the job is initially triggered.
	// Amount of time to delay a job after the job is initially triggered.
	// +kubebuilder:validation:Optional
	WaitTimer *int64 `json:"waitTimer,omitempty" tf:"wait_timer,omitempty"`
}

type ReviewersInitParameters struct {

	// Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-github/apis/team/v1alpha1.Team
	// +listType=set
	Teams []*int64 `json:"teams,omitempty" tf:"teams,omitempty"`

	// References to Team in team to populate teams.
	// +kubebuilder:validation:Optional
	TeamsRefs []v1.Reference `json:"teamsRefs,omitempty" tf:"-"`

	// Selector for a list of Team in team to populate teams.
	// +kubebuilder:validation:Optional
	TeamsSelector *v1.Selector `json:"teamsSelector,omitempty" tf:"-"`

	// Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-github/apis/user/v1alpha1.Membership
	// +listType=set
	Users []*int64 `json:"users,omitempty" tf:"users,omitempty"`

	// References to Membership in user to populate users.
	// +kubebuilder:validation:Optional
	UsersRefs []v1.Reference `json:"usersRefs,omitempty" tf:"-"`

	// Selector for a list of Membership in user to populate users.
	// +kubebuilder:validation:Optional
	UsersSelector *v1.Selector `json:"usersSelector,omitempty" tf:"-"`
}

type ReviewersObservation struct {

	// Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// +listType=set
	Teams []*int64 `json:"teams,omitempty" tf:"teams,omitempty"`

	// Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// +listType=set
	Users []*int64 `json:"users,omitempty" tf:"users,omitempty"`
}

type ReviewersParameters struct {

	// Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-github/apis/team/v1alpha1.Team
	// +kubebuilder:validation:Optional
	// +listType=set
	Teams []*int64 `json:"teams,omitempty" tf:"teams,omitempty"`

	// References to Team in team to populate teams.
	// +kubebuilder:validation:Optional
	TeamsRefs []v1.Reference `json:"teamsRefs,omitempty" tf:"-"`

	// Selector for a list of Team in team to populate teams.
	// +kubebuilder:validation:Optional
	TeamsSelector *v1.Selector `json:"teamsSelector,omitempty" tf:"-"`

	// Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-github/apis/user/v1alpha1.Membership
	// +kubebuilder:validation:Optional
	// +listType=set
	Users []*int64 `json:"users,omitempty" tf:"users,omitempty"`

	// References to Membership in user to populate users.
	// +kubebuilder:validation:Optional
	UsersRefs []v1.Reference `json:"usersRefs,omitempty" tf:"-"`

	// Selector for a list of Membership in user to populate users.
	// +kubebuilder:validation:Optional
	UsersSelector *v1.Selector `json:"usersSelector,omitempty" tf:"-"`
}

// RepositoryEnvironmentSpec defines the desired state of RepositoryEnvironment
type RepositoryEnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryEnvironmentParameters `json:"forProvider"`
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
	InitProvider RepositoryEnvironmentInitParameters `json:"initProvider,omitempty"`
}

// RepositoryEnvironmentStatus defines the observed state of RepositoryEnvironment.
type RepositoryEnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryEnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RepositoryEnvironment is the Schema for the RepositoryEnvironments API. Creates and manages environments for GitHub repositories
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type RepositoryEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.environment) || (has(self.initProvider) && has(self.initProvider.environment))",message="spec.forProvider.environment is a required parameter"
	Spec   RepositoryEnvironmentSpec   `json:"spec"`
	Status RepositoryEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryEnvironmentList contains a list of RepositoryEnvironments
type RepositoryEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RepositoryEnvironment `json:"items"`
}

// Repository type metadata.
var (
	RepositoryEnvironment_Kind             = "RepositoryEnvironment"
	RepositoryEnvironment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RepositoryEnvironment_Kind}.String()
	RepositoryEnvironment_KindAPIVersion   = RepositoryEnvironment_Kind + "." + CRDGroupVersion.String()
	RepositoryEnvironment_GroupVersionKind = CRDGroupVersion.WithKind(RepositoryEnvironment_Kind)
)

func init() {
	SchemeBuilder.Register(&RepositoryEnvironment{}, &RepositoryEnvironmentList{})
}
