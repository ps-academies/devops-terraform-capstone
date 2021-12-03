declare global {
  namespace NodeJS {
    interface ProcessEnv {
      REMOTE_SCHEMA_URL?: string
    }
  }
}

export {}
