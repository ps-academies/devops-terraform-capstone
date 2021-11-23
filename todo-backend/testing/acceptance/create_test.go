//go:build acceptance
// +build acceptance

package acceptance

import (
	"testing"
	"todo-backend/graph"
	"todo-backend/graph/generated"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
)

func TestCreateTodo(t *testing.T) {
	cfg := generated.Config{Resolvers: graph.NewResolver()}
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(cfg)))

	var resp createTodoResponse
	title := "hello world"

	c.MustPost(createTodoTemplate, &resp, client.Var("title", title))

	t.Run("Correct title should be returned", func(t *testing.T) {
		if actual := resp.CreateTodo.Title; actual != title {
			t.Errorf("expected: '%s' | actual: '%s'", title, actual)
		}
	})

	t.Run("ID should not be empty", func(t *testing.T) {
		if resp.CreateTodo.ID == "" {
			t.Error("expected ID but it was empty")
		}
	})
}
