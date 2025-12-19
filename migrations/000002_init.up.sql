CREATE TYPE component_type AS ENUM (
    'STRUCTURE',
    'SET',
    'ITEM'
);

CREATE TABLE products (
    id UUID PRIMARY KEY,
    code VARCHAR(10) NOT NULL UNIQUE,
    group_code SMALLINT NOT NULL,
    name TEXT NOT NULL,

    created_by UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE product_masks (
    id UUID PRIMARY KEY,

    product_id UUID NOT NULL REFERENCES products(id),
    mask TEXT NOT NULL,
    mask_hash CHAR(8) NOT NULL,

    business_id VARCHAR(30) NOT NULL,

    created_by UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,

    UNIQUE (business_id),
    UNIQUE (product_id, mask_hash)
);

CREATE TABLE components (
    id UUID PRIMARY KEY,

    code VARCHAR(10) NOT NULL,
    name TEXT NOT NULL,
    type component_type NOT NULL,

    created_at TIMESTAMPTZ NOT NULL,

    UNIQUE (code, type)
);


CREATE TABLE component_masks (
    id UUID PRIMARY KEY,

    component_id UUID NOT NULL REFERENCES components(id),
    mask TEXT NOT NULL,
    mask_hash CHAR(8) NOT NULL,

    business_id VARCHAR(30) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL,

    UNIQUE (business_id),
    UNIQUE (component_id, mask_hash)
);

CREATE TABLE mask_composition (
    parent_mask_id UUID NOT NULL REFERENCES component_masks(id),
    child_mask_id UUID NOT NULL REFERENCES component_masks(id),

    quantity NUMERIC(10,2) NOT NULL CHECK (quantity > 0),

    PRIMARY KEY (parent_mask_id, child_mask_id)
);


CREATE TABLE material_consumption (
    component_mask_id UUID NOT NULL REFERENCES component_masks(id),
    material_id UUID NOT NULL REFERENCES components(id),

    quantity NUMERIC(10,2) NOT NULL CHECK (quantity > 0),
    unit VARCHAR(10) NOT NULL,

    PRIMARY KEY (component_mask_id, material_id)
);

-- Products
CREATE INDEX idx_products_code ON products(code);

-- Mask product
CREATE INDEX idx_product_masks_business_id 
ON product_masks(business_id);

-- Components
CREATE INDEX idx_components_code_type 
ON components(code, type);

CREATE INDEX idx_components_type 
ON components(type);

-- Mask Components
CREATE INDEX idx_component_masks_business_id 
ON component_masks(business_id);

-- Hierarchy
CREATE INDEX idx_mask_composition_parent 
ON mask_composition(parent_mask_id);
