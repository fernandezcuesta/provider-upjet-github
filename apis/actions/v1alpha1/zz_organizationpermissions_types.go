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

type AllowedActionsConfigInitParameters struct {

	// Whether GitHub-owned actions are allowed in the organization.
	// Whether GitHub-owned actions are allowed in the organization.
	GithubOwnedAllowed *bool `json:"githubOwnedAllowed,omitempty" tf:"github_owned_allowed,omitempty"`

	// Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@, monalisa/octocat@v2, monalisa/."
	// Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, 'monalisa/octocat@', 'monalisa/octocat@v2', 'monalisa/'.
	// +listType=set
	PatternsAllowed []*string `json:"patternsAllowed,omitempty" tf:"patterns_allowed,omitempty"`

	// Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.
	// Whether actions in GitHub Marketplace from verified creators are allowed. Set to 'true' to allow all GitHub Marketplace actions by verified creators.
	VerifiedAllowed *bool `json:"verifiedAllowed,omitempty" tf:"verified_allowed,omitempty"`
}

type AllowedActionsConfigObservation struct {

	// Whether GitHub-owned actions are allowed in the organization.
	// Whether GitHub-owned actions are allowed in the organization.
	GithubOwnedAllowed *bool `json:"githubOwnedAllowed,omitempty" tf:"github_owned_allowed,omitempty"`

	// Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@, monalisa/octocat@v2, monalisa/."
	// Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, 'monalisa/octocat@', 'monalisa/octocat@v2', 'monalisa/'.
	// +listType=set
	PatternsAllowed []*string `json:"patternsAllowed,omitempty" tf:"patterns_allowed,omitempty"`

	// Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.
	// Whether actions in GitHub Marketplace from verified creators are allowed. Set to 'true' to allow all GitHub Marketplace actions by verified creators.
	VerifiedAllowed *bool `json:"verifiedAllowed,omitempty" tf:"verified_allowed,omitempty"`
}

type AllowedActionsConfigParameters struct {

	// Whether GitHub-owned actions are allowed in the organization.
	// Whether GitHub-owned actions are allowed in the organization.
	// +kubebuilder:validation:Optional
	GithubOwnedAllowed *bool `json:"githubOwnedAllowed" tf:"github_owned_allowed,omitempty"`

	// Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@, monalisa/octocat@v2, monalisa/."
	// Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, 'monalisa/octocat@', 'monalisa/octocat@v2', 'monalisa/'.
	// +kubebuilder:validation:Optional
	// +listType=set
	PatternsAllowed []*string `json:"patternsAllowed,omitempty" tf:"patterns_allowed,omitempty"`

	// Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.
	// Whether actions in GitHub Marketplace from verified creators are allowed. Set to 'true' to allow all GitHub Marketplace actions by verified creators.
	// +kubebuilder:validation:Optional
	VerifiedAllowed *bool `json:"verifiedAllowed,omitempty" tf:"verified_allowed,omitempty"`
}

type EnabledRepositoriesConfigInitParameters struct {

	// List of repository IDs to enable for GitHub Actions.
	// List of repository IDs to enable for GitHub Actions.
	// +listType=set
	RepositoryIds []*int64 `json:"repositoryIds,omitempty" tf:"repository_ids,omitempty"`
}

type EnabledRepositoriesConfigObservation struct {

	// List of repository IDs to enable for GitHub Actions.
	// List of repository IDs to enable for GitHub Actions.
	// +listType=set
	RepositoryIds []*int64 `json:"repositoryIds,omitempty" tf:"repository_ids,omitempty"`
}

type EnabledRepositoriesConfigParameters struct {

	// List of repository IDs to enable for GitHub Actions.
	// List of repository IDs to enable for GitHub Actions.
	// +kubebuilder:validation:Optional
	// +listType=set
	RepositoryIds []*int64 `json:"repositoryIds" tf:"repository_ids,omitempty"`
}

type OrganizationPermissionsInitParameters struct {

	// The permissions policy that controls the actions that are allowed to run. Can be one of: all, local_only, or selected.
	// The permissions policy that controls the actions that are allowed to run. Can be one of: 'all', 'local_only', or 'selected'.
	AllowedActions *string `json:"allowedActions,omitempty" tf:"allowed_actions,omitempty"`

	// Sets the actions that are allowed in an organization. Only available when allowed_actions = selected. See Allowed Actions Config below for details.
	// Sets the actions that are allowed in an organization. Only available when 'allowed_actions' = 'selected'
	AllowedActionsConfig []AllowedActionsConfigInitParameters `json:"allowedActionsConfig,omitempty" tf:"allowed_actions_config,omitempty"`

	// The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: all, none, or selected.
	// The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: 'all', 'none', or 'selected'.
	EnabledRepositories *string `json:"enabledRepositories,omitempty" tf:"enabled_repositories,omitempty"`

	// Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when enabled_repositories = selected. See Enabled Repositories Config below for details.
	// Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when 'enabled_repositories' = 'selected'.
	EnabledRepositoriesConfig []EnabledRepositoriesConfigInitParameters `json:"enabledRepositoriesConfig,omitempty" tf:"enabled_repositories_config,omitempty"`
}

