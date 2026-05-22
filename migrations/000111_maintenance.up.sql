BEGIN;

CREATE TYPE maintenance_frequency_enum AS ENUM ('DAILY', 'WEEKLY', 'MONTHLY', 'CUSTOM_DAYS');
CREATE TYPE maintenance_order_status_enum AS ENUM ('PLANNED', 'IN_PROGRESS', 'DONE', 'CANCELLED');

CREATE TABLE maintenance_plans (
    id              BIGSERIAL PRIMARY KEY,
    code            BIGSERIAL UNIQUE NOT NULL,
    machine_id      BIGINT NOT NULL REFERENCES machines(id),
    work_center_id  BIGINT REFERENCES machine_types(id),
    description     TEXT NOT NULL,
    frequency       maintenance_frequency_enum NOT NULL DEFAULT 'MONTHLY',
    frequency_days  INT NOT NULL DEFAULT 30,
    estimated_hours DOUBLE PRECISION NOT NULL DEFAULT 8.0,
    last_executed_at TIMESTAMPTZ,
    next_scheduled_at TIMESTAMPTZ,
    is_active       BOOLEAN NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by      UUID NOT NULL
);

CREATE TABLE maintenance_orders (
    id              BIGSERIAL PRIMARY KEY,
    plan_id         BIGINT NOT NULL REFERENCES maintenance_plans(id),
    work_center_id  BIGINT REFERENCES machine_types(id),
    scheduled_date  DATE NOT NULL,
    estimated_hours DOUBLE PRECISION NOT NULL DEFAULT 8.0,
    actual_hours    DOUBLE PRECISION,
    status          maintenance_order_status_enum NOT NULL DEFAULT 'PLANNED',
    started_at      TIMESTAMPTZ,
    completed_at    TIMESTAMPTZ,
    notes           TEXT,
    is_active       BOOLEAN NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_maintenance_plans_machine ON maintenance_plans(machine_id);
CREATE INDEX idx_maintenance_plans_wc ON maintenance_plans(work_center_id);
CREATE INDEX idx_maintenance_orders_plan ON maintenance_orders(plan_id);
CREATE INDEX idx_maintenance_orders_date ON maintenance_orders(scheduled_date);
CREATE INDEX idx_maintenance_orders_wc ON maintenance_orders(work_center_id);
CREATE INDEX idx_maintenance_orders_status ON maintenance_orders(status);

COMMIT;
