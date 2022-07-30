package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"sort"
	"strconv"
	"task/graph/generated"
	"task/graph/model"
)

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	tasks := make([]*model.Task, 0)

	keys := make([]int, 0)
	for key := range r.Resolver.Store {
		i, _ := strconv.Atoi(key)
		keys = append(keys, i)
	}
	sort.Ints(keys)

	for _, key := range keys {
		task := r.Resolver.Store[strconv.Itoa(key)]
		if task.User.ID == id {
			tasks = append(tasks, &task)
		}
	}

	return &model.User{
		ID:    id,
		Tasks: tasks,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
