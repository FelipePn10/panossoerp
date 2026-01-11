-- name: CreateBomItem :one
INSERT INTO bom_items (
    id,
    bom_id,
    component_id,
    quantity,
    uom,
    scrap_percent,
    operation_id,
    created_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    NOW()
)
RETURNING *;
