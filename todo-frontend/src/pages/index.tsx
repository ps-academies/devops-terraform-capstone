import React, { useMemo } from "react";

import { useSetVisibiltyFilter, useUpdateTodo } from "../operations";
import {
  useGetTodosQuery,
  useGetVisibilityFilterQuery,
  VisibilityFilterOptions,
} from "../state";

import { Header } from "../components/Header";
import {
  Footer,
  FooterClearCompletedButton,
  FooterActiveCount,
  FooterVisibiltyFilters,
} from "../components/Footer";
import { Main } from "../components/Main";
import { TodoList, TodoListItem } from "../components/TodoList";

const IndexRoute: React.FC = () => {
  const { data: todosConnection } = useGetTodosQuery();
  const {
    data: { visibilityFilter: activeFilter },
  } = useGetVisibilityFilterQuery();

  const [setVisibilityFilter] = useSetVisibiltyFilter();
  const [updateTodo] = useUpdateTodo();

  const [completed, active] = useMemo(() => {
    const { edges } = todosConnection.todos;
    return edges.reduce(
      ([a, b], edge) => {
        return edge.node.completed ? [[...a, edge], b] : [a, [...b, edge]];
      },
      [[], []]
    );
  }, [todosConnection]);

  const filtered = useMemo(() => {
    if (activeFilter.id === VisibilityFilterOptions["SHOW_ACTIVE"].id) {
      return active;
    }

    if (activeFilter.id === VisibilityFilterOptions["SHOW_COMPLETED"].id) {
      return completed;
    }

    return todosConnection.todos.edges;
  }, [activeFilter, completed, active]);

  const hasSelectedItems = filtered.length > 0;

  return (
    <>
      <div className="todoapp">
        <Header />

        {hasSelectedItems && (
          <Main>
            {completed.length > 0 && (
              <>
                <input id="toggle-all" className="toggle-all" type="checkbox" />
                <label htmlFor="toggle-all">Mark all as complete</label>
              </>
            )}

            <TodoList>
              {filtered.map(({ node }) => (
                <TodoListItem key={node.id} todo={node} onUpdate={updateTodo} />
              ))}
            </TodoList>
          </Main>
        )}

        <Footer>
          {active.length > 0 && <FooterActiveCount count={active.length} />}

          <FooterVisibiltyFilters
            activeFilter={activeFilter}
            setVisibilityFilter={setVisibilityFilter}
          />

          {completed.length > 0 && <FooterClearCompletedButton />}
        </Footer>
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
