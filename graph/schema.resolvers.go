package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/valllabh/security-cli/graph/model"
)

func (r *mutationResolver) AddPackage(ctx context.Context, input []*model.AddPackageInput) (*model.AddPackagePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePackage(ctx context.Context, input model.UpdatePackageInput) (*model.UpdatePackagePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePackage(ctx context.Context, filter model.PackageFilter) (*model.DeletePackagePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddLicense(ctx context.Context, input []*model.AddLicenseInput, upsert *bool) (*model.AddLicensePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteLicense(ctx context.Context, filter model.LicenseFilter) (*model.DeleteLicensePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddPackageVersion(ctx context.Context, input []*model.AddPackageVersionInput, upsert *bool) (*model.AddPackageVersionPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePackageVersion(ctx context.Context, filter model.PackageVersionFilter) (*model.DeletePackageVersionPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddProgrammingLanguage(ctx context.Context, input []*model.AddProgrammingLanguageInput, upsert *bool) (*model.AddProgrammingLanguagePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProgrammingLanguage(ctx context.Context, filter model.ProgrammingLanguageFilter) (*model.DeleteProgrammingLanguagePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddPackageType(ctx context.Context, input []*model.AddPackageTypeInput, upsert *bool) (*model.AddPackageTypePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePackageType(ctx context.Context, filter model.PackageTypeFilter) (*model.DeletePackageTypePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryPackage(ctx context.Context, filter *model.PackageFilter, order *model.PackageOrder, first *int, offset *int) ([]*model.Package, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AggregatePackage(ctx context.Context, filter *model.PackageFilter) (*model.PackageAggregateResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetLicense(ctx context.Context, id string) (*model.License, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryLicense(ctx context.Context, filter *model.LicenseFilter, order *model.LicenseOrder, first *int, offset *int) ([]*model.License, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AggregateLicense(ctx context.Context, filter *model.LicenseFilter) (*model.LicenseAggregateResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPackageVersion(ctx context.Context, id string) (*model.PackageVersion, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryPackageVersion(ctx context.Context, filter *model.PackageVersionFilter, order *model.PackageVersionOrder, first *int, offset *int) ([]*model.PackageVersion, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AggregatePackageVersion(ctx context.Context, filter *model.PackageVersionFilter) (*model.PackageVersionAggregateResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProgrammingLanguage(ctx context.Context, id string) (*model.ProgrammingLanguage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryProgrammingLanguage(ctx context.Context, filter *model.ProgrammingLanguageFilter, order *model.ProgrammingLanguageOrder, first *int, offset *int) ([]*model.ProgrammingLanguage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AggregateProgrammingLanguage(ctx context.Context, filter *model.ProgrammingLanguageFilter) (*model.ProgrammingLanguageAggregateResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPackageType(ctx context.Context, id string) (*model.PackageType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryPackageType(ctx context.Context, filter *model.PackageTypeFilter, order *model.PackageTypeOrder, first *int, offset *int) ([]*model.PackageType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AggregatePackageType(ctx context.Context, filter *model.PackageTypeFilter) (*model.PackageTypeAggregateResult, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
