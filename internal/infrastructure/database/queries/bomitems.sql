-- name: CreateBomItem :one
INSERT INTO bom_items (
    bom_id,
    component_id,
    quantity,
    uom,
    scrap_percent,
    operation_id,
    mask_component
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;
