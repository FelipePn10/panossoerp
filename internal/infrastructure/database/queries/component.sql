-- name: CreateComponent :one
INSERT INTO components (
    name,
    group_code,
    warehouse,
    created_by
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;

-- name: ExistsComponentByCode :one
SELECT *
FROM components
WHERE code = $1;
