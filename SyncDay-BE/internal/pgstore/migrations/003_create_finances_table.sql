CREATE TABLE IF NOT EXISTS user_finances (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    transaction_type BIGINT NOT NULL,
    category_id BIGINT NOT NULL,
    transaction_date TIMESTAMP NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)

);

---- create above / drop below ----

DROP TABLE IF EXISTS user_finances
