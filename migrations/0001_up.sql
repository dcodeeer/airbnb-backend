-- USERS

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

-- BOOKING

CREATE SCHEMA booking;

CREATE TABLE booking.statuses (
  id   SERIAL PRIMARY KEY,
  name VARCHAR(535) NOT NULL
);
-- 'pending', 'booked', 'finished'

CREATE TABLE booking.booking (
  id          SERIAL PRIMARY KEY,
  estate_id   SERIAL NOT NULL,
  customer_id SERIAL NOT NULL,
  status      SERIAL NOT NULL,
  date_from   DATE NOT NULL,
  date_to     DATE NOT NULL,
  total       SERIAL NOT NULL,
  created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ESTATES

CREATE SCHEMA estates;

CREATE TABLE estates.estates (
  id          SERIAL PRIMARY KEY,
  owner_id    SERIAL NOT NULL,
  price_night INTEGER NOT NULL,
  price_week  INTEGER NOT NULL,

  area        INTEGER NOT NULL,
  rooms       INTEGER NOT NULL,
  showers     INTEGER NOT NULL,
  baby_rooms  INTEGER NOT NULL,

  created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE estates.images (
  id        SERIAL PRIMARY KEY,
  estate_id SERIAL NOT NULL,
  path      VARCHAR(535) NOT NULL,
  name      VARCHAR(535) NOT NULL,
  priority  SERIAL NOT NULL
);

CREATE TABLE estates.addresses (
  id       SERIAL PRIMARY KEY,
  number   SERIAL NOT NULL,
  street   VARCHAR(535) NOT NULL,
  city     VARCHAR(535) NOT NULL,
  district VARCHAR(535) NOT NULL
);

CREATE TABLE estates.amenities (
  id   SERIAL PRIMARY KEY,
  name VARCHAR(535) NOT NULL
);

CREATE TABLE estates.reviews (
  id         SERIAL PRIMARY KEY,
  user_id    SERIAL NOT NULL,
  estate_id  SERIAL NOT NULL,
  content    TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);