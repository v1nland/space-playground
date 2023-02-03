package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"
	"space-playground/app/missions/domain"
	"space-playground/app/shared/infrastructure/graph"
	"space-playground/app/shared/infrastructure/graph/model"
	"space-playground/app/shared/infrastructure/log"

	"github.com/google/uuid"
)

// CreateAstronaut is the resolver for the createAstronaut field.
func (r *mutationResolver) CreateAstronaut(ctx context.Context, input model.NewAstronautInput) (string, error) {
	id, err := r.registerAstronautUseCase.Register(ctx, input.Name, input.IsPilot)

	if err != nil {
		log.WithError(err).Error("error at mutation[createAstronaut]")
		return "", err
	}

	return id.String(), nil
}

// CreateMission is the resolver for the createMission field.
func (r *mutationResolver) CreateMission(ctx context.Context, input model.NewMissionInput) (int, error) {
	registerMissionUseCaseInput := domain.RegisterMissionUseCaseInput{}
	registerMissionUseCaseInput.FromGql(input)

	id, err := r.registerMissionUseCase.Register(ctx, registerMissionUseCaseInput)

	if err != nil {
		log.WithError(err).Error("error at mutation[createMission]")
		return 0, err
	}

	return *id, nil
}

// GetAstronautByID is the resolver for the getAstronautById field.
func (r *queryResolver) GetAstronautByID(ctx context.Context, id string) (*model.Astronaut, error) {
	astronautUuid, err := uuid.Parse(id)

	if err != nil {
		log.WithError(err).Error("error parsing [id] argument")
		return nil, err
	}

	astronaut, err := r.astronautDetailsUsecase.ById(ctx, astronautUuid)

	if err != nil {
		log.WithError(err).Error("error at query[getAstronautById]")
		return nil, err
	}

	return &model.Astronaut{
		ID:      astronaut.ID.String(),
		Name:    astronaut.Name,
		IsPilot: astronaut.IsPilot,
	}, nil
}

// GetMissionByID is the resolver for the getMissionById field.
func (r *queryResolver) GetMissionByID(ctx context.Context, id int) (*model.Mission, error) {
	panic(fmt.Errorf("not implemented: GetMissionByID - getMissionById"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
