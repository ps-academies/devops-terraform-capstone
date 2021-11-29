import gql from "graphql-tag";

export const CREATE_TODO = gql`
  mutation CreateTodo($title: String!) {
    createTodo(input: { title: $title }) {
      id
      title
    }
  }
`;
