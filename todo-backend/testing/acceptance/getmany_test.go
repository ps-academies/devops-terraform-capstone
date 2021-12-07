//go:build acceptance
// +build acceptance

package acceptance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMany(t *testing.T) {
	c, err := NewClient()
	assert.Nil(t, err)

	expectedCount := 5

	var resp getManyTodoResponse
	var titles []string

	for i := 0; i < expectedCount; i++ {
		titles = append(titles, Create(c).Title)
	}

	c.MustPost(getManyTodoTemplate, &resp)

	t.Run("correct titles should be returned", func(t *testing.T) {
		var actual []string
		for _, edge := range resp.Todos.Edges {
			actual = append(actual, edge.Node.Title)
		}
		assert.Subset(t, actual, titles)
	})
}
