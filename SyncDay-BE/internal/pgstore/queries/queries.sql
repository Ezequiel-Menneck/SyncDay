-- name: CreateUser :one
INSERT INTO users (name, last_name, email, cellphone, base_salary)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: UpdateUserBaseSalary :exec
UPDATE users SET base_salary = $1 WHERE id = $2;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1;