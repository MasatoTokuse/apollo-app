package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"user/graph/generated"
	"user/graph/model"
)

// UpsertUser is the resolver for the upsertUser field.
func (r *mutationResolver) UpsertUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	id := input.ID
	var user model.User
	user.Name = input.Name

	n := len(r.Resolver.Store)
	if id != nil {
		stored, ok := r.Resolver.Store[*id]
		if !ok {
			return nil, fmt.Errorf("not found")
		}
		_ = stored
		r.Resolver.Store[*id] = user
	} else {
		nid := strconv.Itoa(n + 1)
		user.ID = nid
		r.Resolver.Store[nid] = user
	}

	return &user, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, ok := r.Resolver.Store[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return &user, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users := make([]*model.User, 0)

	keys := make([]int, 0)
	for key := range r.Store {
		i, _ := strconv.Atoi(key)
		keys = append(keys, i)
	}
	sort.Ints(keys)

	for _, key := range keys {
		user := r.Store[strconv.Itoa(key)]
		users = append(users, &user)
	}
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
