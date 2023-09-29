-- name: CreateAccount :exec
INSERT INTO accounts 
    (id, owner_name, balance, number, bank_id) VALUES($1,$2,$3,$4,$5);

-- name: CreateBank :exec
INSERT INTO banks 
    (id, code, name, created_at)
    VALUES($1,$2,$3,$4);

-- name: CreatePixKey :exec
INSERT INTO pix_keys 
    (id, kind, key, status, account_id, created_at)
    VALUES($1,$2,$3,$4,$5,$6);

-- name: CreateTransaction :exec
INSERT INTO transactions 
    (id, account_from_id, pix_key_id, amount, status, description, created_at)
    VALUES($1,$2,$3,$4,$5,$6,$7);

-- name: FindPixKeyByID :one
SELECT * FROM pix_keys WHERE id = $1 LIMIT 1;

-- name: FindAccountByID :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: FindBankByID :one
SELECT * FROM banks WHERE id = $1 LIMIT 1;

-- name: FindTransactionByID :one
SELECT * FROM transactions WHERE id = $1 LIMIT 1;

-- name: GetTransactionForUpdate :one
SELECT * FROM transactions WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: UpdateTransactionStatus :one
UPDATE transactions SET status = $2 WHERE id = $1 RETURNING *;