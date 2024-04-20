-- name: CreateAccount :one
INSERT INTO Account (owner,balance,currency,country_code) Values ($1 , $2, $3, $4) RETURNING *;

-- name: GetAccount :one
SELECT * FROM Account WHERE id = $1 LIMIT 1;
