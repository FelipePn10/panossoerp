-- 1. ENUM
DO $$
BEGIN
    CREATE TYPE component_type AS ENUM ('STRUCTURE', 'SET', 'ITEM');
EXCEPTION
    WHEN duplicate_object THEN null;
END
$$;

-- 2. Users
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- 3. Products
CREATE TABLE products (
    id BIGINT PRIMARY KEY,
    code VARCHAR(10) NOT NULL UNIQUE,
    group_code VARCHAR(20),
    name TEXT NOT NULL,
    created_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

-- 4. Product Masks
CREATE TABLE product_masks (
    id BIGINT PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id),
    product_code VARCHAR(10) NOT NULL,
    mask TEXT NOT NULL,
    mask_hash CHAR(8) NOT NULL,
    business_id VARCHAR(30) NOT NULL,
    created_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL,
    UNIQUE (business_id),
    UNIQUE (product_id, mask_hash)
);

-- 5. Components
CREATE TABLE components (
    id BIGINT PRIMARY KEY,
    code VARCHAR(10) NOT NULL,
    name TEXT NOT NULL,
    type component_type NOT NULL,
    created_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL,
    UNIQUE (code, type)
);

-- 6. Component Masks
CREATE TABLE component_masks (
    id BIGINT PRIMARY KEY,
    component_id BIGINT NOT NULL REFERENCES components(id),
    mask TEXT NOT NULL,
    mask_hash CHAR(8) NOT NULL,
    business_id VARCHAR(30) NOT NULL,
    created_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL,
    UNIQUE (business_id),
    UNIQUE (component_id, mask_hash)
);

-- 7. Mask Composition
CREATE TABLE mask_composition (
    parent_mask_id BIGINT NOT NULL REFERENCES component_masks(id),
    child_mask_id BIGINT NOT NULL REFERENCES component_masks(id),
    quantity NUMERIC(10,2) NOT NULL CHECK (quantity > 0),
    PRIMARY KEY (parent_mask_id, child_mask_id)
);

-- 8. Material Consumption
CREATE TABLE material_consumption (
    component_mask_id BIGINT NOT NULL REFERENCES component_masks(id),
    material_id BIGINT NOT NULL REFERENCES components(id),
    quantity NUMERIC(10,2) NOT NULL CHECK (quantity > 0),
    unit VARCHAR(10) NOT NULL,
    PRIMARY KEY (component_mask_id, material_id)
);

-- 9. Complements
CREATE TABLE complement_a (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    value TEXT NOT NULL UNIQUE
);

CREATE TABLE complement_b (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    value TEXT NOT NULL UNIQUE
);

-- 10. Questions
CREATE TABLE questions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    createdby UUID NOT NULL REFERENCES users(id),
    complement_a_id BIGINT NOT NULL REFERENCES complement_a(id),
    complement_b_id BIGINT NOT NULL REFERENCES complement_b(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- 11. Indexes
CREATE INDEX idx_products_code ON products(code);
CREATE INDEX idx_product_masks_business_id ON product_masks(business_id);
CREATE INDEX idx_components_code_type ON components(code, type);
CREATE INDEX idx_components_type ON components(type);
CREATE INDEX idx_component_masks_business_id ON component_masks(business_id);
CREATE INDEX idx_mask_composition_parent ON mask_composition(parent_mask_id);
CREATE INDEX idx_products_created_by ON products(created_by);
CREATE INDEX idx_components_created_by ON components(created_by);
CREATE INDEX idx_component_masks_created_by ON component_masks(created_by);