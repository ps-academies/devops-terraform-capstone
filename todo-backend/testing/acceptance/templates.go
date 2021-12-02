package acceptance

import "todo-backend/graph/model"

var createTodoTemplate = `
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

var getManyTodoTemplate = `
query GetManyTodos {
	todos {
		edges {
			node {
				id
				title
				completed
			}
		}
	}
}`

type getManyTodoResponse struct {
	Todos model.TodoConnection
}

type deleteTodoResponse struct {
	DeleteTodo model.Todo
}

var deleteTodoTemplate = `
mutation DeleteTodoTest($id: ID!) {
  deleteTodo(
    id: $id
  ) {
    id
    title
  }
}`
