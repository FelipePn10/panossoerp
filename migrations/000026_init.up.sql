-- Remove colunas antigas que foram substituídas pela estrutura PDM
ALTER TABLE items DROP COLUMN IF EXISTS name;
ALTER TABLE items DROP COLUMN IF EXISTS description;
ALTER TABLE items DROP COLUMN IF EXISTS type;
ALTER TABLE items DROP COLUMN IF EXISTS status;

-- Root
ALTER TABLE items ADD COLUMN complement            TEXT;
ALTER TABLE items ADD COLUMN nature                SMALLINT    NOT NULL DEFAULT 0;
ALTER TABLE items ADD COLUMN situation             SMALLINT    NOT NULL DEFAULT 0;
-- health já existe

-- PDM
ALTER TABLE items ADD COLUMN pdm_group_id              INT         NOT NULL DEFAULT 0;
ALTER TABLE items ADD COLUMN pdm_modifier_id            INT         NOT NULL DEFAULT 0;
ALTER TABLE items ADD COLUMN pdm_attributes             JSONB       NOT NULL DEFAULT '[]';
ALTER TABLE items ADD COLUMN pdm_description_technique  TEXT        NOT NULL DEFAULT '';

-- Warehouse (warehouse_id já existe)
ALTER TABLE items ADD COLUMN warehouse_unit_of_measurement          SMALLINT    NOT NULL DEFAULT 0;
ALTER TABLE items ADD COLUMN warehouse_automatic_low                BOOLEAN     NOT NULL DEFAULT FALSE;
ALTER TABLE items ADD COLUMN warehouse_cyclical_count_config        JSONB;
ALTER TABLE items ADD COLUMN warehouse_minimum_stock                INT         NOT NULL DEFAULT 0;
ALTER TABLE items ADD COLUMN warehouse_avg_monthly_consumption_manual INT;

-- Engineering
ALTER TABLE items ADD COLUMN engineering_item_base_cod  INT;
ALTER TABLE items ADD COLUMN engineering_weight         JSONB       NOT NULL DEFAULT '{}';
ALTER TABLE items ADD COLUMN engineering_dimensions     JSONB;
ALTER TABLE items ADD COLUMN engineering_type           SMALLINT    NOT NULL DEFAULT 0;
ALTER TABLE items ADD COLUMN engineering_type_struct    SMALLINT    NOT NULL DEFAULT 0;
ALTER TABLE items ADD COLUMN engineering_oem            BOOLEAN     NOT NULL DEFAULT FALSE;

-- Planning
ALTER TABLE items ADD COLUMN planning_type_mrp      SMALLINT    NOT NULL DEFAULT 0;
ALTER TABLE items ADD COLUMN planning_llc           INT         NOT NULL DEFAULT 0;
ALTER TABLE items ADD COLUMN planning_reorder_point JSONB;
ALTER TABLE items ADD COLUMN planning_tank_id       INT;
ALTER TABLE items ADD COLUMN planning_ghost         BOOLEAN     NOT NULL DEFAULT FALSE;

-- Planners
ALTER TABLE items ADD COLUMN planner_employee_id    INT;

-- Supplies
ALTER TABLE items ADD COLUMN supplies_type_of_use   SMALLINT    NOT NULL DEFAULT 0;

-- Tabela separada para MachineUsage (relação 1:N com items)
CREATE TABLE item_machine_usages (
    id          BIGSERIAL   PRIMARY KEY,
    item_id     BIGINT      NOT NULL REFERENCES items(id) ON DELETE CASCADE,
    machine_id  INT         NOT NULL,
    usage_time  INT         NOT NULL
);

CREATE INDEX idx_item_machine_usages_item_id ON item_machine_usages(item_id);
