-- industrial operations (steps)
CREATE TABLE operations (
    id BIGSERIAL PRIMARY KEY,
    code VARCHAR(50) NOT NULL UNIQUE,   -- CORTE, COSTURA, MONTAGEM
    name VARCHAR(255) NOT NULL
);


CREATE TABLE boms (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id),

    bom_type VARCHAR(10) NOT NULL,             -- EBOM | MBOM
    version INTEGER NOT NULL,
    status VARCHAR(20) NOT NULL,               -- draft | released | obsolete

    valid_from DATE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),

    CHECK (bom_type IN ('EBOM','MBOM')),
    CHECK (status IN ('draft','released','obsolete')),

    UNIQUE (product_id, bom_type, version)
);


CREATE TABLE bom_items (
    id BIGSERIAL PRIMARY KEY,
    bom_id BIGINT NOT NULL REFERENCES boms(id),

    component_id BIGINT NOT NULL REFERENCES products(id),

    quantity NUMERIC(14,6) NOT NULL,
    uom VARCHAR(10),

    scrap_percent NUMERIC(5,2) NOT NULL DEFAULT 0,

    operation_sequence INTEGER NOT NULL,
    operation_id BIGINT NOT NULL REFERENCES operations(id),

    created_at TIMESTAMP NOT NULL DEFAULT now(),

    UNIQUE (bom_id, component_id, operation_sequence)
);


-- Production Order (where the process takes place)
CREATE TABLE production_orders (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id),
    bom_id BIGINT NOT NULL REFERENCES boms(id),
    quantity NUMERIC(14,6) NOT NULL,
    status VARCHAR(20) NOT NULL,
    current_operation_id BIGINT REFERENCES operations(id),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);


CREATE TABLE stock (
    product_id BIGINT PRIMARY KEY REFERENCES products(id),
    quantity NUMERIC(14,6) NOT NULL DEFAULT 0
);

CREATE TABLE stock_movements (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id),

    movement_type VARCHAR(10) NOT NULL, -- IN | OUT
    quantity NUMERIC(14,6) NOT NULL,

    reference_type VARCHAR(20),          -- PURCHASE, PRODUCTION, SALE
    reference_id BIGINT,

    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE mrp_parameters (
    product_id BIGINT PRIMARY KEY REFERENCES products(id),

    planning_type VARCHAR(10) NOT NULL, -- buy | make
    lead_time_days INTEGER NOT NULL,
    safety_stock NUMERIC(14,6) NOT NULL DEFAULT 0
);

ALTER TABLE products
DROP COLUMN updated_at;

ALTER TABLE products
ADD COLUMN product_type VARCHAR(20) NOT NULL DEFAULT 'finished',
ADD COLUMN uom VARCHAR(20);

UPDATE products
SET uom = 'UN'
WHERE uom IS NULL;

ALTER TABLE products
ALTER COLUMN uom SET NOT NULL;



CREATE INDEX idx_bom_items_bom ON bom_items(bom_id);
CREATE INDEX idx_stock_movements_product ON stock_movements(product_id);
CREATE INDEX idx_production_orders_status ON production_orders(status);

-- finished | semi_finished | raw_material | indirect
