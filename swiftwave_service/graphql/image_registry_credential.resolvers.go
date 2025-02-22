package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"

	"github.com/swiftwave-org/swiftwave/swiftwave_service/core"
	"github.com/swiftwave-org/swiftwave/swiftwave_service/graphql/model"
)

// Deployments is the resolver for the deployments field.
func (r *imageRegistryCredentialResolver) Deployments(ctx context.Context, obj *model.ImageRegistryCredential) ([]*model.Deployment, error) {
	// fetch record
	records, err := core.FindDeploymentsByImageRegistryCredentialId(ctx, r.ServiceManager.DbClient, obj.ID)
	if err != nil {
		return nil, err
	}
	// convert to graphql object
	var result = make([]*model.Deployment, 0)
	for _, record := range records {
		result = append(result, deploymentToGraphqlObject(record))
	}
	return result, nil
}

// CreateImageRegistryCredential is the resolver for the createImageRegistryCredential field.
func (r *mutationResolver) CreateImageRegistryCredential(ctx context.Context, input model.ImageRegistryCredentialInput) (*model.ImageRegistryCredential, error) {
	record := imageRegistryCredentialInputToDatabaseObject(&input)
	err := record.Create(ctx, r.ServiceManager.DbClient)
	if err != nil {
		return nil, err
	}
	return imageRegistryCredentialToGraphqlObject(record), nil
}

// UpdateImageRegistryCredential is the resolver for the updateImageRegistryCredential field.
func (r *mutationResolver) UpdateImageRegistryCredential(ctx context.Context, id uint, input model.ImageRegistryCredentialInput) (*model.ImageRegistryCredential, error) {
	// fetch record
	var record = &core.ImageRegistryCredential{}
	err := record.FindById(ctx, r.ServiceManager.DbClient, id)
	if err != nil {
		return nil, err
	}
	// update record
	record.Url = input.URL
	record.Username = input.Username
	record.Password = input.Password
	err = record.Update(ctx, r.ServiceManager.DbClient)
	if err != nil {
		return nil, err
	}
	return imageRegistryCredentialToGraphqlObject(record), nil
}

// DeleteImageRegistryCredential is the resolver for the deleteImageRegistryCredential field.
func (r *mutationResolver) DeleteImageRegistryCredential(ctx context.Context, id uint) (bool, error) {
	// fetch record
	var record = &core.ImageRegistryCredential{}
	err := record.FindById(ctx, r.ServiceManager.DbClient, id)
	if err != nil {
		return false, err
	}
	// delete record
	tx := r.ServiceManager.DbClient.Begin()
	err = record.Delete(ctx, *tx)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	err = tx.Commit().Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// ImageRegistryCredentials is the resolver for the imageRegistryCredentials field.
func (r *queryResolver) ImageRegistryCredentials(ctx context.Context) ([]*model.ImageRegistryCredential, error) {
	records, err := core.FindAllImageRegistryCredentials(ctx, r.ServiceManager.DbClient)
	if err != nil {
		return nil, err
	}
	var result []*model.ImageRegistryCredential
	for _, record := range records {
		result = append(result, imageRegistryCredentialToGraphqlObject(record))
	}
	return result, nil
}

// ImageRegistryCredential is the resolver for the imageRegistryCredential field.
func (r *queryResolver) ImageRegistryCredential(ctx context.Context, id uint) (*model.ImageRegistryCredential, error) {
	var record = &core.ImageRegistryCredential{}
	err := record.FindById(ctx, r.ServiceManager.DbClient, id)
	if err != nil {
		return nil, err
	}
	return imageRegistryCredentialToGraphqlObject(record), nil
}

// ImageRegistryCredential returns ImageRegistryCredentialResolver implementation.
func (r *Resolver) ImageRegistryCredential() ImageRegistryCredentialResolver {
	return &imageRegistryCredentialResolver{r}
}

type imageRegistryCredentialResolver struct{ *Resolver }
