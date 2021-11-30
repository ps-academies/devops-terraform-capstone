import { useCallback } from "react";
import gql from "graphql-tag";

import {
  UpdatedTodo,
  todosVar,
  useUpdateTodoMutation,
  TodoEdge,
} from "../../../state";

export const UPDATE_TODO = gql`
  mutation UpdateTodo($id: ID!, $input: UpdatedTodo!) {
    updateTodo(id: $id, input: $input) {
      id
      title
      completed
    }
  }
`;

const useUpdateTodoLocal = () => {
  const updateTodo = useCallback<(id: string, input: UpdatedTodo) => void>(
    (id, input) => {
      const prev = todosVar();
      const { edges } = prev;

      const nextEdges = edges.reduce((acc: TodoEdge[], edge) => {
        if (edge.node.id === id)
          return acc.concat({ ...edge, node: { ...edge.node, ...input } });

        return acc.concat(edge);
      }, []);

      todosVar({ ...prev, edges: nextEdges });
    },
    []
  );

  return [updateTodo];
};

const useUpdateTodoRemote = () => {
  const [updateTodoMutation] = useUpdateTodoMutation();

  const updateTodo = useCallback<(id: string, input: UpdatedTodo) => void>(
    (id, input) => {
      updateTodoMutation({ variables: { id, input } });
    },
    []
  );

  return [updateTodo];
};

// TODO: use ENV variable
// eslint-disable-next-line no-constant-condition
export const useUpdateTodo = true ? useUpdateTodoLocal : useUpdateTodoRemote;
