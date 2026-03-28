-- name: CreateWarehouse :one
INSERT INTO warehouse (
    code,
    description,
    location,
    type,
    disposition,
    reservations_allowed,
    created_by
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
) RETURNING
    code,
    description,
    location,
    type,
    disposition,
    reservations_allowed,
    created_by,
    created_at;
