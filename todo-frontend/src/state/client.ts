import { ApolloClient } from "@apollo/client";
import { cache } from "./cache";

import { useRemote } from '../env'

export const client = new ApolloClient({
  cache,
  connectToDevTools: true,
  ...(useRemote() && {
    uri: process.env.REMOTE_SCHEMA_URL
  })
});


