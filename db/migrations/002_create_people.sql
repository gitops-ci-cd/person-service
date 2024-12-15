CREATE TABLE IF NOT EXISTS
  people (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    name VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

---- create above / drop below ----

DROP TABLE people;
