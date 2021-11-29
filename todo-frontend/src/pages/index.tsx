import React from "react";

import { useGetTodosQuery } from "../state";

import { Header } from "../components/Header";
import { Main } from "../components/Main";
import { TodoList, TodoListItem } from "../components/TodoList";

const IndexRoute: React.FC = () => {
  const { data } = useGetTodosQuery();

  const hasTodos = data?.todos.edges && data?.todos.edges.length > 0;

  return (
    <>
      <div className="todoapp">
        <Header />

        {hasTodos && (
          <Main>
            <input id="toggle-all" className="toggle-all" type="checkbox" />
            <label htmlFor="toggle-all">Mark all as complete</label>

            <TodoList>
              {data.todos?.edges.map(({ node }) => (
                <TodoListItem key={node.id} todo={node} />
              ))}
            </TodoList>
          </Main>
        )}

        {hasTodos && (
          <footer className="footer">
            <span className="todo-count">
              <strong>0</strong> item left
            </span>

            <ul className="filters">
              <li>
                <a className="selected" href="#/">
                  All
                </a>
              </li>
              <li>
                <a href="#/active">Active</a>
              </li>
              <li>
                <a href="#/completed">Completed</a>
              </li>
            </ul>
            <button className="clear-completed">Clear completed</button>
          </footer>
        )}
      </div>

      <footer className="info">
        <p>Double-click to edit a todo</p>
        <p>
          Template by <a href="http://sindresorhus.com">Sindre Sorhus</a>
        </p>
        <p>
          Created by <a href="http://todomvc.com">you</a>
        </p>
        <p>
          Part of <a href="http://todomvc.com">TodoMVC</a>
        </p>
      </footer>
    </>
  );
};

export default IndexRoute;
