package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"todo-backend/graph/generated"
	"todo-backend/graph/model"

	"github.com/google/uuid"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	next := &model.Todo{
		ID:    uuid.New().String(),
		Title: input.Title,
	}
	nextEdge := model.TodoEdge{Node: next}
	r.todos.Edges = append(r.todos.Edges, &nextEdge)

	return next, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*model.Todo, error) {
	found := -1
	var todo *model.Todo
	for i, edge := range r.todos.Edges {
		if edge.Node.ID == id {
			todo = edge.Node
			found = i
			break
		}
	}
	if found >= 0 {
		r.todos.Edges = append(r.todos.Edges[:found], r.todos.Edges[found+1:]...)
		return todo, nil
	}

	return nil, fmt.Errorf("todo with id '%s' could not be found", id)
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context, first *int, after *string) (*model.TodoConnection, error) {
	return r.todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
