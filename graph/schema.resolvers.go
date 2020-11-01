package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/cynthiawilliamsa/meetmeup/data"
	"github.com/cynthiawilliamsa/meetmeup/graph/generated"
	"github.com/cynthiawilliamsa/meetmeup/graph/model"
	"github.com/cynthiawilliamsa/meetmeup/graph/shared"
)

func (r *meetupResolver) User(ctx context.Context, obj *shared.Meetup) (*shared.User, error) {
	// return GetUserLoader(ctx).Load(obj.UserID)
	return r.DBRepo.FindByID(obj.UserID)
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*shared.Meetup, error) {
	if len(input.Name) < 2 {
		return nil, errors.New("name not long enough")
	}
	if len(input.Description) < 3 {
		return nil, errors.New("description not long enough")
	}
	meetup := &shared.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      "4",
	}
	return r.DBRepo.CreateMeetup(meetup)
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*shared.Meetup, error) {
	return r.DBRepo.GetMeetups()
}

func (r *userResolver) Meetups(ctx context.Context, obj *shared.User) ([]*shared.Meetup, error) {
	meetups := data.Meetups
	for _, m := range meetups {
		if m.ID == obj.ID {
			meetups = append(meetups, m)
		}
	}
	return meetups, nil
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
