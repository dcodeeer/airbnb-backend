CREATE SCHEMA users;

CREATE TABLE users.users (
  id            SERIAL PRIMARY KEY,
  email         VARCHAR(254) UNIQUE NOT NULL,
  password      VARCHAR(100) NOT NULL,
  phone         VARCHAR(25),
  firstname     VARCHAR(256),
  lastname      VARCHAR(256),
  patronymic    VARCHAR(256),
  date_of_birth DATE,
  photo         VARCHAR(256),
  created_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users.sessions (
  user_id    SERIAL NOT NULL,
  token      VARCHAR(128) UNIQUE NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users.recovery_keys (
  user_id    SERIAL NOT NULL,
  email      VARCHAR(254) NOT NULL,
  key        VARCHAR(128) UNIQUE NOT NULL,
  expire     TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);