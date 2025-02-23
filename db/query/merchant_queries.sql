-- name: FetchMerchantByNameQuery :one
SELECT m.id, m.name, m.type_id, mt.type as merchant_type
FROM merchant m
         JOIN merchant_type mt ON m.type_id = mt.id
WHERE m.name = ?;

-- name: FetchMerchantByIDQuery :one
SELECT m.id, m.name, m.type_id, mt.type as merchant_type
FROM merchant m
         JOIN merchant_type mt ON m.type_id = mt.id
WHERE m.id = ?;

-- name: FetchMerchantTypeQuery :one
SELECT id, type
FROM merchant_type
WHERE type = ?;

-- name: InsertMerchantTypeQuery :execresult
INSERT INTO merchant_type (type) VALUES (?);

-- name: InsertMerchantQuery :execresult
INSERT INTO merchant (name, type_id) VALUES (?, ?);

-- name: FetchMerchantQuery :one
SELECT m.id, m.name, m.type_id, mt.type as merchant_type
FROM merchant m
         JOIN merchant_type mt ON m.type_id = mt.id
WHERE type = ?;