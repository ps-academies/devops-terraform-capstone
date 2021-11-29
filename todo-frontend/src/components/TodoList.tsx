import React from "react";
import { Todo } from "../state";

export const TodoList: React.FC = (props) => (
  <ul className="todo-list" {...props} />
);

interface TodoListItemProps {
  todo: Todo;
}

export const TodoListItem: React.FC<TodoListItemProps> = (props) => {
  const { children, todo, ...rest } = props;

  return (
    <li className="completed" {...rest}>
      <div className="view">
        <input className="toggle" type="checkbox" checked />
        <label>{todo.title}</label>
        <button className="destroy"></button>
      </div>

      <input className="edit" value="Create a TodoMVC template" />
    </li>
  );
};
