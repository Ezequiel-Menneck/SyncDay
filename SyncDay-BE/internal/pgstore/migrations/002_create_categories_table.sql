CREATE TABLE IF NOT EXISTS categories(
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(30) NOT NULL,
    description VARCHAR(30) NULL
);
---- create above / drop below ----

DROP TABLE IF EXISTS categories