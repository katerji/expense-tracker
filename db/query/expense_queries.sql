-- name: FetchExpenseByIDQuery :one
SELECT e.id,
       e.amount,
       e.currency,
       e.time_of_purchase,
       e.description,
       e.merchant_id,
       e.account_id,
       m.name  as merchant_name,
       mt.id   as merchant_type_id,
       mt.type as merchant_type
FROM expense e
         JOIN merchant m ON e.merchant_id = m.id
         JOIN merchant_type mt ON mt.id = m.type_id
WHERE e.id = ?;

-- name: InsertExpenseQuery :execresult
INSERT INTO expense (amount, currency, time_of_purchase, description, merchant_id, account_id)
VALUES (?, ?, ?, ?, ?, ?);
