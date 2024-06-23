-- name: CreateAccount :one
INSERT INTO account (
    owner,
    balance,
    currency
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: ListAccounts :many
SELECT * FROM account
ORDER BY owner
LIMIT $1 OFFSET $2;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: UpdateAccount :one
UPDATE account
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: GetAccountForUpdate :one
SELECT * FROM account
WHERE id = $1
FOR NO KEY UPDATE;

-- name: AddAccountBalance :one
UPDATE account
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;