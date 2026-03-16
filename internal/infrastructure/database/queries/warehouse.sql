-- name: CreateWarehouse :one
INSERT INTO questions (
    name,
    description,
    code,
    types,
    createdby
) VALUES (
    $1,
    $2
    $3,
    $4
) RETURNING *;

-- name: ExistsWarehouseByName :one
SELECT *
FROM questions
WHERE name = $1;
