CREATE TABLE IF NOT EXISTS  users (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    last_name VARCHAR(60) NOT NULL,
    email VARCHAR(60) NOT NULL,
    cellphone VARCHAR(20) NULL,
    base_salary NUMERIC(10, 2) NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS users