ALTER TABLE overhead_allocations
    ADD CONSTRAINT uq_overhead_allocations_code UNIQUE (code);

CREATE INDEX IF NOT EXISTS idx_overhead_allocations_code
    ON overhead_allocations(code);

CREATE INDEX IF NOT EXISTS idx_overhead_allocations_plan_account_id
    ON overhead_allocations(plan_account_id);

CREATE INDEX IF NOT EXISTS idx_overhead_allocations_base_id
    ON overhead_allocations(base_id);