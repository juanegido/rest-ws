DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP TABLE IF EXISTS products;

CREATE TABLE products (
  id VARCHAR(255) NOT NULL PRIMARY KEY,
  reference VARCHAR(255) NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  price FLOAT NOT NULL,
  quantity INTEGER NOT NULL DEFAULT 0
);