//go:build acceptance
// +build acceptance

package acceptance

import (
	"fmt"
	"strconv"
	"testing"
	"todo-backend/graph"
	"todo-backend/graph/generated"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
)

func TestGetMany(t *testing.T) {
	cfg := generated.Config{Resolvers: graph.NewResolver()}
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(cfg)))

	expectedCount := 5

	var resp getManyTodoResponse
	var titles []string

	for i := 0; i < expectedCount; i++ {
		title := strconv.Itoa(i)
		titles = append(titles, title)
		c.MustPost(createTodoTemplate, &createTodoResponse{}, client.Var("title", title))
	}

	c.MustPost(getManyTodoTemplate, &resp)

	t.Run(fmt.Sprintf("%d items should be returned", expectedCount), func(t *testing.T) {
		if actual := len(resp.Todos.Edges); actual != expectedCount {
			t.Errorf("expected: '%d' | actual: '%d'", expectedCount, actual)
		}
	})
}
