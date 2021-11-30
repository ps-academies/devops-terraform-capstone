//go:build acceptance
// +build acceptance

package acceptance

import (
	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, resp.CreateTodo.Title, title)
	})

	t.Run("ID should not be empty", func(t *testing.T) {
		assert.NotEmpty(t, resp.CreateTodo.ID, "expected a non-empty ID")
	})
}
