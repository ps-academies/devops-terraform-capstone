//go:build acceptance
// +build acceptance

package acceptance

import (
	"fmt"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"todo-backend/graph"
	"todo-backend/graph/generated"
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

	t.Run("correct titles should be returned", func(t *testing.T) {
		var actual []string
		for _, edge := range resp.Todos.Edges {
			actual = append(actual, edge.Node.Title)
		}
		assert.ElementsMatch(t, titles, actual)
	})
}
