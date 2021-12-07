package graph

import (
	"context"
	"todo-backend/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	database *db.Database
}

func NewResolver() (*Resolver, error) {
	connString := "postgres://todo_admin:postgres@127.0.0.1:5432/todos"

	ctx := context.Background()
	conn, err := db.NewDatabase(ctx, connString)
	if err != nil {
		return nil, err
	}

	return &Resolver{database: conn}, nil
}
