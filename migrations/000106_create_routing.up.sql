BEGIN;

-- ─── enums ───────────────────────────────────────────────────────────────────

DO $$ BEGIN
    CREATE TYPE operation_origin_enum AS ENUM ('INTERNA', 'EXTERNA', 'TERCEIROS');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

DO $$ BEGIN
    CREATE TYPE operation_situation_enum AS ENUM ('APROVADA', 'INATIVA');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

DO $$ BEGIN
    CREATE TYPE route_situation_enum AS ENUM ('APROVADA', 'INATIVA');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

DO $$ BEGIN
    CREATE TYPE route_op_situation_enum AS ENUM ('APROVADA', 'INATIVA', 'FANTASMA');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

-- ─── operations ───────────────────────────────────────────────────────────────
-- Drop legacy table from migration 000015 (old schema: code VARCHAR, no routing cols).
-- CASCADE removes any FK constraints that point to it from the legacy bom_items /
-- production_orders tables created in that same migration.

DROP TABLE IF EXISTS operations CASCADE;

CREATE TABLE operations (
    id                     BIGSERIAL PRIMARY KEY,
    code                   BIGINT NOT NULL UNIQUE,
    name                   VARCHAR(100) NOT NULL,
    description            TEXT,
    origin                 operation_origin_enum NOT NULL DEFAULT 'INTERNA',
    situation              operation_situation_enum NOT NULL DEFAULT 'APROVADA',
    default_work_center_id BIGINT REFERENCES machine_types(id),
    standard_time          NUMERIC(15,4) NOT NULL DEFAULT 0,
    setup_time             NUMERIC(15,4) NOT NULL DEFAULT 0,
    is_active              BOOLEAN NOT NULL DEFAULT TRUE,
    created_at             TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at             TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by             UUID NOT NULL
);

CREATE INDEX idx_operations_code ON operations(code);
CREATE INDEX idx_operations_work_center ON operations(default_work_center_id);

-- ─── manufacturing_routes ─────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS manufacturing_routes (
    id          BIGSERIAL PRIMARY KEY,
    code        BIGINT NOT NULL UNIQUE,
    item_code   BIGINT NOT NULL REFERENCES items(code),
    mask        VARCHAR(200),
    alternative SMALLINT NOT NULL DEFAULT 1,
    description VARCHAR(200),
    situation   route_situation_enum NOT NULL DEFAULT 'APROVADA',
    is_standard BOOLEAN NOT NULL DEFAULT FALSE,
    is_active   BOOLEAN NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by  UUID NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_mfg_routes_item ON manufacturing_routes(item_code);
CREATE INDEX IF NOT EXISTS idx_mfg_routes_code ON manufacturing_routes(code);
CREATE UNIQUE INDEX IF NOT EXISTS idx_mfg_routes_unique ON manufacturing_routes(item_code, COALESCE(mask, ''), alternative);

-- ─── route_operations ─────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS route_operations (
    id             BIGSERIAL PRIMARY KEY,
    route_id       BIGINT NOT NULL REFERENCES manufacturing_routes(id) ON DELETE CASCADE,
    sequence       SMALLINT NOT NULL,
    operation_id   BIGINT NOT NULL REFERENCES operations(id),
    work_center_id BIGINT REFERENCES machine_types(id),
    standard_time  NUMERIC(15,4),
    setup_time     NUMERIC(15,4),
    situation      route_op_situation_enum NOT NULL DEFAULT 'APROVADA',
    notes          TEXT,
    is_active      BOOLEAN NOT NULL DEFAULT TRUE,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (route_id, sequence)
);

CREATE INDEX IF NOT EXISTS idx_route_ops_route ON route_operations(route_id);
CREATE INDEX IF NOT EXISTS idx_route_ops_operation ON route_operations(operation_id);
CREATE INDEX IF NOT EXISTS idx_route_ops_wc ON route_operations(work_center_id);

-- ─── route_operation_network ──────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS route_operation_network (
    id             BIGSERIAL PRIMARY KEY,
    predecessor_id BIGINT NOT NULL REFERENCES route_operations(id) ON DELETE CASCADE,
    successor_id   BIGINT NOT NULL REFERENCES route_operations(id) ON DELETE CASCADE,
    overlap_pct    NUMERIC(5,2) NOT NULL DEFAULT 0 CHECK (overlap_pct BETWEEN 0 AND 100),
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (predecessor_id, successor_id)
);

CREATE INDEX IF NOT EXISTS idx_ron_predecessor ON route_operation_network(predecessor_id);
CREATE INDEX IF NOT EXISTS idx_ron_successor ON route_operation_network(successor_id);

COMMIT;
