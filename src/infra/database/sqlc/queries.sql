-- name: GetClient :one
SELECT * FROM CLIENTS WHERE ID = $1;


-- name: UpdateBalance :exec
UPDATE CLIENTS SET BALANCE = $1 WHERE ID = $2;


-- name: CreateTransaction :exec
INSERT INTO TRANSACTIONS (VALUE, TYPE, DESCRIPTION, CLIENT_ID)
VALUES ($1, $2, $3, $4);

-- name: FindLastTransactionsByClient :many
SELECT * FROM TRANSACTIONS WHERE CLIENT_ID = $1 ORDER BY ID DESC LIMIT 10;
