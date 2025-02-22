package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"
	"errors"

	"github.com/swiftwave-org/swiftwave/swiftwave_service/core"
	"github.com/swiftwave-org/swiftwave/swiftwave_service/graphql/model"
)

// CreateRedirectRule is the resolver for the createRedirectRule field.
func (r *mutationResolver) CreateRedirectRule(ctx context.Context, input model.RedirectRuleInput) (*model.RedirectRule, error) {
	record := redirectRuleInputToDatabaseObject(&input)
	err := record.Create(ctx, r.ServiceManager.DbClient)
	if err != nil {
		return nil, err
	}
	// publish event
	err = r.WorkerManager.EnqueueRedirectRuleApplyRequest(record.ID)
	if err != nil {
		return nil, errors.New("failed to enqueue redirect rule apply request")
	}
	return redirectRuleToGraphqlObject(record), nil
}

// DeleteRedirectRule is the resolver for the deleteRedirectRule field.
func (r *mutationResolver) DeleteRedirectRule(ctx context.Context, id uint) (bool, error) {
	record := core.RedirectRule{}
	err := record.FindById(ctx, r.ServiceManager.DbClient, id)
	if err != nil {
		return false, err
	}
	err = record.Delete(ctx, r.ServiceManager.DbClient, false)
	if err != nil {
		return false, err
	}
	// publish event
	err = r.WorkerManager.EnqueueRedirectRuleDeleteRequest(record.ID)
	if err != nil {
		return false, errors.New("failed to enqueue redirect rule delete request")
	}
	return true, nil
}

// RedirectRule is the resolver for the redirectRule field.
func (r *queryResolver) RedirectRule(ctx context.Context, id uint) (*model.RedirectRule, error) {
	record := core.RedirectRule{}
	err := record.FindById(ctx, r.ServiceManager.DbClient, id)
	if err != nil {
		return nil, err
	}
	return redirectRuleToGraphqlObject(&record), nil
}

// RedirectRules is the resolver for the redirectRules field.
func (r *queryResolver) RedirectRules(ctx context.Context) ([]*model.RedirectRule, error) {
	records, err := core.FindAllRedirectRules(ctx, r.ServiceManager.DbClient)
	if err != nil {
		return nil, err
	}
	var result []*model.RedirectRule
	for _, record := range records {
		result = append(result, redirectRuleToGraphqlObject(record))
	}
	return result, nil
}

// Domain is the resolver for the domain field.
func (r *redirectRuleResolver) Domain(ctx context.Context, obj *model.RedirectRule) (*model.Domain, error) {
	domain := &core.Domain{}
	err := domain.FindById(ctx, r.ServiceManager.DbClient, obj.DomainID)
	if err != nil {
		return nil, err
	}
	return domainToGraphqlObject(domain), nil
}

// RedirectRule returns RedirectRuleResolver implementation.
func (r *Resolver) RedirectRule() RedirectRuleResolver { return &redirectRuleResolver{r} }

type redirectRuleResolver struct{ *Resolver }