type OrganizationPermissionsObservation struct {

	// The permissions policy that controls the actions that are allowed to run. Can be one of: all, local_only, or selected.
	// The permissions policy that controls the actions that are allowed to run. Can be one of: 'all', 'local_only', or 'selected'.
	AllowedActions *string `json:"allowedActions,omitempty" tf:"allowed_actions,omitempty"`

	// Sets the actions that are allowed in an organization. Only available when allowed_actions = selected. See Allowed Actions Config below for details.
	// Sets the actions that are allowed in an organization. Only available when 'allowed_actions' = 'selected'
	AllowedActionsConfig []AllowedActionsConfigObservation `json:"allowedActionsConfig,omitempty" tf:"allowed_actions_config,omitempty"`

	// The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: all, none, or selected.
	// The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: 'all', 'none', or 'selected'.
	EnabledRepositories *string `json:"enabledRepositories,omitempty" tf:"enabled_repositories,omitempty"`

	// Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when enabled_repositories = selected. See Enabled Repositories Config below for details.
	// Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when 'enabled_repositories' = 'selected'.
	EnabledRepositoriesConfig []EnabledRepositoriesConfigObservation `json:"enabledRepositoriesConfig,omitempty" tf:"enabled_repositories_config,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OrganizationPermissionsParameters struct {

	// The permissions policy that controls the actions that are allowed to run. Can be one of: all, local_only, or selected.
	// The permissions policy that controls the actions that are allowed to run. Can be one of: 'all', 'local_only', or 'selected'.
	// +kubebuilder:validation:Optional
	AllowedActions *string `json:"allowedActions,omitempty" tf:"allowed_actions,omitempty"`

	// Sets the actions that are allowed in an organization. Only available when allowed_actions = selected. See Allowed Actions Config below for details.
	// Sets the actions that are allowed in an organization. Only available when 'allowed_actions' = 'selected'
	// +kubebuilder:validation:Optional
	AllowedActionsConfig []AllowedActionsConfigParameters `json:"allowedActionsConfig,omitempty" tf:"allowed_actions_config,omitempty"`

	// The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: all, none, or selected.
	// The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: 'all', 'none', or 'selected'.
	// +kubebuilder:validation:Optional
	EnabledRepositories *string `json:"enabledRepositories,omitempty" tf:"enabled_repositories,omitempty"`

	// Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when enabled_repositories = selected. See Enabled Repositories Config below for details.
	// Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when 'enabled_repositories' = 'selected'.
	// +kubebuilder:validation:Optional
	EnabledRepositoriesConfig []EnabledRepositoriesConfigParameters `json:"enabledRepositoriesConfig,omitempty" tf:"enabled_repositories_config,omitempty"`
}

// OrganizationPermissionsSpec defines the desired state of OrganizationPermissions
type OrganizationPermissionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationPermissionsParameters `json:"forProvider"`
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
	InitProvider OrganizationPermissionsInitParameters `json:"initProvider,omitempty"`
}

// OrganizationPermissionsStatus defines the observed state of OrganizationPermissions.
type OrganizationPermissionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationPermissionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OrganizationPermissions is the Schema for the OrganizationPermissionss API. Creates and manages Actions permissions within a GitHub organization
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type OrganizationPermissions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabledRepositories) || (has(self.initProvider) && has(self.initProvider.enabledRepositories))",message="spec.forProvider.enabledRepositories is a required parameter"
	Spec   OrganizationPermissionsSpec   `json:"spec"`
	Status OrganizationPermissionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationPermissionsList contains a list of OrganizationPermissionss
type OrganizationPermissionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationPermissions `json:"items"`
}

// Repository type metadata.
var (
	OrganizationPermissions_Kind             = "OrganizationPermissions"
	OrganizationPermissions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationPermissions_Kind}.String()
	OrganizationPermissions_KindAPIVersion   = OrganizationPermissions_Kind + "." + CRDGroupVersion.String()
	OrganizationPermissions_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationPermissions_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationPermissions{}, &OrganizationPermissionsList{})
}
