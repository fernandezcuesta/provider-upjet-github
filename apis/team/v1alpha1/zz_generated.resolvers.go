/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/coopnorge/provider-github/apis/repo/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this TeamMembership.
func (mg *TeamMembership) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TeamID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TeamIDRef,
		Selector:     mg.Spec.ForProvider.TeamIDSelector,
		To: reference.To{
			List:    &TeamList{},
			Managed: &Team{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TeamID")
	}
	mg.Spec.ForProvider.TeamID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TeamIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TeamID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.TeamIDRef,
		Selector:     mg.Spec.InitProvider.TeamIDSelector,
		To: reference.To{
			List:    &TeamList{},
			Managed: &Team{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TeamID")
	}
	mg.Spec.InitProvider.TeamID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TeamIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TeamRepository.
func (mg *TeamRepository) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &v1alpha1.RepositoryList{},
			Managed: &v1alpha1.Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TeamID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TeamIDRef,
		Selector:     mg.Spec.ForProvider.TeamIDSelector,
		To: reference.To{
			List:    &TeamList{},
			Managed: &Team{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TeamID")
	}
	mg.Spec.ForProvider.TeamID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TeamIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &v1alpha1.RepositoryList{},
			Managed: &v1alpha1.Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TeamID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.TeamIDRef,
		Selector:     mg.Spec.InitProvider.TeamIDSelector,
		To: reference.To{
			List:    &TeamList{},
			Managed: &Team{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TeamID")
	}
	mg.Spec.InitProvider.TeamID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TeamIDRef = rsp.ResolvedReference

	return nil
}
