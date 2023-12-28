-- migrate:up
CREATE TABLE users (
  id uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
  name varchar(32) NOT NULL,
  email varchar(32) UNIQUE NOT NULL,
  password varchar(255) NOT NULL,
  verified boolean NOT NULL DEFAULT false,
  last_seen_at timestamptz NOT NULL DEFAULT now(),
  created_at timestamptz NOT NULL DEFAULT now()
);
-- migrate:down
DROP TABLE users;