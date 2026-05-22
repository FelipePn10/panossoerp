BEGIN;

CREATE TYPE inspection_point_type AS ENUM ('RECEBIMENTO', 'PROCESSO', 'EXPEDICAO');
CREATE TYPE inspection_result_enum AS ENUM ('APROVADO', 'REJEITADO', 'CONDICIONAL', 'PENDENTE');
CREATE TYPE nc_severity_enum AS ENUM ('CRITICA', 'MAIOR', 'MENOR', 'OBSERVACAO');
CREATE TYPE nc_disposition_enum AS ENUM ('SUCATA', 'RETRABALHO', 'APROVADO_CONDICIONAL', 'DEVOLVIDO');

-- Plano de inspeção por item (+ operação do roteiro opcional)
CREATE TABLE inspection_plans (
    id                   BIGSERIAL PRIMARY KEY,
    item_code            BIGINT NOT NULL REFERENCES items(code),
    route_operation_id   BIGINT REFERENCES route_operations(id),
    point_type           inspection_point_type NOT NULL,
    description          VARCHAR(200) NOT NULL,
    sample_size          NUMERIC(10,4) NOT NULL DEFAULT 1,
    acceptance_level     NUMERIC(5,2) NOT NULL DEFAULT 0,  -- AQL %
    instructions         TEXT,
    is_active            BOOLEAN NOT NULL DEFAULT TRUE,
    created_at           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by           UUID NOT NULL
);

CREATE INDEX idx_ip_item ON inspection_plans(item_code);
CREATE INDEX idx_ip_route_op ON inspection_plans(route_operation_id);

-- Características a inspecionar dentro do plano
CREATE TABLE inspection_plan_characteristics (
    id              BIGSERIAL PRIMARY KEY,
    plan_id         BIGINT NOT NULL REFERENCES inspection_plans(id) ON DELETE CASCADE,
    name            VARCHAR(100) NOT NULL,
    nominal         NUMERIC(15,4),
    tolerance_upper NUMERIC(15,4),
    tolerance_lower NUMERIC(15,4),
    unit            VARCHAR(20),
    is_critical     BOOLEAN NOT NULL DEFAULT FALSE
);

-- Registros de inspeção (por OP ou lote de recebimento)
CREATE TABLE quality_records (
    id                   BIGSERIAL PRIMARY KEY,
    plan_id              BIGINT NOT NULL REFERENCES inspection_plans(id),
    production_order_id  BIGINT REFERENCES production_orders(id),
    lot                  VARCHAR(50),
    item_code            BIGINT NOT NULL,
    inspected_qty        NUMERIC(15,4) NOT NULL,
    approved_qty         NUMERIC(15,4) NOT NULL DEFAULT 0,
    rejected_qty         NUMERIC(15,4) NOT NULL DEFAULT 0,
    result               inspection_result_enum NOT NULL DEFAULT 'PENDENTE',
    inspector_id         BIGINT REFERENCES employees(id),
    inspected_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    notes                TEXT,
    created_at           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by           UUID NOT NULL
);

CREATE INDEX idx_qr_order ON quality_records(production_order_id);
CREATE INDEX idx_qr_item ON quality_records(item_code);
CREATE INDEX idx_qr_lot ON quality_records(lot);

-- Medições das características
CREATE TABLE quality_measurements (
    id              BIGSERIAL PRIMARY KEY,
    record_id       BIGINT NOT NULL REFERENCES quality_records(id) ON DELETE CASCADE,
    characteristic_id BIGINT NOT NULL REFERENCES inspection_plan_characteristics(id),
    measured_value  NUMERIC(15,4) NOT NULL,
    is_conformant   BOOLEAN NOT NULL
);

-- Não-conformidades
CREATE TABLE non_conformances (
    id                  BIGSERIAL PRIMARY KEY,
    code                BIGINT NOT NULL UNIQUE,
    quality_record_id   BIGINT REFERENCES quality_records(id),
    production_order_id BIGINT REFERENCES production_orders(id),
    item_code           BIGINT NOT NULL,
    lot                 VARCHAR(50),
    nonconform_qty      NUMERIC(15,4) NOT NULL,
    description         TEXT NOT NULL,
    severity            nc_severity_enum NOT NULL DEFAULT 'MENOR',
    root_cause          TEXT,
    corrective_action   TEXT,
    disposition         nc_disposition_enum,
    disposed_at         TIMESTAMPTZ,
    disposed_by         UUID,
    is_open             BOOLEAN NOT NULL DEFAULT TRUE,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by          UUID NOT NULL
);

CREATE INDEX idx_nc_item ON non_conformances(item_code);
CREATE INDEX idx_nc_order ON non_conformances(production_order_id);
CREATE INDEX idx_nc_open ON non_conformances(is_open) WHERE is_open = TRUE;

COMMIT;
