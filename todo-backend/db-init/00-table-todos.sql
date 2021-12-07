BEGIN;
CREATE TABLE IF NOT EXISTS todos (
  id varchar PRIMARY KEY,
  title varchar,
  completed boolean,
  created date DEFAULT now(),
  updated date DEFAULT now()
);
END;
