import { useCallback } from "react";
import gql from "graphql-tag";

export const CREATE_TODO = gql`
  mutation CreateTodo($title: String!) {
    createTodo(input: { title: $title }) {
      id
      title
    }
  }
`;

const useCreateTodoLocal = () => {
  const createTodo = useCallback(() => {
    //
  }, []);

  return [createTodo];
};

const useCreateTodoRemote = () => {
  const createTodo = useCallback(() => {
    //
  }, []);

  return [createTodo];
};

// TODO: use ENV variable
// eslint-disable-next-line no-constant-condition
export const useCreateTodo = true ? useCreateTodoLocal : useCreateTodoRemote;
