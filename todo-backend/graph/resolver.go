package graph

import "todo-backend/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos *model.TodoConnection
}

func NewResolver() *Resolver {
	defaultTotalCount := 0
	defaultHasNextPage := false
	return &Resolver{
		todos: &model.TodoConnection{
			PageInfo: &model.PageInfo{
				StartCursor: "cur",
				EndCursor:   "cur",
				HasNextPage: &defaultHasNextPage,
			},
			Edges:      []*model.TodoEdge{},
			TotalCount: &defaultTotalCount,
		},
	}
}
