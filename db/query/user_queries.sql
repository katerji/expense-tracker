-- name: FetchUserByEmailQuery :one
SELECT id, email, first_name, last_name, password FROM user
WHERE email = ?;

-- name: InsertUserQuery :exec
INSERT INTO user (email, first_name, last_name, password) VALUES (?, ?, ?, ?);

-- name: FetchUserAccount :one
SELECT id, name, user_id
from account
where user_id = ?;

-- name: InsertAccount :exec
INSERT INTO account (name, user_id) VALUES (?, ?);