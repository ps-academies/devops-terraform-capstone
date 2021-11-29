import { ApolloClient } from "@apollo/client";
import { cache } from "./cache";

export const client = new ApolloClient({
  cache,
  connectToDevTools: true,
});
