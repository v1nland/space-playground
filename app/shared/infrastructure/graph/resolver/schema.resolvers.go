package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"space-playground/app/shared/infrastructure/graph"
	"space-playground/app/shared/infrastructure/graph/model"
)

// CreateAstronaut is the resolver for the createAstronaut field.
func (r *mutationResolver) CreateAstronaut(ctx context.Context, input model.NewAstronaut) (*model.Astronaut, error) {
	panic(fmt.Errorf("not implemented: CreateAstronaut - createAstronaut"))
}

// GetAstronautByID is the resolver for the getAstronautById field.
func (r *queryResolver) GetAstronautByID(ctx context.Context, id string) (*model.Astronaut, error) {
	panic(fmt.Errorf("not implemented: GetAstronautByID - getAstronautById"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }