package graph

import "user/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Store map[string]model.User
}

func NewResolver() *Resolver {
	return &Resolver{
		Store: make(map[string]model.User),
	}
}
