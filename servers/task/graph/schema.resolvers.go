package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"task/graph/generated"
	"task/graph/model"
)

// UpsertTask is the resolver for the upsertTask field.
func (r *mutationResolver) UpsertTask(ctx context.Context, input model.TaskInput) (*model.Task, error) {
	id := input.ID
	var task model.Task
	task.Text = input.Text
	task.User = &model.User{
		ID: input.UserID,
	}

	n := len(r.Resolver.Store)
	if id != nil {
		stored, ok := r.Resolver.Store[*id]
		if !ok {
			return nil, fmt.Errorf("not found")
		}
		_ = stored
		task.ID = *id
		r.Resolver.Store[*id] = task
	} else {
		nid := strconv.Itoa(n + 1)
		task.ID = nid
		r.Resolver.Store[nid] = task
	}

	return &task, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	tasks := make([]*model.Task, 0)

	keys := make([]int, 0)
	for key := range r.Resolver.Store {
		i, _ := strconv.Atoi(key)
		keys = append(keys, i)
	}
	sort.Ints(keys)

	for _, key := range keys {
		task := r.Resolver.Store[strconv.Itoa(key)]
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
