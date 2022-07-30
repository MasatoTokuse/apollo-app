package graph

import "task/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Store map[string]model.Task
}

func NewResolver() *Resolver {
	return &Resolver{
		Store: make(map[string]model.Task),
	}
}
