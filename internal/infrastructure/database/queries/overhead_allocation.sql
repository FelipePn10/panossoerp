-- name: CreateOverheadAllocation :one
INSERT INTO overhead_allocations (cost_center_code, plan_account_code, account_code, period_start, period_end, allocation_type, base_code, created_by)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING *;

-- name: AddAllocationTarget :one
INSERT INTO overhead_allocation_targets (overhead_code, cost_center_code, percentage, amount)
VALUES ($1, $2, $3, $4)
    RETURNING *;

-- name: GetOverheadAllocationByCode :one
SELECT * FROM overhead_allocations WHERE code = $1;

-- name: GetAllocationTargets :many
SELECT * FROM overhead_allocation_targets WHERE overhead_code = $1;

-- name: ListOverheadAllocations :many
SELECT * FROM overhead_allocations ORDER BY created_at DESC;

-- name: ListOverheadAllocationsByCostCenter :many
SELECT * FROM overhead_allocations WHERE cost_center_code = $1;

-- name: DeleteOverheadAllocation :exec
DELETE FROM overhead_allocations WHERE code = $1;

-- name: DeleteAllocationTargets :exec
DELETE FROM overhead_allocation_targets WHERE overhead_code = $1;
