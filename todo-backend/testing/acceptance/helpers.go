package acceptance

import (
	"todo-backend/graph"
	"todo-backend/graph/generated"
	"todo-backend/graph/model"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/google/uuid"
)

func NewClient() (*client.Client, error) {
	res, err := graph.NewResolver()
	if err != nil {
		return nil, err
	}
	cfg := generated.Config{Resolvers: res}
	return client.New(handler.NewDefaultServer(generated.NewExecutableSchema(cfg))), nil
}

func Create(c *client.Client) model.Todo {
	title := uuid.NewString()
	var resp createTodoResponse
	c.MustPost(createTodoTemplate, &resp, client.Var("title", title))
	return resp.CreateTodo
}

func Delete(c *client.Client, id string) {
	var resp deleteTodoResponse
	c.MustPost(deleteTodoTemplate, &resp, client.Var("id", id))
}

func GetAll(c *client.Client) (todos []model.Todo) {
	var resp getManyTodoResponse

	c.MustPost(getManyTodoTemplate, &resp)
	for _, edge := range resp.Todos.Edges {
		todos = append(todos, *edge.Node)
	}

	return
}
