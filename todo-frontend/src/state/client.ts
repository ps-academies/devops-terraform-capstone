import { ApolloClient } from "@apollo/client";
import { cache } from "./cache";


export const client = new ApolloClient({
  cache,
  connectToDevTools: true,
  uri: "http://localhost:8080/query",

//  uri: process.env.GATSBY_REMOTE_SCHEMA_URL
});


