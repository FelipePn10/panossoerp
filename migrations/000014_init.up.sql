-- -- ===============================
-- -- UNITS (Unidade de Medida)
-- -- ===============================  
-- CREATE TABLE units (
--     id BIGSERIAL PRIMARY KEY,
--     code VARCHAR(10) NOT NULL UNIQUE,
--     description VARCHAR(50) NOT NULL,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW()
-- );

-- -- ===============================
-- -- WAREHOUSES (Almoxarifados)
-- -- ===============================
-- CREATE TABLE warehouses (
--     id BIGSERIAL PRIMARY KEY,
--     code VARCHAR(50) NOT NULL UNIQUE,
--     name VARCHAR(150) NOT NULL,
--     active BOOLEAN NOT NULL DEFAULT TRUE,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     deleted_at TIMESTAMP
-- );

-- -- ===============================
-- -- WAREHOUSE LOCATIONS
-- -- ===============================
-- CREATE TABLE warehouse_locations (
--     id BIGSERIAL PRIMARY KEY,
--     warehouse_id BIGINT NOT NULL,
--     code VARCHAR(50) NOT NULL,
--     description VARCHAR(150),
--     active BOOLEAN NOT NULL DEFAULT TRUE,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     deleted_at TIMESTAMP,

--     CONSTRAINT fk_location_warehouse
--         FOREIGN KEY (warehouse_id)
--         REFERENCES warehouses(id),

--     CONSTRAINT uq_warehouse_location
--         UNIQUE (warehouse_id, code)
-- );

-- -- ===============================
-- -- ITEMS (Materiais / Produtos)
-- -- ===============================
-- CREATE TABLE items (
--     id BIGSERIAL PRIMARY KEY,

--     internal_code VARCHAR(50) NOT NULL,
--     sku VARCHAR(80),
--     name VARCHAR(200) NOT NULL,

--     type VARCHAR(30) NOT NULL,
--     -- PRODUCT | RAW_MATERIAL | CONSUMABLE | SERVICE

--     unit_id BIGINT NOT NULL,

--     active BOOLEAN NOT NULL DEFAULT TRUE,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     deleted_at TIMESTAMP,

--     CONSTRAINT uq_item_internal_code UNIQUE (internal_code),
--     CONSTRAINT uq_item_sku UNIQUE (sku),

--     CONSTRAINT fk_item_unit
--         FOREIGN KEY (unit_id)
--         REFERENCES units(id)
-- );

-- -- ===============================
-- -- STOCK LOTS (Lotes - opcional)
-- -- ===============================
-- CREATE TABLE stock_lots (
--     id BIGSERIAL PRIMARY KEY,
--     item_id BIGINT NOT NULL,
--     lot_code VARCHAR(80) NOT NULL,
--     expiration_date DATE,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW(),

--     CONSTRAINT uq_item_lot UNIQUE (item_id, lot_code),

--     CONSTRAINT fk_lot_item
--         FOREIGN KEY (item_id)
--         REFERENCES items(id)
-- );

-- -- ===============================
-- -- STOCK MOVEMENTS (EVENT STORE)
-- -- ===============================
-- CREATE TABLE stock_movements (
--     id BIGSERIAL PRIMARY KEY,

--     item_id BIGINT NOT NULL,
--     warehouse_id BIGINT NOT NULL,
--     location_id BIGINT,
--     lot_id BIGINT,

--     quantity NUMERIC(14,4) NOT NULL,
--     movement_type VARCHAR(20) NOT NULL,
--     -- IN | OUT | TRANSFER | ADJUST | INVENTORY

--     reference_type VARCHAR(50),
--     reference_id BIGINT,

--     created_by BIGINT,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW(),

--     CONSTRAINT fk_movement_item
--         FOREIGN KEY (item_id)
--         REFERENCES items(id),

--     CONSTRAINT fk_movement_warehouse
--         FOREIGN KEY (warehouse_id)
--         REFERENCES warehouses(id),

--     CONSTRAINT fk_movement_location
--         FOREIGN KEY (location_id)
--         REFERENCES warehouse_locations(id),

--     CONSTRAINT fk_movement_lot
--         FOREIGN KEY (lot_id)
--         REFERENCES stock_lots(id)
-- );

-- -- ===============================
-- -- STOCK BALANCES (DERIVED CACHE)
-- -- ===============================
-- CREATE TABLE stock_balances (
--     item_id BIGINT NOT NULL,
--     warehouse_id BIGINT NOT NULL,
--     location_id BIGINT,
--     lot_id BIGINT,

--     quantity NUMERIC(14,4) NOT NULL DEFAULT 0,

--     PRIMARY KEY (item_id, warehouse_id, location_id, lot_id),

--     CONSTRAINT fk_balance_item
--         FOREIGN KEY (item_id)
--         REFERENCES items(id),

--     CONSTRAINT fk_balance_warehouse
--         FOREIGN KEY (warehouse_id)
--         REFERENCES warehouses(id),

--     CONSTRAINT fk_balance_location
--         FOREIGN KEY (location_id)
--         REFERENCES warehouse_locations(id),

--     CONSTRAINT fk_balance_lot
--         FOREIGN KEY (lot_id)
--         REFERENCES stock_lots(id)
-- );