import gql from "graphql-tag";

export const GET_TODOS = gql`
  query GetTodos {
    todos @client {
      edges {
        cursor
        node {
          id
          title
        }
      }
      pageInfo {
        endCursor
        hasNextPage
        startCursor
      }
    }
  }
`;
