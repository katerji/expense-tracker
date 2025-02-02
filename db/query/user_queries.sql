-- name: FetchUserByEmailQuery :one
SELECT id, email, first_name, last_name FROM user
WHERE email = ?;

-- name: InsertUserQuery :exec
INSERT INTO user (email, first_name, last_name, password) VALUES (?, ?, ?, ?)