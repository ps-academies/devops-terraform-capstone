package graph

import "todo-backend/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos *model.TodoConnection
}

func NewResolver() *Resolver {
	return &Resolver{
		todos: &model.TodoConnection{
			Edges: []*model.TodoEdge{},
		},
	}
}
