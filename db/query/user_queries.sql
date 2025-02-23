-- name: FetchUserByEmailQuery :one
SELECT id, email, first_name, last_name, password FROM user
WHERE email = ?;

-- name: FetchUserByID :one
SELECT id, email, first_name, last_name, password FROM user
WHERE id = ?;

-- name: InsertUserQuery :exec
INSERT INTO user (email, first_name, last_name, password) VALUES (?, ?, ?, ?);

-- name: FetchUserAccount :one
SELECT id, name, user_id
from account
where user_id = ?;

-- name: FetchAccountByID :one
SELECT id, name, user_id
from account
where id = ?;

-- name: InsertAccount :exec
INSERT INTO account (name, user_id) VALUES (?, ?);