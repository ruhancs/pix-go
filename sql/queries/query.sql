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

-- name: FindPixKeyByID :one
SELECT * FROM pix_keys WHERE id = $1 LIMIT 1;

-- name: FindAccountByID :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: FindBankByID :one
SELECT * FROM banks WHERE id = $1 LIMIT 1;