BEGIN;

-- CRP: carga calculada por CT por dia
CREATE TABLE capacity_requirements (
    id              BIGSERIAL PRIMARY KEY,
    plan_code       BIGINT NOT NULL,
    work_center_id  BIGINT NOT NULL REFERENCES machine_types(id),
    req_date        DATE NOT NULL,
    required_hours  NUMERIC(15,4) NOT NULL DEFAULT 0,
    available_hours NUMERIC(15,4) NOT NULL DEFAULT 0,
    load_pct        NUMERIC(6,2) GENERATED ALWAYS AS (
        CASE WHEN available_hours > 0 THEN (required_hours / available_hours) * 100 ELSE 0 END
    ) STORED,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (plan_code, work_center_id, req_date)
);

CREATE INDEX idx_crp_plan ON capacity_requirements(plan_code);
CREATE INDEX idx_crp_wc_date ON capacity_requirements(work_center_id, req_date);
CREATE INDEX idx_crp_overloaded ON capacity_requirements(plan_code) WHERE load_pct > 100;

COMMIT;
