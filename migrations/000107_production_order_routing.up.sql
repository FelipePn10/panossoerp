BEGIN;

-- Link OP ao roteiro
ALTER TABLE production_orders ADD COLUMN IF NOT EXISTS route_id BIGINT REFERENCES manufacturing_routes(id);

-- Operações da OP (espelho das route_operations no momento da abertura)
CREATE TABLE production_order_operations (
    id                  BIGSERIAL PRIMARY KEY,
    production_order_id BIGINT NOT NULL REFERENCES production_orders(id) ON DELETE CASCADE,
    route_operation_id  BIGINT REFERENCES route_operations(id),
    sequence            SMALLINT NOT NULL,
    operation_name      VARCHAR(100) NOT NULL,
    work_center_id      BIGINT REFERENCES machine_types(id),
    planned_hours       NUMERIC(15,4) NOT NULL DEFAULT 0,
    setup_hours         NUMERIC(15,4) NOT NULL DEFAULT 0,
    actual_hours        NUMERIC(15,4) NOT NULL DEFAULT 0,
    status              VARCHAR(20) NOT NULL DEFAULT 'PENDING', -- PENDING, IN_PROGRESS, DONE, SKIPPED
    started_at          TIMESTAMPTZ,
    completed_at        TIMESTAMPTZ,
    notes               TEXT,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (production_order_id, sequence)
);

CREATE INDEX idx_poo_order ON production_order_operations(production_order_id);
CREATE INDEX idx_poo_wc ON production_order_operations(work_center_id);

-- Apontamento por operação (complementa production_appointments)
ALTER TABLE production_appointments
    ADD COLUMN IF NOT EXISTS operation_id BIGINT REFERENCES production_order_operations(id);

-- Sequenciamento APS
CREATE TABLE production_sequences (
    id                  BIGSERIAL PRIMARY KEY,
    production_order_id BIGINT NOT NULL REFERENCES production_orders(id) ON DELETE CASCADE,
    operation_id        BIGINT REFERENCES production_order_operations(id),
    work_center_id      BIGINT NOT NULL REFERENCES machine_types(id),
    sequence_position   INT NOT NULL,
    scheduled_start     TIMESTAMPTZ NOT NULL,
    scheduled_end       TIMESTAMPTZ NOT NULL,
    status              VARCHAR(20) NOT NULL DEFAULT 'SCHEDULED', -- SCHEDULED, CONFIRMED, DONE
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_ps_order ON production_sequences(production_order_id);
CREATE INDEX idx_ps_wc ON production_sequences(work_center_id);
CREATE INDEX idx_ps_start ON production_sequences(scheduled_start);

COMMIT;
