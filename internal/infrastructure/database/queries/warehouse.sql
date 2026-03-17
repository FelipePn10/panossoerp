-- name: CreateWarehouse :one
INSERT INTO warehouse (
    name,
    description,
    code,
    types,
    created_by
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: ExistsWarehouseByName :one
SELECT *
FROM questions
WHERE name = $1;
