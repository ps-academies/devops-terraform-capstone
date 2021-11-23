package acceptance

import "todo-backend/graph/model"

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

var getManyTodoTemplate string = `
query GetManyTodos {
	todos {
		edges {
			node {
				title
			}
		}
	}
}`

type getManyTodoResponse struct {
	Todos model.TodoConnection
}
