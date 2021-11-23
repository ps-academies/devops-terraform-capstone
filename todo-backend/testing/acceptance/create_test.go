//go:build acceptance
// +build acceptance

package testing

import (
	"testing"
	"todo-backend/graph"
	"todo-backend/graph/generated"
	"todo-backend/graph/model"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
)

var createTodoTemplate string = `
mutation CreateTodoTest($title: String!) {
  createTodo(
    input: { title: $title }
  ) {
    id
    title
  }
}`

type createTodoResponse struct {
	CreateTodo model.Todo
}

func TestCreateTodo(t *testing.T) {
	cfg := generated.Config{Resolvers: &graph.Resolver{}}
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
