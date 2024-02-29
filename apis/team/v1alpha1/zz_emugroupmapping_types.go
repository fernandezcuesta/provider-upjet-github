// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type EmuGroupMappingInitParameters struct {

	// Integer corresponding to the external group ID to be linked
	// Integer corresponding to the external group ID to be linked.
	GroupID *float64 `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Slug of the GitHub team
	// Slug of the GitHub team.
	// +crossplane:generate:reference:type=Team
	TeamSlug *string `json:"teamSlug,omitempty" tf:"team_slug,omitempty"`

	// Reference to a Team to populate teamSlug.
	// +kubebuilder:validation:Optional
	TeamSlugRef *v1.Reference `json:"teamSlugRef,omitempty" tf:"-"`

	// Selector for a Team to populate teamSlug.
	// +kubebuilder:validation:Optional
	TeamSlugSelector *v1.Selector `json:"teamSlugSelector,omitempty" tf:"-"`
}

type EmuGroupMappingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Integer corresponding to the external group ID to be linked
	// Integer corresponding to the external group ID to be linked.
	GroupID *float64 `json:"groupId,omitempty" tf:"group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Slug of the GitHub team
	// Slug of the GitHub team.
	TeamSlug *string `json:"teamSlug,omitempty" tf:"team_slug,omitempty"`
}

type EmuGroupMappingParameters struct {

	// Integer corresponding to the external group ID to be linked
	// Integer corresponding to the external group ID to be linked.
	// +kubebuilder:validation:Optional
	GroupID *float64 `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Slug of the GitHub team
	// Slug of the GitHub team.
	// +crossplane:generate:reference:type=Team
	// +kubebuilder:validation:Optional
	TeamSlug *string `json:"teamSlug,omitempty" tf:"team_slug,omitempty"`

	// Reference to a Team to populate teamSlug.
	// +kubebuilder:validation:Optional
	TeamSlugRef *v1.Reference `json:"teamSlugRef,omitempty" tf:"-"`

	// Selector for a Team to populate teamSlug.
	// +kubebuilder:validation:Optional
	TeamSlugSelector *v1.Selector `json:"teamSlugSelector,omitempty" tf:"-"`
}

// EmuGroupMappingSpec defines the desired state of EmuGroupMapping
type EmuGroupMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EmuGroupMappingParameters `json:"forProvider"`
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
	InitProvider EmuGroupMappingInitParameters `json:"initProvider,omitempty"`
}

// EmuGroupMappingStatus defines the observed state of EmuGroupMapping.
type EmuGroupMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EmuGroupMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EmuGroupMapping is the Schema for the EmuGroupMappings API. Manages mappings between external groups for enterprise managed users.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type EmuGroupMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupId) || (has(self.initProvider) && has(self.initProvider.groupId))",message="spec.forProvider.groupId is a required parameter"
	Spec   EmuGroupMappingSpec   `json:"spec"`
	Status EmuGroupMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmuGroupMappingList contains a list of EmuGroupMappings
type EmuGroupMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmuGroupMapping `json:"items"`
}

// Repository type metadata.
var (
	EmuGroupMapping_Kind             = "EmuGroupMapping"
	EmuGroupMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EmuGroupMapping_Kind}.String()
	EmuGroupMapping_KindAPIVersion   = EmuGroupMapping_Kind + "." + CRDGroupVersion.String()
	EmuGroupMapping_GroupVersionKind = CRDGroupVersion.WithKind(EmuGroupMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&EmuGroupMapping{}, &EmuGroupMappingList{})
}
