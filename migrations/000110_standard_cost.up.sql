BEGIN;

-- Custo-padrão calculado por item (+ máscara)
CREATE TABLE item_standard_costs (
    id              BIGSERIAL PRIMARY KEY,
    item_code       BIGINT NOT NULL REFERENCES items(code),
    mask            VARCHAR(200) NOT NULL DEFAULT '',
    material_cost   NUMERIC(20,6) NOT NULL DEFAULT 0,
    labor_cost      NUMERIC(20,6) NOT NULL DEFAULT 0,  -- custo de mão-de-obra (tempo × custo CT)
    overhead_cost   NUMERIC(20,6) NOT NULL DEFAULT 0,
    total_cost      NUMERIC(20,6) GENERATED ALWAYS AS (material_cost + labor_cost + overhead_cost) STORED,
    currency        VARCHAR(3) NOT NULL DEFAULT 'BRL',
    calculated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    calculated_by   UUID NOT NULL,
    UNIQUE (item_code, mask)
);

CREATE INDEX idx_isc_item ON item_standard_costs(item_code);

-- Custo unitário por CT (hora-máquina + mão-de-obra)
CREATE TABLE work_center_costs (
    id                  BIGSERIAL PRIMARY KEY,
    work_center_id      BIGINT NOT NULL REFERENCES machine_types(id) UNIQUE,
    cost_per_hour       NUMERIC(20,6) NOT NULL DEFAULT 0,
    currency            VARCHAR(3) NOT NULL DEFAULT 'BRL',
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by          UUID NOT NULL
);

-- Log de rolagem (histórico de cada cálculo)
CREATE TABLE cost_rollup_log (
    id              BIGSERIAL PRIMARY KEY,
    item_code       BIGINT NOT NULL,
    mask            VARCHAR(200) NOT NULL DEFAULT '',
    bom_level       INT NOT NULL,
    material_cost   NUMERIC(20,6) NOT NULL,
    labor_cost      NUMERIC(20,6) NOT NULL,
    overhead_cost   NUMERIC(20,6) NOT NULL,
    run_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_crl_item ON cost_rollup_log(item_code, run_at);

-- Custo de compra por item (item comprado = referência para custo de material)
CREATE TABLE item_purchase_costs (
    id          BIGSERIAL PRIMARY KEY,
    item_code   BIGINT NOT NULL REFERENCES items(code) UNIQUE,
    unit_cost   NUMERIC(20,6) NOT NULL DEFAULT 0,
    currency    VARCHAR(3) NOT NULL DEFAULT 'BRL',
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by  UUID NOT NULL
);

COMMIT;
