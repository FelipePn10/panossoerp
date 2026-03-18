-- name: CreateWarehouse :one
INSERT INTO warehouse (
    code,
    description,
    location,
    type,
    disposition,
    reservation_allowed,
    created_by
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
) RETURNING *;

-- name: ExistsWarehouseByCode :one
SELECT EXISTS (
    SELECT 1
    FROM warehouse
    WHERE code = $1
);