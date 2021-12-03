import { useCallback } from "react";
import gql from "graphql-tag";

import { todosVar, useDeleteTodoMutation } from "../../../state";

export const DELETE_TODO = gql`
  mutation DeleteTodo($id: ID!) {
    deleteTodo(id: $id) {
      id
    }
  }
`;

const useDeleteTodoLocal = () => {
  const deleteTodo = useCallback<(id: string) => void>((id) => {
    const prev = todosVar();
    const { edges } = prev;

    const nextEdges = edges.filter((edge) => edge.node.id !== id, []);

    todosVar({ ...prev, edges: nextEdges });
  }, []);

  return [deleteTodo];
};

const useDeleteTodoRemote = () => {
  const [deleteTodoMutation] = useDeleteTodoMutation();

  const deleteTodo = useCallback<(id: string) => void>((id) => {
      deleteTodoMutation({
        refetchQueries: ['GetTodos'],
        variables: { id },
      });
  }, []);

  return [deleteTodo];
};

// TODO: use ENV variable
// eslint-disable-next-line no-constant-condition
export const useDeleteTodo = false ? useDeleteTodoLocal : useDeleteTodoRemote;
